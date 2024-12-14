package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type UTF1632Probe struct {
	CharSetProbe

	// how many logical characters to scan before feeling confident of prediction
	MinCharsForDetection float64
	// a fixed constant ratio of expected zeros or non-zeros in modulo-position.
	ExpectedRatio float64

	position     int
	zerosAtMod   [4]float64
	nonzeroAtMod [4]float64
	quad         [4]byte

	invalidUtf16be, invalidUtf16le     bool
	invalidUtf32be, invalidUtf32le     bool
	firstHalfSurrogatePairDetected16be bool
	firstHalfSurrogatePairDetected16le bool
}

func NewUTF1632Probe() *UTF1632Probe {
	p := &UTF1632Probe{
		CharSetProbe: NewCharSetProbe(consts.UnknownLangFilter),

		MinCharsForDetection: 4,
		ExpectedRatio:        0.94,

		position:                           0,
		zerosAtMod:                         [4]float64{},
		nonzeroAtMod:                       [4]float64{},
		quad:                               [4]byte{},
		invalidUtf16be:                     false,
		invalidUtf16le:                     false,
		invalidUtf32be:                     false,
		invalidUtf32le:                     false,
		firstHalfSurrogatePairDetected16be: false,
		firstHalfSurrogatePairDetected16le: false,
	}
	return p
}

func (u *UTF1632Probe) Reset() {
	u.CharSetProbe.Reset()
	u.position = 0
	u.zerosAtMod = [4]float64{}
	u.nonzeroAtMod = [4]float64{}
	u.quad = [4]byte{}
	u.invalidUtf16be = false
	u.invalidUtf16le = false
	u.invalidUtf32be = false
	u.invalidUtf32le = false
	u.firstHalfSurrogatePairDetected16be = false
	u.firstHalfSurrogatePairDetected16le = false
}

func (u *UTF1632Probe) CharSetName() string {
	switch {
	case u.isLikelyUtf32be():
		return consts.UTF32Be
	case u.isLikelyUtf32le():
		return consts.UTF32Le
	case u.isLikelyUtf16be():
		return consts.UTF16Be
	case u.isLikelyUtf16le():
		return consts.UTF16Le
	default:
		return consts.UTF16
	}
}

func (u *UTF1632Probe) Language() string {
	return ""
}

func (u *UTF1632Probe) approx32bitChars() float64 {
	return max(1.0, float64(u.position)/4.0)
}

func (u *UTF1632Probe) approx16bitChars() float64 {
	return max(1.0, float64(u.position)/2.0)
}

func (u *UTF1632Probe) isLikelyUtf32be() bool {
	approxChars := u.approx32bitChars()
	return approxChars >= u.MinCharsForDetection &&
		(u.zerosAtMod[0]/approxChars > u.ExpectedRatio &&
			u.zerosAtMod[1]/approxChars > u.ExpectedRatio &&
			u.zerosAtMod[2]/approxChars > u.ExpectedRatio &&
			u.nonzeroAtMod[3]/approxChars > u.ExpectedRatio &&
			!u.invalidUtf32be)
}

func (u *UTF1632Probe) isLikelyUtf32le() bool {
	approxChars := u.approx32bitChars()
	return approxChars >= u.MinCharsForDetection &&
		(u.nonzeroAtMod[0]/approxChars > u.ExpectedRatio &&
			u.zerosAtMod[1]/approxChars > u.ExpectedRatio &&
			u.zerosAtMod[2]/approxChars > u.ExpectedRatio &&
			u.zerosAtMod[3]/approxChars > u.ExpectedRatio &&
			!u.invalidUtf32le)
}

func (u *UTF1632Probe) isLikelyUtf16be() bool {
	approxChars := u.approx16bitChars()
	return approxChars >= u.MinCharsForDetection &&
		((u.nonzeroAtMod[1]+u.nonzeroAtMod[3])/approxChars > u.ExpectedRatio &&
			(u.zerosAtMod[0]+u.zerosAtMod[2])/approxChars > u.ExpectedRatio &&
			!u.invalidUtf16be)
}

func (u *UTF1632Probe) isLikelyUtf16le() bool {
	approxChars := u.approx16bitChars()
	return approxChars >= u.MinCharsForDetection &&
		((u.nonzeroAtMod[0]+u.nonzeroAtMod[2])/approxChars > u.ExpectedRatio &&
			(u.zerosAtMod[1]+u.zerosAtMod[3])/approxChars > u.ExpectedRatio &&
			!u.invalidUtf16le)
}

// Validate if the quad of bytes is valid UTF-32.
//
//	UTF-32 is valid in the range 0x00000000 - 0x0010FFFF
//	excluding 0x0000D800 - 0x0000DFFF
//
//	https://en.wikipedia.org/wiki/UTF-32
func (u *UTF1632Probe) validateUtf32Characters(quard []byte) {
	if quard[0] != 0 || quard[1] > 0x10 || (quard[0] == 0 && quard[1] == 0 && quard[2] >= 0xD8 && quard[2] <= 0xDF) {
		u.invalidUtf32be = true
	}

	if quard[3] != 0 || quard[2] > 0x10 || (quard[3] == 0 && quard[2] == 0 && quard[1] >= 0xD8 && quard[1] <= 0xDF) {
		u.invalidUtf32le = true
	}
}

// Validate if the pair of bytes is  valid UTF-16.
//
//	UTF-16 is valid in the range 0x0000 - 0xFFFF excluding 0xD800 - 0xFFFF
//	with an exception for surrogate pairs, which must be in the range
//	0xD800-0xDBFF followed by 0xDC00-0xDFFF
//
//	https://en.wikipedia.org/wiki/UTF-16
func (u *UTF1632Probe) validateUtf16Characters(quard []byte) {
	if !u.firstHalfSurrogatePairDetected16be {
		if quard[0] >= 0xD8 && quard[0] <= 0xDB {
			u.firstHalfSurrogatePairDetected16be = true
		} else if quard[0] >= 0xDC && quard[0] <= 0xDF {
			u.invalidUtf16be = true
		}
	} else {
		if quard[0] >= 0xDC && quard[0] <= 0xDF {
			u.firstHalfSurrogatePairDetected16be = false
		} else {
			u.invalidUtf16le = true
		}
	}

	if !u.firstHalfSurrogatePairDetected16le {
		if quard[1] >= 0xD8 && quard[1] <= 0xDB {
			u.firstHalfSurrogatePairDetected16le = true
		} else if quard[1] >= 0xDC && quard[1] <= 0xDF {
			u.invalidUtf16le = true
		}
	} else {
		if quard[1] >= 0xDC && quard[1] <= 0xDF {
			u.firstHalfSurrogatePairDetected16le = false
		} else {
			u.invalidUtf16le = true
		}
	}
}

func (u *UTF1632Probe) Feed(buf []byte) consts.ProbingState {
	for _, b := range buf {
		mod4 := u.position % 4
		u.quad[mod4] = b
		if mod4 == 3 {
			u.validateUtf32Characters(u.quad[:])
			u.validateUtf16Characters(u.quad[:2])
			u.validateUtf16Characters(u.quad[2:4])
		}

		if b == 0 {
			u.zerosAtMod[mod4]++
		} else {
			u.nonzeroAtMod[mod4]++
		}
		u.position++
	}
	return u.State()
}

func (u *UTF1632Probe) State() consts.ProbingState {
	if u.state == consts.NotMeProbingState || u.state == consts.FoundItProbingState {
		// terminal, decided states
		return u.state
	}

	if u.GetConfidence() > 0.80 {
		u.state = consts.FoundItProbingState
	} else if u.position > 4*1024 {
		// if we get to 4kb into the file, and we can't conclude it's UTF, let's give up
		u.state = consts.NotMeProbingState
	}
	return u.state
}

func (u *UTF1632Probe) GetConfidence() float64 {
	if u.isLikelyUtf16be() || u.isLikelyUtf16le() || u.isLikelyUtf32be() || u.isLikelyUtf32le() {
		return 0.85
	}
	return 0.0
}
