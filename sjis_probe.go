package chardet

type SJISProbe struct {
	state                ProbingState
	codingSM             *CodingStateMachine
	distributionAnalyzer interface{}
	numMbChars           int
}

func NewSJISProbe() *SJISProbe {
	return &SJISProbe{
		codingSM: NewCodingStateMachine(*SjisSmModel()),
	}
}
