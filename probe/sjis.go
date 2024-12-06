package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/log"
	"github.com/wlynxg/chardet/smm"
	"go.uber.org/zap"
)

type SJISProbe struct {
	MultiByteCharSetProbe

	log                  *zap.SugaredLogger
	state                consts.ProbingState
	codingSM             *CodingStateMachine
	distributionAnalyzer cda.Analyzer
	contextAnalyzer      cda.Analyzer
	numMbChars           int
}

func NewSJISProbe() *SJISProbe {
	sp := &SJISProbe{
		log:                  log.New("SJISProbe"),
		codingSM:             NewCodingStateMachine(smm.SjisSmModel()),
		distributionAnalyzer: cda.NewSJISDistributionAnalysis(),
		contextAnalyzer:      cda.NewSJISContextAnalysis(),
	}
	sp.MultiByteCharSetProbe = NewMultiByteCharSetProbe(
		sp.contextAnalyzer.CharSetName(),
		consts.JapaneseLanguage,
		consts.UnknownLangFilter)
	return sp
}

func (s *SJISProbe) Reset() {
	s.MultiByteCharSetProbe.Reset()
	if s.contextAnalyzer != nil {
		s.contextAnalyzer.Reset()
	}
}

func (s *SJISProbe) CharSetName() string {
	return s.contextAnalyzer.CharSetName()
}

func (s *SJISProbe) Language() string {
	return consts.JapaneseLanguage
}

func (s *SJISProbe) Feed(buf []byte) consts.ProbingState {
loop:
	for i := 0; i < len(buf); i++ {
		codingState := s.codingSm.NextState(buf[i])
		switch codingState {
		case consts.ErrorMachineState:
			s.log.Debugf("%s %s prober hit error at byte %d", s.charsetName, s.language, i)
			s.state = consts.NotMeProbingState
			break loop
		case consts.ItsMeMachineState:
			s.state = consts.FoundItProbingState
			break loop
		case consts.StartMachineState:
			charLen := s.codingSM.CurrentCharLength()
			if i == 0 {
				s.lastChar[1] = buf[0]
				s.contextAnalyzer.Feed(s.lastChar[2-charLen:], charLen)
				s.distributionAnalyzer.Feed(s.lastChar[:], charLen)
			} else {
				s.contextAnalyzer.Feed(buf[i+1-charLen:i+3-charLen], charLen)
				s.distributionAnalyzer.Feed(buf[i-1:i+1], charLen)
			}
		default:
		}
	}

	s.
	return s.state
}
