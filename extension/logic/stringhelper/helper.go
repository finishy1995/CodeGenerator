package stringhelper

import "CodeGenerator/generator/define"

func Register() {
	define.Register(NewLogicKeyLower(), NewLogicKeyUpper())
}
