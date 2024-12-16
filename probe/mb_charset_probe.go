package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
)

type MultiByteCharSetProbe struct {
	CharSetProbe

	charsetName, language string
	distributionAnalyzer  cda.Analyzer
	codingSM              *CodingStateMachine
	lastChar              [2]byte
}

func NewMultiByteCharSetProbe(charsetName, language string, filter consts.LangFilter,
	distributionAnalyzer cda.Analyzer, codingSM *CodingStateMachine) MultiByteCharSetProbe {
	return MultiByteCharSetProbe{
		CharSetProbe:         NewCharSetProbe(filter),
		charsetName:          charsetName,
		language:             language,
		distributionAnalyzer: distributionAnalyzer,
		codingSM:             codingSM,
		lastChar:             [2]byte{0, 0},
	}
}

func (m *MultiByteCharSetProbe) Reset() {
	m.CharSetProbe.Reset()
	if m.codingSM != nil {
		m.codingSM.Reset()
	}
	if m.distributionAnalyzer != nil {
		m.distributionAnalyzer.Reset()
	}
	m.lastChar = [2]byte{0, 0}
}

func (m *MultiByteCharSetProbe) Feed(buf []byte) consts.ProbingState {
	if len(buf) == 0 {
		return m.state
	}

loop:
	for i := 0; i < len(buf); i++ {
		codingState := m.codingSM.NextState(buf[i])
		switch codingState {
		case consts.ErrorMachineState:
			m.state = consts.NotMeProbingState
			break loop
		case consts.ItsMeMachineState:
			m.state = consts.FoundItProbingState
			break loop
		case consts.StartMachineState:
			charLen := m.codingSM.CurrentCharLength()
			if i == 0 {
				m.lastChar[1] = buf[0]
				m.distributionAnalyzer.Feed(m.lastChar[:], charLen)
			} else {
				m.distributionAnalyzer.Feed(buf[i-1:i+1], charLen)
			}
		default:
		}
	}

	m.lastChar[0] = buf[len(buf)-1]

	if m.state == consts.DetectingProbingState {
		if m.distributionAnalyzer.GotEnoughData() && m.GetConfidence() > m.ShortcutThreshold {
			m.state = consts.FoundItProbingState
		}
	}
	return m.state
}

func (m *MultiByteCharSetProbe) GetConfidence() float64 {
	return m.distributionAnalyzer.GetConfidence()
}

func (m *MultiByteCharSetProbe) CharSetName() string {
	return m.charsetName
}

func (m *MultiByteCharSetProbe) Language() string {
	return m.language
}
