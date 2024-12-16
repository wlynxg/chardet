package probe

import (
	"bytes"

	"github.com/wlynxg/chardet/consts"
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

This filter applies to all scripts that do not use English characters.
*/
func (p *CharSetProbe) FilterInternationalWords(buf []byte) []byte {
	var (
		filtered         = make([]byte, 0, len(buf))
		word             = make([]byte, 0, 16)
		hasInternational bool
	)

	for i := 0; i < len(buf); i++ {
		b := buf[i]

		// Check if the byte is an English alphabet
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			word = append(word, b)
		} else if b >= 0x80 { // Check if the byte is an international character
			word = append(word, b)
			hasInternational = true
		} else { // It's a marker character
			if hasInternational {
				filtered = append(filtered, word...)
				// Replace the last character with a space if it's a marker
				if lastChar := word[len(word)-1]; (lastChar < 'a' || lastChar > 'z') &&
					(lastChar < 'A' || lastChar > 'Z') {
					filtered = append(filtered, ' ')
				} else {
					filtered = append(filtered, lastChar)
				}
			}
			word = word[:0]          // Reset the word buffer
			hasInternational = false // Reset international flag
		}
	}

	// Handle the last word if it was not followed by a marker
	if hasInternational && len(word) > 0 {
		filtered = append(filtered, word...)
	}

	return filtered
}

/*
FilterWithEnglishLetters :
Returns a copy of "buf" that retains only the sequences of English
alphabet and high byte characters that are not between <> characters.
Also retains English alphabet and high byte characters immediately
before occurrences of >.

This filter can be applied to all scripts that contain both English
characters and extended ASCII characters, but is currently only used by
"Latin1Probe".
*/
func (p *CharSetProbe) FilterWithEnglishLetters(buf []byte) []byte {
	// Pre-allocate a buffer based on an estimate of the filtered content size
	filtered := make([]byte, 0, len(buf))
	inTag := false
	prev := 0

	for curr := 0; curr < len(buf); curr++ {
		bufChar := buf[curr]
		// Check if we're coming out of or entering an HTML tag
		if bufChar == '>' {
			inTag = false
		} else if bufChar == '<' {
			inTag = true
		}

		// Inline the check for alphabetic characters
		if bufChar < 0x80 && !((bufChar >= 'a' && bufChar <= 'z') || (bufChar >= 'A' && bufChar <= 'Z')) {
			// If we're not in a tag, and we've found some text to keep
			if curr > prev && !inTag {
				// Append the slice from prev to curr
				filtered = append(filtered, buf[prev:curr]...)
				filtered = append(filtered, ' ')
			}
			prev = curr + 1
		}
	}

	// If we're not in a tag, and we've got some remaining text to keep
	if !inTag && prev < len(buf) {
		filtered = append(filtered, buf[prev:]...)
	}

	// If no filtering occurred, return the original buffer
	if len(filtered) == 0 {
		return buf
	}

	return filtered
}

// RemoveXMLTags removes XML tags from the input buffer and retains only the sequences
// of English alphabet and high byte characters that are not between <> characters.
func (p *CharSetProbe) RemoveXMLTags(buf []byte) []byte {
	// Pre-allocate a buffer based on an estimate of the filtered content size
	filtered := make([]byte, 0, len(buf))
	inTag := false
	prev := 0

	for curr := 0; curr < len(buf); curr++ {
		bufChar := buf[curr]

		// Check for end of tag
		if bufChar == '>' {
			inTag = false
			prev = curr + 1
			continue
		}

		// Check to begin tag
		if bufChar == '<' {
			if curr > prev && !inTag {
				// Append the slice from prev to curr
				filtered = append(filtered, buf[prev:curr]...)
				filtered = append(filtered, ' ')
			}
			inTag = true
		}
	}

	// If we're not in a tag, and we've got some remaining text to keep
	if !inTag && prev < len(buf) {
		filtered = append(filtered, buf[prev:]...)
	}

	// If no filtering occurred, return the original buffer
	if len(filtered) == 0 {
		return buf
	}

	return filtered
}
