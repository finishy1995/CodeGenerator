package stringhelper

import "github.com/finishy1995/codegenerator/generator/define"

func Register() {
	define.Register(NewLogicKeyLower(), NewLogicKeyUpper())
}
