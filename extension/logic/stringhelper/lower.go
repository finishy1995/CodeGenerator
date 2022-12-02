package stringhelper

import (
	"CodeGenerator/generator/define"
	"CodeGenerator/generator/logic"
	"strings"
)

// Lower get lower output
//
// Example
//
//	Lower HELLO => hello
//	Lower HELLO 1 => hELLO
type Lower struct {
}

func (l *Lower) GetKey() string {
	return "Lower"
}

func (l *Lower) Exec(step define.Step, args ...string) {
	if len(args) < 1 {
		return
	}
	count := 0
	if len(args) > 1 {
		num, err := logic.String2Int(args[1])
		if err == nil {
			count = num
		}
		length := len(args[0])
		if count > length {
			count = 0
		}
	}
	if count < 1 {
		step.SetOutput(strings.ToLower(args[0]))
	} else {
		step.SetOutput(strings.ToLower(args[0][:count]) + args[0][count:])
	}
}

func NewLogicKeyLower() define.LogicKey {
	return &Lower{}
}
