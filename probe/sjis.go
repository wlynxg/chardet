package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
)

type SJISProbe struct {
	MultiByteCharSetProbe

	state           consts.ProbingState
	contextAnalyzer cda.Analyzer
}

func NewSJISProbe() *SJISProbe {
	return &SJISProbe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			"",
			consts.Japanese,
			consts.UnknownLangFilter,
			cda.NewSJISDistributionAnalysis(),
			NewCodingStateMachine(SjisSmModel()),
		),
		contextAnalyzer: cda.NewSJISContextAnalysis(),
	}
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

func (s *SJISProbe) Feed(buf []byte) consts.ProbingState {
loop:
	for i := 0; i < len(buf); i++ {
		codingState := s.codingSM.NextState(buf[i])
		switch codingState {
		case consts.ErrorMachineState:
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

	s.lastChar[0] = buf[1]
	if s.state == consts.DetectingProbingState &&
		s.contextAnalyzer.GotEnoughData() &&
		(s.GetConfidence() > s.ShortcutThreshold) {
		s.state = consts.FoundItProbingState
	}
	return s.state
}

func (s *SJISProbe) GetConfidence() float64 {
	return max(s.contextAnalyzer.GetConfidence(), s.distributionAnalyzer.GetConfidence())
}
