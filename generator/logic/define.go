package logic

import (
	"CodeGenerator/generator/define"
)

// Define key-value pair in temp dict, will overwrite key-value pair if exist already
type Define struct {
}

func (d *Define) GetKey() string {
	return "Define"
}

func (d *Define) Exec(step define.Step, args ...string) {
	arg := MergeArgs(args...)
	m := GetArgsMap(arg)
	tempDict := step.GetTempDict()
	for key, value := range m {
		tempDict.AddKeyValuePair(key, value)
	}
}

func NewLogicKeyDefine() define.LogicKey {
	return &Define{}
}
