package chardet

type CodingStateMachine struct {
	currState MachineState
	model     *StateMachineModel
}

func NewCodingStateMachine(sm *StateMachineModel) *CodingStateMachine {
	return &CodingStateMachine{
		model: sm,
	}
}

func (o *CodingStateMachine) Reset() {
	o.currState = StartMachineState
}
