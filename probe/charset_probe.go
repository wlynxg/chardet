package probe

import (
	"bytes"
	"regexp"

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
	pattern := regexp.MustCompile(`[\x00-\x7F]`)
	buf = pattern.ReplaceAll(buf, []byte(" "))
	return buf
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
	pattern := regexp.MustCompile(`[a-zA-Z]*[\x80-\xFF]+[a-zA-Z]*[^a-zA-Z\x80-\xFF]?`)
	words := pattern.FindAll(buf, -1)

	var filtered bytes.Buffer
	for _, word := range words {
		filtered.Write(word[:len(word)-1])

		// If the last character in the word is a marker, replace it with a
		// space as markers shouldn't affect our analysis (they are used
		// similarly across all languages and may thus have similar
		// frequencies).
		lastChar := word[len(word)-1]
		if !util.IsAlpha(lastChar) && lastChar < 0x80 {
			filtered.WriteByte(' ')
		} else {
			filtered.WriteByte(lastChar)
		}
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

		// If current character is not extended-ASCII and not alphabetic...
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
