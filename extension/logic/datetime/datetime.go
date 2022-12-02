package datetime

import (
	"CodeGenerator/generator/define"
	"time"
)

// DateTime print datetime as format
type DateTime struct {
}

func (dt *DateTime) GetKey() string {
	return "DateTime"
}

func (dt *DateTime) Exec(step define.Step, args ...string) {
	now := time.Now()
	step.SetOutput(now.Format("2006-01-02 15:04:05"))
}

func NewLogicKeyDateTime() define.LogicKey {
	return &DateTime{}
}

func Register() {
	define.Register(NewLogicKeyDateTime())
}
