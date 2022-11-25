package logic

import "CodeGenerator/generator/define"

// EndIf end if block
type EndIf struct {
	i *If
}

func (ei *EndIf) GetKey() string {
	return "EndIf"
}

func (ei *EndIf) Exec(step define.Step, args ...string) {
	ei.i.Stack.Pop()
	step.GetLine().Skip = false
}

func NewLogicKeyEndIf(i *If) define.LogicKey {
	return &EndIf{
		i: i,
	}
}
