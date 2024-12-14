package probe

import (
	"github.com/wlynxg/chardet/consts"
)

const (
	FreqCatNum = 4

	UDF              = 0 // undefined
	OTH              = 1 // other
	ASC              = 2 // ascii capital letter
	ASS              = 3 // ascii small letter
	ACV              = 4 // accent capital vowel
	ACO              = 5 // accent capital other
	ASV              = 6 // accent small vowel
	ASO              = 7 // accent small other
	ODD              = 8 // character that is unlikely to appear
	Latin1ClassNum   = 8 // latin1 total classes
	MacRomanClassNum = 9 // macroman total classes
)

type Latin1Probe struct {
	CharSetProbe

	Char2Class []int
	ClassModel []int

	lastCharClass int
	freqCounter   []int
}

func NewLatin1Probe() *Latin1Probe {
	return &Latin1Probe{
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
			OTH, UDF, OTH, ASO, OTH, OTH, OTH, OTH, // 80 - 87
			OTH, OTH, ACO, OTH, ACO, UDF, ACO, UDF, // 88 - 8F
			UDF, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 90 - 97
			OTH, OTH, ASO, OTH, ASO, UDF, ASO, ACO, // 98 - 9F
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // A0 - A7
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // A8 - AF
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // B0 - B7
			OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // B8 - BF
			ACV, ACV, ACV, ACV, ACV, ACV, ACO, ACO, // C0 - C7
			ACV, ACV, ACV, ACV, ACV, ACV, ACV, ACV, // C8 - CF
			ACO, ACO, ACV, ACV, ACV, ACV, ACV, OTH, // D0 - D7
			ACV, ACV, ACV, ACV, ACV, ACO, ACO, ACO, // D8 - DF
			ASV, ASV, ASV, ASV, ASV, ASV, ASO, ASO, // E0 - E7
			ASV, ASV, ASV, ASV, ASV, ASV, ASV, ASV, // E8 - EF
			ASO, ASO, ASV, ASV, ASV, ASV, ASV, OTH, // F0 - F7
			ASV, ASV, ASV, ASV, ASV, ASO, ASO, ASO, // F8 - FF
		},
		// 0: illegal
		// 1: very unlikely
		// 2: normal
		// 3: very likely
		ClassModel: []int{
			// UDF OTH ASC ASS ACV ACO ASV ASO
			0, 0, 0, 0, 0, 0, 0, 0, // UDF
			0, 3, 3, 3, 3, 3, 3, 3, // OTH
			0, 3, 3, 3, 3, 3, 3, 3, // ASC
			0, 3, 3, 3, 1, 1, 3, 3, // ASS
			0, 3, 3, 3, 1, 2, 1, 2, // ACV
			0, 3, 3, 3, 3, 3, 3, 3, // ACO
			0, 3, 1, 3, 1, 1, 1, 3, // ASV
			0, 3, 1, 3, 1, 1, 3, 3, // ASO
		},

		lastCharClass: OTH,
		freqCounter:   make([]int, FreqCatNum),
	}
}

func (l *Latin1Probe) Reset() {
	l.lastCharClass = OTH
	l.freqCounter = make([]int, FreqCatNum)
	l.CharSetProbe.Reset()
}

func (l *Latin1Probe) CharSetName() string {
	return consts.ISO88591
}

func (l *Latin1Probe) Language() string {
	return ""
}

func (l *Latin1Probe) Feed(buf []byte) consts.ProbingState {
	buf = l.FilterWithEnglishLetters(buf)
	for _, b := range buf {
		charCls := l.Char2Class[int(b)]
		freq := l.ClassModel[(l.lastCharClass*Latin1ClassNum)+charCls]
		if freq == 0 {
			l.state = consts.NotMeProbingState
			break
		}
		l.freqCounter[freq]++
		l.lastCharClass = charCls
	}
	return l.state
}

func (l *Latin1Probe) GetConfidence() float64 {
	if l.state == consts.NotMeProbingState {
		return 0.01
	}
	total := 0
	for _, freq := range l.freqCounter {
		total += freq
	}

	confidence := 0.0
	if float64(total) >= 0.01 {
		confidence = float64(l.freqCounter[3]-l.freqCounter[1]*20) / float64(total)
	}
	if confidence < 0.0 {
		confidence = 0.0
	}
	// lower the confidence of latin1 so that another more accurate
	// detector can take priority.
	confidence = confidence * 0.73
	return confidence
}
