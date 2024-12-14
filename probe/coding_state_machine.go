package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type CodingStateMachine struct {
	model       *StateMachineModel
	currState   consts.MachineState
	currBytePos int
	currCharLen int

	Active bool
}

func NewCodingStateMachine(sm StateMachineModel) *CodingStateMachine {
	return &CodingStateMachine{
		currState: consts.StartMachineState,
		model:     &sm,
	}
}

func (o *CodingStateMachine) Reset() {
	o.currState = consts.StartMachineState
}

func (o *CodingStateMachine) CurrentCharLength() int {
	return o.currCharLen
}

func (o *CodingStateMachine) CodingStateMachine() string {
	if o.model == nil {
		return ""
	}
	return o.model.Name
}

func (o *CodingStateMachine) Language() string {
	if o.model == nil {
		return ""
	}
	return o.model.Language
}

func (o *CodingStateMachine) NextState(b byte) consts.MachineState {
	// for each byte we get its class
	// if it is first byte, we also get byte length
	byteClass := o.model.ClassTable[b]
	if o.currState == consts.StartMachineState {
		o.currBytePos = 0
		o.currCharLen = int(o.model.CharLenTable[byteClass])
	}

	// from byte's class and state_table, we get its next state
	currState := byte(o.currState)*o.model.ClassFactor + byteClass
	o.currState = o.model.StateTable[currState]
	o.currBytePos++
	return o.currState
}
