package chardet

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/log"
	"github.com/wlynxg/chardet/probe"
	"github.com/wlynxg/chardet/smm"
	"go.uber.org/zap"
)

type Result struct {
	Encoding   string  `json:"encoding,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
	Language   string  `json:"language,omitempty"`
}

type UniversalDetector struct {
	MinimumThreshold float64
	HighByteDetector *regexp.Regexp
	EscDetector      *regexp.Regexp
	WinByteDetector  *regexp.Regexp
	IsoWinMap        map[string]string

	log             *zap.SugaredLogger
	done            bool
	gotData         bool
	hasWinBytes     bool
	inputState      consts.InputState
	lastChars       []byte
	escCharsetProbe *smm.EscCharSetProbe
	langFilter      consts.LangFilter
	charsetProbes   []probe.ICharSetProbe
	filter          consts.LangFilter
	result          *Result
}

func NewUniversalDetector(filter consts.LangFilter) *UniversalDetector {
	if filter == consts.UnknownLangFilter {
		filter = consts.AllLangFilter
	}

	return &UniversalDetector{
		MinimumThreshold: 0.20,
		HighByteDetector: regexp.MustCompile(`[\x80-\xFF]`),
		EscDetector:      regexp.MustCompile(`(\x1B|~{)`),
		WinByteDetector:  regexp.MustCompile(`[\x80-\x9F]`),
		IsoWinMap: map[string]string{
			"iso-8859-1":  "Windows-1252",
			"iso-8859-2":  "Windows-1250",
			"iso-8859-5":  "Windows-1251",
			"iso-8859-6":  "Windows-1256",
			"iso-8859-7":  "Windows-1253",
			"iso-8859-8":  "Windows-1255",
			"iso-8859-9":  "Windows-1254",
			"iso-8859-13": "Windows-1257",
		},

		log:        log.New("UniversalDetector"),
		inputState: consts.PureAsciiInputState,
		lastChars:  []byte{},
		filter:     filter,
	}
}

func (u *UniversalDetector) Reset() {
	u.result = &Result{}
	u.done = false
	u.gotData = false
	u.hasWinBytes = false
	u.inputState = consts.PureAsciiInputState
	u.lastChars = []byte{}

	if u.escCharsetProbe != nil {
		u.escCharsetProbe.Reset()
	}

	for _, p := range u.charsetProbes {
		if p != nil {
			p.Reset()
		}
	}
}

func (u *UniversalDetector) Feed(data []byte) {
	if u.done || len(data) == 0 {
		return
	}

	// First check for known BOMs, since these are guaranteed to be correct
	if !u.gotData {
		// If the data starts with BOM, we know it is UTF
		var encoding string
		switch {
		case bytes.HasPrefix(data, consts.UTF8BOM):
			// EF BB BF  UTF-8 with BOM
			encoding = consts.UTF8SIG
		case bytes.HasPrefix(data, consts.UTF32LEBOM) || bytes.HasPrefix(data, consts.UTF32BEBOM):
			// FF FE 00 00  UTF-32, little-endian BOM
			// 00 00 FE FF  UTF-32, big-endian BOM
			encoding = consts.UTF32
		case bytes.HasPrefix(data, consts.UTF16LEBOM) || bytes.HasPrefix(data, consts.UTF16BEBOM):
			// FF FE  UTF-16, little endian BOM
			// FE FF  UTF-16, big endian BOM
			encoding = consts.UTF16
		case bytes.HasPrefix(data, consts.UCS43412BOM):
			// FE FF 00 00  UCS-4, unusual octet order BOM (3412)
			encoding = consts.UCS43412
		case bytes.HasPrefix(data, consts.UCS42143BOM):
			// 00 00 FF FE  UCS-4, unusual octet order BOM (2143)
			encoding = consts.UCS42143
		}

		u.gotData = true
		if encoding != "" {
			u.result = &Result{
				Encoding:   encoding,
				Confidence: 1.0,
				Language:   "",
			}
			u.done = true
			return
		}
	}

	// If none of those matched, and we've only seen ASCII so far, check
	// for high bytes and escape sequences.
	if u.inputState == consts.PureAsciiInputState {
		if u.HighByteDetector.Match(data) {
			u.inputState = consts.HighByteInputState
		} else if u.inputState == consts.PureAsciiInputState &&
			u.EscDetector.Match(bytes.Join([][]byte{u.lastChars, data}, nil)) {
			u.inputState = consts.EcsAsciiInputState
		}
	}

	u.lastChars = append(u.lastChars, data[len(data)-1])

	switch u.inputState {
	case consts.EcsAsciiInputState:
		// If we've seen escape sequences, use the EscCharSetProbe, which
		// uses a simple state machine to check for known escape sequences in
		// HZ and ISO-2022 encodings, since those are the only encodings that
		// use such sequences.
		if u.escCharsetProbe == nil {
			u.escCharsetProbe = smm.NewEscCharSetProbe(u.langFilter)
		}

		if u.escCharsetProbe.Feed(data) == consts.FoundItProbingState {
			u.result = &Result{
				Encoding:   u.escCharsetProbe.CharsetName(),
				Confidence: u.escCharsetProbe.Confidence(),
				Language:   u.escCharsetProbe.Language(),
			}
			u.done = true
		}
	case consts.HighByteInputState:
		// If we've seen high bytes (i.e., those with values greater than 127),
		// we need to do more complicated checks using all our multi-byte and
		// single-byte probes that are left.  The single-byte probes
		// use character bigram distributions to determine the encoding, whereas
		// the multi-byte probes use a combination of character unigram and
		// bigram distributions.
		if len(u.charsetProbes) == 0 {
			u.charsetProbes = []probe.ICharSetProbe{probe.MBCGroupProbe(u.filter)}
			// If we're checking non-CJK encodings, use single-byte probe
			if u.filter&consts.NonCjkLangFilter != 0 {
				u.charsetProbes = append(u.charsetProbes, probe.NewSBCSGroupProbe())
			}
			u.charsetProbes = append(u.charsetProbes, probe.NewLatin1Probe())
		}

		for _, charsetProbe := range u.charsetProbes {
			if charsetProbe == nil {
				continue
			}

			if charsetProbe.Feed(data) == consts.FoundItProbingState {
				u.result = &Result{
					Encoding:   charsetProbe.CharSetName(),
					Confidence: charsetProbe.GetConfidence(),
					Language:   charsetProbe.Language(),
				}
				u.done = true
			}
		}

		if u.WinByteDetector.Match(data) {
			u.hasWinBytes = true
		}
	default:
	}
}

func (u *UniversalDetector) GetResult() Result {
	if u.done {
		return *u.result
	}
	u.done = true

	switch {
	case !u.gotData:
		u.log.Debug("no data received!")
	case u.inputState == consts.PureAsciiInputState:
		u.result = &Result{
			Encoding:   consts.Ascii,
			Confidence: 1.0,
			Language:   "",
		}
	case u.inputState == consts.HighByteInputState:
		var (
			probeConfidence    = 0.0
			maxProbeConfidence = 0.0
			maxConfidenceProbe probe.ICharSetProbe
		)

		for _, charsetProbe := range u.charsetProbes {
			if charsetProbe == nil {
				continue
			}

			probeConfidence = charsetProbe.GetConfidence()
			if probeConfidence > maxProbeConfidence {
				maxProbeConfidence = probeConfidence
				maxConfidenceProbe = charsetProbe
			}
		}

		if maxConfidenceProbe != nil && maxProbeConfidence > u.MinimumThreshold {
			charsetName := maxConfidenceProbe.CharSetName()
			confidence := maxConfidenceProbe.GetConfidence()

			if u.hasWinBytes {
				// Use Windows encoding name instead of ISO-8859 if we saw any
				// extra Windows-specific bytes
				if n, ok := u.IsoWinMap[strings.ToLower(maxConfidenceProbe.CharSetName())]; ok {
					charsetName = n
				}
			}
			u.result = &Result{
				Encoding:   charsetName,
				Confidence: confidence,
				Language:   maxConfidenceProbe.Language(),
			}
		}
	}
	return *u.result
}
