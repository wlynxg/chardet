package probe

import (
	"math"

	"github.com/wlynxg/chardet/consts"
)

type UTF8Probe struct {
	CharSetProbe
	OneCharProb       float64
	ShortcutThreshold float64

	codingSM   *CodingStateMachine
	numMbChars int
}

func NewUTF8Probe() *UTF8Probe {
	return &UTF8Probe{
		OneCharProb:       0.5,
		ShortcutThreshold: 0.95,
		codingSM:          NewCodingStateMachine(UTF8SmModel()),
		numMbChars:        0,
	}
}

func (u *UTF8Probe) CharSetName() string {
	return consts.UTF8
}

func (u *UTF8Probe) Language() string {
	return ""
}

func (u *UTF8Probe) Reset() {
	u.CharSetProbe.Reset()
	u.numMbChars = 0
	if u.codingSM != nil {
		u.codingSM.Reset()
	}
}

func (u *UTF8Probe) Feed(data []byte) consts.ProbingState {
loop:
	for _, datum := range data {
		codingState := u.codingSM.NextState(datum)
		switch codingState {
		case consts.ErrorMachineState:
			u.state = consts.NotMeProbingState
			break loop
		case consts.ItsMeMachineState:
			u.state = consts.FoundItProbingState
			break loop
		case consts.StartMachineState:
			if u.codingSM.CurrentCharLength() >= 2 {
				u.numMbChars++
			}
		default:
		}
	}

	if u.state == consts.DetectingProbingState && u.GetConfidence() > u.ShortcutThreshold {
		u.state = consts.FoundItProbingState
	}
	return u.state
}

func (u *UTF8Probe) GetConfidence() float64 {
	unlike := 0.99
	if u.numMbChars < 6 {
		unlike *= math.Pow(u.OneCharProb, float64(u.numMbChars))
		return 1.0 - unlike
	} else {
		return unlike
	}
}
