package numberhelper

import (
	"CodeGenerator/generator/define"
	"CodeGenerator/generator/logic"
	"CodeGenerator/library/log"
	"fmt"
	"github.com/dengsgo/math-engine/engine"
)

// Calc calculate a math expression
type Calc struct {
}

func (c *Calc) GetKey() string {
	return "Calc"
}

func (c *Calc) Exec(step define.Step, args ...string) {
	expr := logic.MergeArgs(args...)
	result, err := engine.ParseAndExec(expr)
	if err != nil {
		log.Error("calc execution failed, error: %s", err.Error())
	}
	step.SetOutput(fmt.Sprintf("%v", result))
}

func NewLogicKeyCalc() define.LogicKey {
	return &Calc{}
}
