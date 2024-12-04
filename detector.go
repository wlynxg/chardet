package chardet

import (
	"bytes"
	"regexp"
)

type Result struct {
	Encoding   string  `json:"encoding,omitempty"`
	Confidence float32 `json:"confidence,omitempty"`
	Language   string  `json:"language,omitempty"`
}

type UniversalDetector struct {
	done       bool
	gotData    bool
	inputState InputState

	highByteDetector *regexp.Regexp
	escDetector      *regexp.Regexp
	lastChars        []byte

	escCharsetProbe *EscCharSetProbe
	langFilter      LangFilter

	result *Result
}

func NewUniversalDetector() *UniversalDetector {
	highByteDetector, _ := regexp.Compile(`[\x80-\xFF]`)
	escDetector, _ := regexp.Compile(`(\x1B|~{)`)

	return &UniversalDetector{
		inputState:       PureAsciiInputState,
		highByteDetector: highByteDetector,
		escDetector:      escDetector,
		lastChars:        []byte{},
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
		case bytes.HasPrefix(data, utf8BOM):
			// EF BB BF  UTF-8 with BOM
			encoding = UTF8SIG
		case bytes.HasPrefix(data, utf32LEBOM) || bytes.HasPrefix(data, utf32BEBOM):
			// FF FE 00 00  UTF-32, little-endian BOM
			// 00 00 FE FF  UTF-32, big-endian BOM
			encoding = UTF32
		case bytes.HasPrefix(data, utf16LEBOM) || bytes.HasPrefix(data, utf16BEBOM):
			// FF FE  UTF-16, little endian BOM
			// FE FF  UTF-16, big endian BOM
			encoding = UTF16
		case bytes.HasPrefix(data, ucs43412BOM):
			// FE FF 00 00  UCS-4, unusual octet order BOM (3412)
			encoding = UCS43412
		case bytes.HasPrefix(data, ucs42143BOM):
			// 00 00 FF FE  UCS-4, unusual octet order BOM (2143)
			encoding = UCS42143
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
	if u.inputState == PureAsciiInputState {
		if u.highByteDetector.Match(data) {
			u.inputState = HighByteInputState
		} else if u.inputState == PureAsciiInputState &&
			u.escDetector.Match(bytes.Join([][]byte{u.lastChars, data}, nil)) {
			u.inputState = EcsAsciiInputState
		}
	}

	u.lastChars = append(u.lastChars, data[len(data)-1:]...)

	// If we've seen escape sequences, use the EscCharSetProbe, which
	// uses a simple state machine to check for known escape sequences in
	// HZ and ISO-2022 encodings, since those are the only encodings that
	// use such sequences.
	if u.inputState == EcsAsciiInputState {
		if u.escCharsetProbe == nil {
			u.escCharsetProbe = NewEscCharSetProbe(u.langFilter)
		}
	}
}
