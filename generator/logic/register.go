package logic

import "github.com/finishy1995/codegenerator/generator/define"

var (
	loopInstance *Loop
	ifInstance   *If
)

func RegisterAll() {
	RegisterBase()
	RegisterLoop()
	RegisterIf()
	RegisterFile()
}

func RegisterBase() {
	define.Register(NewLogicKeyDefine(), NewLogicKeyGetKey())
}

func RegisterLoop() {
	loopInstance = NewLogicKeyLoop()
	define.Register(loopInstance, NewLogicKeyEndLoop(loopInstance))
}

func RegisterIf() {
	ifInstance = NewLogicKeyIf()
	define.Register(ifInstance, NewLogicKeyEndIf(ifInstance), NewLogicKeyElse(ifInstance))
}

func RegisterFile() {
	define.Register(NewLogicKeyInsert(), NewLogicKeyEndFile(), NewLogicKeyStartFile())
}

func GetLoop() *Loop {
	return loopInstance
}

func GetIf() *If {
	return ifInstance
}
