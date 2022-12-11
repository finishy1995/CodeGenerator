package logic

import (
	"github.com/finishy1995/codegenerator/generator/define"
)

// EndFile end file output
type EndFile struct {
}

func (ef *EndFile) GetKey() string {
	return "EndFile"
}

func (ef *EndFile) Exec(step define.Step, args ...string) {
	step.GetLine().GetTask().EndFile()
}

func NewLogicKeyEndFile() define.LogicKey {
	return &EndFile{}
}
