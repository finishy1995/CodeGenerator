package main

import (
	"github.com/finishy1995/codegenerator/generator"
	"github.com/finishy1995/codegenerator/generator/define"
	"github.com/finishy1995/codegenerator/generator/logic"
)

func main() {
	d := define.NewDictionary()
	logic.RegisterAll()
	m := generator.NewMission(d, "./tpl", "./output")
	m.Run()
}
