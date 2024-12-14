package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type EscCharSetProbe struct {
	CharSetProbe

	activeSmCount    int
	detectedCharset  string
	detectedLanguage string

	codingSM []*CodingStateMachine
}

func NewEscCharSetProbe(filter consts.LangFilter) *EscCharSetProbe {
	probe := &EscCharSetProbe{
		CharSetProbe:     NewCharSetProbe(filter),
		activeSmCount:    0,
		detectedCharset:  "",
		detectedLanguage: "",
		codingSM:         []*CodingStateMachine{},
	}

	if probe.filter&consts.ChineseLangFilter != 0 {
		probe.codingSM = append(probe.codingSM,
			NewCodingStateMachine(HzSmModel()),
			NewCodingStateMachine(Iso2022cnSmModel()))
	}

	if probe.filter&consts.JapaneseLangFilter != 0 {
		probe.codingSM = append(probe.codingSM, NewCodingStateMachine(Iso2022jpSmModel()))
	}

	if probe.filter&consts.KoreanLangFilter != 0 {
		probe.codingSM = append(probe.codingSM, NewCodingStateMachine(Iso2022krSmModel()))
	}

	probe.Reset()
	return probe
}

func (e *EscCharSetProbe) Reset() {
	for _, model := range e.codingSM {
		if model == nil {
			continue
		}
		model.Active = true
		model.Reset()
	}

	e.state = consts.DetectingProbingState
	e.activeSmCount = len(e.codingSM)
	e.detectedCharset = ""
	e.detectedLanguage = ""
}

func (e *EscCharSetProbe) GetConfidence() float64 {
	if e.detectedCharset != "" {
		return 0.99
	} else {
		return 0.0
	}
}

func (e *EscCharSetProbe) CharSetName() string {
	return e.detectedCharset
}

func (e *EscCharSetProbe) Language() string {
	return e.detectedLanguage
}

func (e *EscCharSetProbe) Feed(buf []byte) consts.ProbingState {
	for _, b := range buf {
		for _, machine := range e.codingSM {
			if machine == nil || !machine.Active {
				continue
			}

			codingState := machine.NextState(b)

			switch codingState {
			case consts.ErrorMachineState:
				machine.Active = false
				e.activeSmCount--
				if e.activeSmCount <= 0 {
					e.state = consts.NotMeProbingState
					return e.state
				}
			case consts.ItsMeMachineState:
				e.state = consts.FoundItProbingState
				e.detectedCharset = machine.CodingStateMachine()
				e.detectedLanguage = machine.Language()
			default:
			}
		}
	}
	return e.state
}
