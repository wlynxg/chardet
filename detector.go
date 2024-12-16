package chardet

import (
	"bytes"

	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/probe"
)

// Result represents the character encoding detection result
type Result struct {
	// Encoding is the detected character encoding name
	Encoding string `json:"encoding,omitempty"`
	// Confidence indicates how confident the detector is about the result (0.0-1.0)
	Confidence float64 `json:"confidence,omitempty"`
	// Language represents the detected language (if applicable)
	Language string `json:"language,omitempty"`
}

// UniversalDetector implements universal character encoding detection
type UniversalDetector struct {
	// MinimumThreshold is the minimum confidence threshold for detection
	MinimumThreshold float64
	// IsoWinMap maps ISO encodings to Windows encodings
	IsoWinMap map[string]string

	// done indicates if detection is complete
	done bool
	// gotData indicates if any data has been processed
	gotData bool
	// hasWinBytes indicates if Windows-specific bytes were detected
	hasWinBytes bool

	// lastChars stores the last processed characters
	lastChars []byte
	// inputState tracks the current input processing state
	inputState consts.InputState
	// filter specifies which languages to detect
	filter consts.LangFilter

	// Various probes used for detection
	escCharsetProbe probe.Probe
	utf1632Probe    *probe.UTF1632Probe
	charsetProbes   []probe.Probe

	// result stores the final detection result
	result Result
}

// NewUniversalDetector creates a new UniversalDetector instance with the specified language filter
func NewUniversalDetector(filter consts.LangFilter) *UniversalDetector {
	if filter == consts.UnknownLangFilter {
		filter = consts.AllLangFilter
	}

	return &UniversalDetector{
		MinimumThreshold: 0.20,
		IsoWinMap: map[string]string{
			consts.ISO88591:  consts.Windows1252,
			consts.ISO88592:  consts.Windows1250,
			consts.ISO88595:  consts.Windows1251,
			consts.ISO88596:  consts.Windows1256,
			consts.ISO88597:  consts.Windows1253,
			consts.ISO88598:  consts.Windows1255,
			consts.ISO88599:  consts.Windows1254,
			consts.ISO885913: consts.Windows1257,
		},

		inputState: consts.PureAsciiInputState,
		lastChars:  []byte{},
		filter:     filter,
	}
}

// Reset resets the detector to its initial state
func (u *UniversalDetector) Reset() {
	u.result = Result{}
	u.done = false
	u.gotData = false
	u.hasWinBytes = false
	u.inputState = consts.PureAsciiInputState
	u.lastChars = []byte{}

	if u.escCharsetProbe != nil {
		u.escCharsetProbe.Reset()
	}

	if u.utf1632Probe != nil {
		u.utf1632Probe.Reset()
	}

	for _, p := range u.charsetProbes {
		if p != nil {
			p.Reset()
		}
	}
}

// Feed processes a chunk of bytes for character encoding detection
// It analyzes the input data and updates the internal state accordingly
func (u *UniversalDetector) Feed(buf []byte) {
	if u.done || len(buf) == 0 {
		return
	}

	// First check for known BOMs, since these are guaranteed to be correct
	if !u.gotData {
		// If the buf starts with BOM, we know it is UTF
		var encoding string

		if bytes.HasPrefix(buf, []byte(consts.UTF8BOM)) {
			// EF BB BF  UTF-8 with BOM
			encoding = consts.UTF8SIG
		} else if bytes.HasPrefix(buf, []byte(consts.UTF32LEBOM)) || bytes.HasPrefix(buf, []byte(consts.UTF32BEBOM)) {
			// FF FE 00 00  UTF-32, little-endian BOM
			// 00 00 FE FF  UTF-32, big-endian BOM
			encoding = consts.UTF32
		} else if bytes.HasPrefix(buf, []byte(consts.UCS43412BOM)) {
			// FE FF 00 00  UCS-4, unusual octet order BOM (3412)
			encoding = consts.UCS43412
		} else if bytes.HasPrefix(buf, []byte(consts.UCS42143BOM)) {
			// 00 00 FF FE  UCS-4, unusual octet order BOM (2143)
			encoding = consts.UCS42143
		} else if bytes.HasPrefix(buf, []byte(consts.UTF16LEBOM)) || bytes.HasPrefix(buf, []byte(consts.UTF16BEBOM)) {
			// FF FE  UTF-16, little endian BOM
			// FE FF  UTF-16, big endian BOM
			encoding = consts.UTF16
		}

		u.gotData = true
		if encoding != "" {
			u.result = Result{
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
		if HighByteDetector(buf) {
			u.inputState = consts.HighByteInputState
		} else if u.inputState == consts.PureAsciiInputState &&
			EscDetector(bytes.Join([][]byte{u.lastChars, buf}, nil)) {
			u.inputState = consts.EcsAsciiInputState
		}
	}

	u.lastChars = append(u.lastChars, buf[len(buf)-1])

	// next we will look to see if it is appears to be either a UTF-16 or UTF-32 encoding
	if u.utf1632Probe == nil {
		u.utf1632Probe = probe.NewUTF1632Probe()
	}

	if u.utf1632Probe.State() == consts.DetectingProbingState {
		if u.utf1632Probe.Feed(buf) == consts.FoundItProbingState {
			u.result = Result{
				Encoding:   u.utf1632Probe.CharSetName(),
				Confidence: u.utf1632Probe.GetConfidence(),
				Language:   "",
			}
			u.done = true
			return
		}
	}

	switch u.inputState {
	case consts.EcsAsciiInputState:
		// If we've seen escape sequences, use the EscCharSetProbe, which
		// uses a simple state machine to check for known escape sequences in
		// HZ and ISO-2022 encodings, since those are the only encodings that
		// use such sequences.
		if u.escCharsetProbe == nil {
			u.escCharsetProbe = probe.NewEscCharSetProbe(u.filter)
		}

		if u.escCharsetProbe.Feed(buf) == consts.FoundItProbingState {
			u.result = Result{
				Encoding:   u.escCharsetProbe.CharSetName(),
				Confidence: u.escCharsetProbe.GetConfidence(),
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
			u.charsetProbes = []probe.Probe{probe.MBCGroupProbe(u.filter)}
			// If we're checking non-CJK encodings, use single-byte probe
			if u.filter&consts.NonCjkLangFilter != 0 {
				u.charsetProbes = append(u.charsetProbes, probe.NewSBCSGroupProbe())
			}
			u.charsetProbes = append(u.charsetProbes, probe.NewLatin1Probe(), probe.NewMacRomanProbe())
		}

		for _, charsetProbe := range u.charsetProbes {
			if charsetProbe == nil {
				continue
			}

			if charsetProbe.Feed(buf) == consts.FoundItProbingState {
				u.result = Result{
					Encoding:   charsetProbe.CharSetName(),
					Confidence: charsetProbe.GetConfidence(),
					Language:   charsetProbe.Language(),
				}
				u.done = true
				break
			}
		}

		if WinByteDetector(buf) {
			u.hasWinBytes = true
		}
	default:
	}
}

// GetResult returns the final character encoding detection result
// If detection is not complete, it will finalize the detection process
func (u *UniversalDetector) GetResult() Result {
	if u.done {
		return u.result
	}
	u.done = true

	switch {
	case !u.gotData:
	case u.inputState == consts.PureAsciiInputState:
		u.result = Result{
			Encoding:   consts.Ascii,
			Confidence: 1.0,
			Language:   "",
		}
	case u.inputState == consts.HighByteInputState:
		var (
			confidence, maxProbeConfidence float64
			maxConfidenceProbe             probe.Probe
		)

		for _, charsetProbe := range u.charsetProbes {
			if charsetProbe == nil {
				continue
			}

			confidence = charsetProbe.GetConfidence()
			if confidence > maxProbeConfidence {
				maxProbeConfidence = confidence
				maxConfidenceProbe = charsetProbe
			}
		}

		if maxConfidenceProbe != nil && maxProbeConfidence > u.MinimumThreshold {
			charsetName := maxConfidenceProbe.CharSetName()
			confidence := maxConfidenceProbe.GetConfidence()

			if u.hasWinBytes {
				// Use Windows encoding name instead of ISO-8859 if we saw any
				// extra Windows-specific bytes
				if n, ok := u.IsoWinMap[maxConfidenceProbe.CharSetName()]; ok {
					charsetName = n
				}
			}
			u.result = Result{
				Encoding:   charsetName,
				Confidence: confidence,
				Language:   maxConfidenceProbe.Language(),
			}
		}
	}
	return u.result
}

// HighByteDetector checks if the buffer contains any bytes with values >= 0x80
func HighByteDetector(buf []byte) bool {
	for _, b := range buf {
		if b >= 0x80 { // Check if the byte is in the international range
			return true
		}
	}
	return false
}

// EscDetector checks if the buffer contains escape sequences
// commonly used in certain character encodings
func EscDetector(buf []byte) bool {
	for i := 0; i < len(buf); i++ {
		// Check for ESC character (0x1B)
		if buf[i] == 0x1B {
			return true
		}
		// Check for ~ character (0x7E) followed by {
		if buf[i] == 0x7E && i+1 < len(buf) && buf[i+1] == '{' {
			return true
		}
	}
	return false
}

// WinByteDetector checks if the buffer contains Windows-specific byte values
// in the range 0x80-0x9F
func WinByteDetector(buf []byte) bool {
	for _, b := range buf {
		if b >= 0x80 && b <= 0x9F { // Check if the byte is in the specified range
			return true
		}
	}
	return false
}
