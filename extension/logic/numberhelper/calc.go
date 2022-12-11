package numberhelper

import (
	"fmt"
	"github.com/dengsgo/math-engine/engine"
	"github.com/finishy1995/codegenerator/generator/define"
	"github.com/finishy1995/codegenerator/generator/logic"
	"github.com/finishy1995/codegenerator/library/log"
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
