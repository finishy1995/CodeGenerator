package numberhelper

import "CodeGenerator/generator/define"

func Register() {
	define.Register(NewLogicKeyCalc())
}
