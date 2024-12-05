package chardet

type EscCharSetProbe struct {
	filter           LangFilter
	state            ProbingState
	activeSmCount    int
	detectedCharset  string
	detectedLanguage string

	codingSM []*CodingStateMachine
}

func NewEscCharSetProbe(filter LangFilter) *EscCharSetProbe {
	probe := &EscCharSetProbe{
		filter:           filter,
		activeSmCount:    0,
		detectedCharset:  "",
		detectedLanguage: "",
		state:            DetectingProbingState,
		codingSM:         []*CodingStateMachine{},
	}

	if probe.filter&ChineseLangFilter != 0 {
		probe.codingSM = append(probe.codingSM, NewCodingStateMachine(*HzSmModel()),
			NewCodingStateMachine(*Iso2022cnSmModel()))
	}

	if probe.filter&JapaneseLangFilter != 0 {
		probe.codingSM = append(probe.codingSM, NewCodingStateMachine(*Iso2022jpSmModel()))
	}

	if probe.filter&KoreanLangFilter != 0 {
		probe.codingSM = append(probe.codingSM, NewCodingStateMachine(*Iso2022krSmModel()))
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

	e.state = DetectingProbingState
	e.activeSmCount = len(e.codingSM)
	e.detectedCharset = ""
	e.detectedLanguage = ""
}

func (e *EscCharSetProbe) Feed(data []byte) ProbingState {
	for _, b := range data {
		for _, machine := range e.codingSM {
			if machine == nil || !machine.Active {
				continue
			}

			codingState := machine.NextState(b)

			switch codingState {
			case ErrorMachineState:
				machine.Active = false
				e.activeSmCount--
				if e.activeSmCount <= 0 {
					e.state = NotMeProbingState
					return e.state
				}
			case ItsMeMachineState:
				e.state = FoundItProbingState
				e.detectedCharset = machine.CodingStateMachine()
				e.detectedLanguage = machine.Language()
			default:
			}
		}
	}
	return e.state
}

func (e *EscCharSetProbe) CharsetName() string {
	return e.detectedCharset
}

func (e *EscCharSetProbe) Language() string {
	return e.detectedLanguage
}

func (e *EscCharSetProbe) Confidence() float32 {
	if e.detectedCharset != "" {
		return 0.99
	} else {
		return 0.0
	}
}
