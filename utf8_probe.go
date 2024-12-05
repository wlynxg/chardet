package chardet

import (
	"math"
)

type UTF8Probe struct {
	OneCharProb       float64
	ShortcutThreshold float64

	state      ProbingState
	codingSM   *CodingStateMachine
	numMbChars int
}

func NewUTF8Probe() *UTF8Probe {
	return &UTF8Probe{
		OneCharProb:       0.5,
		ShortcutThreshold: 0.95,
		state:             DetectingProbingState,
		codingSM:          NewCodingStateMachine(*UTF8SmModel()),
		numMbChars:        0,
	}
}

func (u *UTF8Probe) Confidence() float64 {
	unlike := 0.99
	if u.numMbChars < 6 {
		unlike *= math.Pow(u.OneCharProb, float64(u.numMbChars))
		return 1.0 - unlike
	} else {
		return unlike
	}
}

func (u *UTF8Probe) CharsetName() string {
	return UTF8ProbeCharsetName
}

func (u *UTF8Probe) Language() string {
	return ""
}

func (u *UTF8Probe) Reset() {
	u.state = DetectingProbingState
	u.numMbChars = 0
	if u.codingSM != nil {
		u.codingSM.Reset()
	}
}

func (u *UTF8Probe) Feed(data []byte) ProbingState {
	if u.codingSM == nil {
		return NotMeProbingState
	}

loop:
	for _, datum := range data {
		codingState := u.codingSM.NextState(datum)
		switch codingState {
		case ErrorMachineState:
			u.state = NotMeProbingState
			break loop
		case ItsMeMachineState:
			u.state = FoundItProbingState
			break loop
		case StartMachineState:
			if u.codingSM.CurrentCharLength() >= 2 {
				u.numMbChars++
			}
		default:
		}
	}

	if u.state == DetectingProbingState && u.Confidence() > u.ShortcutThreshold {
		u.state = FoundItProbingState
	}
	return u.state
}
