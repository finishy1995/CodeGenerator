package main

import (
	"CodeGenerator/generator"
	"CodeGenerator/generator/define"
	"CodeGenerator/generator/logic"
)

func main() {
	d := define.NewDictionary()
	logic.RegisterAll()
	m := generator.NewMission(d, "./tpl", "./output")
	m.Run()
}
