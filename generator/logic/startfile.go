package logic

import "CodeGenerator/generator/define"

// StartFile start file output
type StartFile struct {
}

func (sf *StartFile) GetKey() string {
	return "StartFile"
}

func (sf *StartFile) Exec(step define.Step, args ...string) {
	step.GetLine().GetTask().CreateFile()
	if len(args) > 0 {
		step.GetTempDict().AddKeyValuePair(define.DefaultFileName, args[0])
	}
	if len(args) > 1 {
		step.GetTempDict().AddKeyValuePair(define.DefaultFileSuffix, args[1])
	}
}

func NewLogicKeyStartFile() define.LogicKey {
	return &StartFile{}
}
