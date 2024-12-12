package probe

import (
	"bytes"

	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/util"
)

type CharSetProbe struct {
	ShortcutThreshold float64

	active bool
	state  consts.ProbingState
	filter consts.LangFilter
}

func NewCharSetProbe(filter consts.LangFilter) CharSetProbe {
	csp := CharSetProbe{
		ShortcutThreshold: 0.95,

		active: false,
		state:  consts.DetectingProbingState,
		filter: filter,
	}
	return csp
}

func (p *CharSetProbe) Reset() {
	p.state = consts.DetectingProbingState
}

func (p *CharSetProbe) SetActive(state bool) {
	p.active = state
}

func (p *CharSetProbe) IsActive() bool {
	return p.active
}

func (p *CharSetProbe) State() consts.ProbingState {
	return p.state
}

func (p *CharSetProbe) FilterHighByteOnly(buf []byte) []byte {
	var result bytes.Buffer

	inASCII := false // Track if we are currently in a sequence of ASCII characters

	for _, b := range buf {
		if b <= 0x7F { // Check if the byte is ASCII
			if !inASCII {
				// If we are not already in an ASCII sequence, add a space
				result.WriteByte(' ')
				inASCII = true // Set the flag to true
			}
		} else {
			// If we encounter a non-ASCII character, reset the flag
			inASCII = false
			result.WriteByte(b) // Write the non-ASCII byte directly
		}
	}

	return result.Bytes()
}

/*
FilterInternationalWords :

We define three types of bytes:
alphabet: english alphabets [a-zA-Z]
international: international characters [\x80-\xFF]
marker: everything else [^a-zA-Z\x80-\xFF]

The input buffer can be thought to contain a series of words delimited
by markers. This function works to filter all words that contain at
least one international character. All contiguous sequences of markers
are replaced by a single space ascii character.

This filter applies to all scripts which do not use English characters.
*/
func (p *CharSetProbe) FilterInternationalWords(buf []byte) []byte {
	var filtered bytes.Buffer

	var word bytes.Buffer
	hasInternational := false

	for i := 0; i < len(buf); i++ {
		b := buf[i]

		// Check if the byte is an English alphabet
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			word.WriteByte(b)
		} else if b >= 0x80 { // Check if the byte is an international character
			word.WriteByte(b)
			hasInternational = true
		} else { // It's a marker character
			if hasInternational && word.Len() > 0 {
				filtered.Write(word.Bytes())
				// Replace the last character with a space if it's a marker
				if word.Len() > 0 {
					lastChar := word.Bytes()[word.Len()-1]
					if (lastChar < 'a' || lastChar > 'z') && (lastChar < 'A' || lastChar > 'Z') {
						filtered.WriteByte(' ')
					} else {
						filtered.WriteByte(lastChar)
					}
				}
			}
			word.Reset()             // Reset the word buffer
			hasInternational = false // Reset international flag
		}
	}

	// Handle the last word if it was not followed by a marker
	if hasInternational && word.Len() > 0 {
		filtered.Write(word.Bytes())
	}

	return filtered.Bytes()
}

/*
FilterWithEnglishLetters :
Returns a copy of "buf" that retains only the sequences of English
alphabet and high byte characters that are not between <> characters.
Also retains English alphabet and high byte characters immediately
before occurrences of >.

This filter can be applied to all scripts which contain both English
characters and extended ASCII characters, but is currently only used by
"Latin1Probe".
*/
func (p *CharSetProbe) FilterWithEnglishLetters(buf []byte) []byte {
	var (
		filtered bytes.Buffer
		inTag    bool
		prev     int
	)

	for curr := 0; curr < len(buf); curr++ {
		bufChar := buf[curr]
		// Check if we're coming out of or entering an HTML tag
		if bufChar == '>' {
			inTag = false
		} else if bufChar == '<' {
			inTag = true
		}

		// If the current character is not extended-ASCII and not alphabetic...
		if bufChar < 0x80 && !util.IsAlpha(bufChar) {
			// ...and we're not in a tag
			if curr > prev && !inTag {
				// Keep everything after last non-extended-ASCII, non-alphabetic character
				filtered.Write(buf[prev:curr])
				// Output a space to delimit stretch we kept
				filtered.WriteByte(' ')
			}
			prev = curr + 1
		}
	}

	// If we're not in a tag...
	if !inTag {
		// Keep everything after last non-extended-ASCII, non-alphabetic character
		filtered.Write(buf[prev:])
	}

	return filtered.Bytes()
}

// RemoveXMLTags removes XML tags from the input buffer and retains only the sequences
// of English alphabet and high byte characters that are not between <> characters.
func (p *CharSetProbe) RemoveXMLTags(buf []byte) []byte {
	var filtered bytes.Buffer
	inTag := false
	prev := 0

	for curr := 0; curr < len(buf); curr++ {
		bufChar := buf[curr]

		if bufChar == '>' { // End of a tag
			prev = curr + 1
			inTag = false
		} else if bufChar == '<' { // Start of a tag
			if curr > prev && !inTag {
				// Keep everything after last non-extended-ASCII, non-alphabetic character
				filtered.Write(buf[prev:curr])
				// Output a space to delimit stretch we kept
				filtered.WriteByte(' ')
			}
			inTag = true
		}
	}

	// If we're not in a tag, keep everything after last non-extended-ASCII, non-alphabetic character
	if !inTag {
		filtered.Write(buf[prev:])
	}

	return filtered.Bytes()
}
