package main

import (
	"flag"
	"fmt"
	"github.com/finishy1995/codegenerator/dataloader"
	"github.com/finishy1995/codegenerator/extension/dataloader/proto"
	"github.com/finishy1995/codegenerator/extension/logic/datetime"
	"github.com/finishy1995/codegenerator/extension/logic/numberhelper"
	"github.com/finishy1995/codegenerator/extension/logic/stringhelper"
	"github.com/finishy1995/codegenerator/generator"
	"github.com/finishy1995/codegenerator/generator/define"
	"github.com/finishy1995/codegenerator/generator/logic"
	"github.com/finishy1995/codegenerator/library/log"
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
