package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type MacRomanProbe struct {
	CharSetProbe

	Char2Class []int
	ClassModel []int

	lastCharClass int
	freqCounter   []int
}

func NewMacRomanProbe() *MacRomanProbe {
	p := &MacRomanProbe{
		CharSetProbe: NewCharSetProbe(consts.UnknownLangFilter),
		Char2Class: []int{
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 00 - 07
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 08 - 0F
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 10 - 17
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 18 - 1F
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 20 - 27
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 28 - 2F
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 30 - 37
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 38 - 3F
			OTH, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 40 - 47
			ASC, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 48 - 4F
			ASC, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 50 - 57
			ASC, ASC, ASC, OTH, OTH, OTH, OTH, OTH, // 58 - 5F
			OTH, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 60 - 67
			ASS, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 68 - 6F
			ASS, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 70 - 77
			ASS, ASS, ASS, OTH, OTH, OTH, OTH, OTH, // 78 - 7F
			ACV, ACV, ACO, ACV, ACO, ACV, ACV, ASV, // 80 - 87
			ASV, ASV, ASV, ASV, ASV, ASO, ASV, ASV, // 88 - 8F
			ASV, ASV, ASV, ASV, ASV, ASV, ASO, ASV, // 90 - 97
			ASV, ASV, ASV, ASV, ASV, ASV, ASV, ASV, // 98 - 9F
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, ASO, // A0 - A7
			OTH, OTH, ODD, ODD, OTH, OTH, ACV, ACV, // A8 - AF
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // B0 - B7
			OTH, OTH, OTH, OTH, OTH, OTH, ASV, ASV, // B8 - BF
			OTH, OTH, ODD, OTH, ODD, OTH, OTH, OTH, // C0 - C7
			OTH, OTH, OTH, ACV, ACV, ACV, ACV, ASV, // C8 - CF
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, ODD, // D0 - D7
			ASV, ACV, ODD, OTH, OTH, OTH, OTH, OTH, // D8 - DF
			OTH, OTH, OTH, OTH, OTH, ACV, ACV, ACV, // E0 - E7
			ACV, ACV, ACV, ACV, ACV, ACV, ACV, ACV, // E8 - EF
			ODD, ACV, ACV, ACV, ACV, ASV, ODD, ODD, // F0 - F7
			ODD, ODD, ODD, ODD, ODD, ODD, ODD, ODD, // F8 - FF
		},
		ClassModel: []int{
			// UDF OTH ASC ASS ACV ACO ASV ASO ODD
			0, 0, 0, 0, 0, 0, 0, 0, 0, // UDF
			0, 3, 3, 3, 3, 3, 3, 3, 1, // OTH
			0, 3, 3, 3, 3, 3, 3, 3, 1, // ASC
			0, 3, 3, 3, 1, 1, 3, 3, 1, // ASS
			0, 3, 3, 3, 1, 2, 1, 2, 1, // ACV
			0, 3, 3, 3, 3, 3, 3, 3, 1, // ACO
			0, 3, 1, 3, 1, 1, 1, 3, 1, // ASV
			0, 3, 1, 3, 1, 1, 3, 3, 1, // ASO
			0, 1, 1, 1, 1, 1, 1, 1, 1, // ODD
		},

		lastCharClass: OTH,
		freqCounter:   make([]int, FreqCatNum),
	}
	p.Reset()
	return p
}

func (m *MacRomanProbe) Reset() {
	m.lastCharClass = OTH
	m.freqCounter = make([]int, FreqCatNum)

	// express the prior that MacRoman is a somewhat rare encoding;
	// this can be done by starting out in a slightly improbable state
	// that must be overcome
	m.freqCounter[2] = 10

	m.CharSetProbe.Reset()
}

func (m *MacRomanProbe) CharSetName() string {
	return consts.MacRoman
}

func (m *MacRomanProbe) Language() string {
	return ""
}

func (m *MacRomanProbe) Feed(buf []byte) consts.ProbingState {
	buf = m.RemoveXMLTags(buf)
	for _, b := range buf {
		charClass := m.Char2Class[int(b)]
		freq := m.ClassModel[(m.lastCharClass*MacRomanClassNum)+charClass]
		if freq == 0 {
			m.state = consts.NotMeProbingState
			break
		}
		m.freqCounter[freq]++
		m.lastCharClass = charClass
	}
	return m.state
}

func (m *MacRomanProbe) GetConfidence() float64 {
	if m.state == consts.NotMeProbingState {
		return 0.01
	}

	total := 0
	for i := 0; i < len(m.freqCounter); i++ {
		total += m.freqCounter[i]
	}

	confidence := (float64(m.freqCounter[3]) - float64(m.freqCounter[1])*20.0) / float64(total)
	if float64(total) < 0.01 {
		confidence = 0.0
	}
	confidence = max(confidence, 0.0) * 0.73
	return confidence
}
