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
		state:            StartMachineState,
		activeSmCount:    0,
		detectedCharset:  "",
		detectedLanguage: "",
		codingSM:         []*CodingStateMachine{},
	}

	if probe.filter&ChineseLangFilter != 0 {
		probe.codingSM = append(probe.codingSM, NewCodingStateMachine(HzSmModel()),
			NewCodingStateMachine(Iso2022cnSmModel()))
	}

	if probe.filter&JapaneseLangFilter != 0 {
		probe.codingSM = append(probe.codingSM, NewCodingStateMachine(Iso2022jpSmModel()))
	}

	if probe.filter&KoreanLangFilter != 0 {
		probe.codingSM = append(probe.codingSM, NewCodingStateMachine(Iso2022krSmModel()))
	}

	return probe
}

func (e *EscCharSetProbe) Reset() {
	for _, model := range e.codingSM {
		if model == nil {
			continue
		}
		model.Reset()
	}

	e.state = DetectingProbingState
	e.activeSmCount = len(e.codingSM)
	e.detectedCharset = ""
	e.detectedLanguage = ""
}
