package probe

import (
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/log"
	"go.uber.org/zap"
)

type SingleByteCharSetModel struct {
	CharsetName          string
	Language             string
	CharToOrderMap       [256]int
	LanguageModel        map[int]map[int]int
	TypicalPositiveRatio float64
	KeepAsciiLetters     bool
	Alphabet             string
}

type SingleByteCharSetProbe struct {
	CharSetProbe

	SampleSize, SBEnoughRelThreshold                     int
	PositiveShortcutThreshold, NegativeShortcutThreshold float64

	log         *zap.SugaredLogger
	model       *SingleByteCharSetModel
	reversed    bool
	nameProbe   Probe
	lastOrder   int
	seqCounters []int
	totalSeqs   int
	totalChar   int
	controlChar int
	freqChar    int
}

func NewSingleByteCharSetProbe(model *SingleByteCharSetModel, reversed bool, nameProbe Probe) *SingleByteCharSetProbe {
	sp := &SingleByteCharSetProbe{
		CharSetProbe:              NewCharSetProbe(consts.UnknownLangFilter),
		SampleSize:                64,
		SBEnoughRelThreshold:      1024, // 0.25 * SampleSize^2
		PositiveShortcutThreshold: 0.95,
		NegativeShortcutThreshold: 0.05,
		log:                       log.New("SingleByteCharSetProbe"),
		model:                     model,
		// TRUE if we need to reverse every pair in the model lookup
		reversed: reversed,
		// Optional auxiliary probe for a name decision
		nameProbe:   nameProbe,
		lastOrder:   255,
		seqCounters: make([]int, consts.LikelihoodCategories),
		totalSeqs:   0,
		totalChar:   0,
		controlChar: 0,
		freqChar:    0,
	}
	sp.Reset()
	return sp
}

func (s *SingleByteCharSetProbe) Reset() {
	s.CharSetProbe.Reset()
	// char order of last character
	s.lastOrder = 255
	s.seqCounters = make([]int, consts.LikelihoodCategories)
	s.totalSeqs = 0
	s.totalChar = 0
	s.controlChar = 0
	// characters that fall in our sampling range
	s.freqChar = 0
}

func (s *SingleByteCharSetProbe) CharSetName() string {
	if s.nameProbe != nil {
		return s.nameProbe.CharSetName()
	} else {
		return s.model.CharsetName
	}
}

func (s *SingleByteCharSetProbe) Language() string {
	if s.nameProbe != nil {
		return s.nameProbe.Language()
	} else {
		return s.model.Language
	}
}

func (s *SingleByteCharSetProbe) Feed(buf []byte) consts.ProbingState {
	if !s.model.KeepAsciiLetters {
		buf = s.FilterInternationalWords(buf)
	} else {
		buf = s.RemoveXMLTags(buf)
	}

	if len(buf) == 0 {
		return s.state
	}

	for _, b := range buf {
		order := s.model.CharToOrderMap[b]
		// XXX: This was SYMBOL_CAT_ORDER before, with a value of 250, but
		//      CharacterCategory.SYMBOL is actually 253, so we use CONTROL
		//      to make it closer to the original intent. The only difference
		//      is whether we count digits and control characters for
		//      _total_char purposes.
		if order < int(consts.ControlCharacterCategory) {
			s.totalChar++
		}

		if order < s.SampleSize {
			s.freqChar++
			if s.lastOrder < s.SampleSize {
				s.totalSeqs++

				lmCat := s.model.LanguageModel[s.lastOrder][order]
				if s.reversed {
					lmCat = s.model.LanguageModel[order][s.lastOrder]
				}
				s.seqCounters[lmCat]++
			}
		}
		s.lastOrder = order
	}

	charsetName := s.model.CharsetName
	if s.state == consts.DetectingProbingState {
		if s.totalSeqs > s.SBEnoughRelThreshold {
			confidence := s.GetConfidence()
			if confidence > s.PositiveShortcutThreshold {
				s.log.Debugf("%s confidence = %f, we have a winner", charsetName, confidence)
				s.state = consts.FoundItProbingState
			} else if confidence < s.NegativeShortcutThreshold {
				s.log.Debugf("%s confidence = %f, below negative shortcut threshhold %f",
					charsetName, confidence, s.NegativeShortcutThreshold)
				s.state = consts.NotMeProbingState
			}
		}
	}

	return s.state
}

func (s *SingleByteCharSetProbe) GetConfidence() float64 {
	r := 0.01
	if s.totalSeqs > 0 {
		r1 := (float64(s.seqCounters[int(consts.PositiveSequenceLikelihood)]) +
			0.25*float64(s.seqCounters[consts.LikelySequenceLikelihood])) /
			float64(s.totalSeqs) / s.model.TypicalPositiveRatio
		r1 = r1 * float64(s.totalChar-s.controlChar) / float64(s.totalChar)
		r = r1 * float64(s.freqChar) / float64(s.totalChar)
		if r >= 1 {
			r = 0.99
		}
	}
	return r
}
