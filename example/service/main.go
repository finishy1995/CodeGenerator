package main

import (
	"CodeGenerator/dataloader"
	"CodeGenerator/extension/dataloader/proto"
	"CodeGenerator/extension/logic/datetime"
	"CodeGenerator/extension/logic/numberhelper"
	"CodeGenerator/extension/logic/stringhelper"
	"CodeGenerator/generator"
	"CodeGenerator/generator/define"
	"CodeGenerator/generator/logic"
	"CodeGenerator/library/log"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var protoPath = flag.String("proto", "./account.proto", "the proto file path")

func main() {
	flag.Parse()

	//log.SetLevel(log.DEBUG)
	dataloader.AddDataLoader(".proto", proto.NewLoader())
	data := dataloader.LoadFromFile(*protoPath)
	d := define.NewDictionary()
	d.SetData(data)
	d.AddKeyValuePair(".PathSuffix", "ProjectX/service/")
	d.AddKeyValuePair(".PathBase", "ProjectX/base")

	logic.RegisterAll()
	datetime.Register()
	stringhelper.Register()
	numberhelper.Register()
	m := generator.NewMission(d,
		"./tpl",
		fmt.Sprintf("./%s", d.FindOrReturnDefault(".package", "service")))
	m.Run()

	// Generate grpc file using protoc
	serviceName := d.FindOrReturnDefault(".package", "service")
	path := fmt.Sprintf("./%s/pb", serviceName)
	err := os.MkdirAll(path, 0660)
	if err != nil {
		log.Error("cannot create pb dir, error: %s", err.Error())
		return
	}
	cmd := exec.Command("protoc",
		fmt.Sprintf("--go_out=%s", path),
		fmt.Sprintf("--go-grpc_out=%s", path),
		*protoPath,
	)
	if err := cmd.Run(); err != nil {
		log.Error("protoc failed, error: %s", err.Error())
	}
}
