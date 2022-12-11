package logic

import (
	"github.com/finishy1995/codegenerator/generator/define"
	"github.com/finishy1995/codegenerator/library/stack"
)

// If
//
// Example:
//
//	.a undefined || ""
//	.b = false
//	.c = 1
//	.d = true
//	#{If} == #{If false} == #{If .a} == #{If #{.b}} => false
//	#{If true} == #{If .b} == #{If .c} == #{If .d} == #{If #{.d}} => true
type If struct {
	Stack *stack.Stack
}

type ifInfo struct {
	result bool
	inElse bool
}

func (i *If) GetKey() string {
	return "If"
}

func (i *If) Exec(step define.Step, args ...string) {
	// 无参数默认条件失败
	if len(args) == 0 {
		i.exec(step, false)
		return
	}

	// 判断参数是否是 true or false
	value, err := String2Bool(args[0])
	if err != nil {
		// 判断是否是字典里的 key
		str := step.FindInDict(args[0])
		if str == "" {
			i.exec(step, false)
		} else {
			i.exec(step, true)
		}
	} else {
		i.exec(step, value)
	}
}

func (i *If) exec(step define.Step, b bool) {
	i.Stack.Push(&ifInfo{
		result: b,
		inElse: false,
	})

	if !b {
		step.GetLine().Skip = true
	}
}

func NewLogicKeyIf() *If {
	return &If{
		Stack: stack.NewStack(),
	}
}
