package stringhelper

import (
	"github.com/finishy1995/codegenerator/generator/define"
	"github.com/finishy1995/codegenerator/generator/logic"
	"strings"
)

// Upper get lower output
//
// Example
//
//	Upper hello => HELLO
//	Upper hello 1 => Hello
type Upper struct {
}

func (u *Upper) GetKey() string {
	return "Upper"
}

func (u *Upper) Exec(step define.Step, args ...string) {
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
		step.SetOutput(strings.ToUpper(args[0]))
	} else {
		step.SetOutput(strings.ToUpper(args[0][:count]) + args[0][count:])
	}
}

func NewLogicKeyUpper() define.LogicKey {
	return &Upper{}
}
