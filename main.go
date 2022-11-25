package main

import (
	"CodeGenerator/generator"
	"CodeGenerator/generator/define"
	"CodeGenerator/generator/logic"
	"flag"
)

var tplDir = flag.String("tpl", "./example/demo/tpl", "the tpl file directory")
var outputDir = flag.String("o", "./build", "the output file directory")

func main() {
	flag.Parse()

	//log.SetLevel(log.DEBUG)
	//dataloader.AddDataLoader(".proto", proto.NewLoader())
	//data := dataloader.LoadFromFile("./extension/dataloader/proto/account.proto")
	d := define.NewDictionary()
	//d.SetData(data)
	//define.SetLineFeedWindows()
	logic.RegisterAll()
	//fmt.Println(d.String())
	m := generator.NewMission(d, *tplDir, *outputDir)
	m.Run()
}
