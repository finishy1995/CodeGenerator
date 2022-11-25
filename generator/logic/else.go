package logic

import "CodeGenerator/generator/define"

// Else using in if block, jump to another situation
type Else struct {
	l *If
}

func (e *Else) GetKey() string {
	return "Else"
}

func (e *Else) Exec(step define.Step, args ...string) {
	inter := e.l.Stack.Pop()
	if inter == nil {
		return
	}
	info := inter.(*ifInfo)
	info.inElse = true
	if info.result {
		step.GetLine().Skip = true
	} else {
		step.GetLine().Skip = false
	}
	e.l.Stack.Push(info)
}

func NewLogicKeyElse(l *If) define.LogicKey {
	return &Else{
		l: l,
	}
}
