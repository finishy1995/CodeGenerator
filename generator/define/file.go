package define

import (
	"CodeGenerator/library/log"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File interface {
	Append(content string)
	Generate()
}

type NewFileFunc func(Task) File

var (
	newFileFuncInstance = NewDefaultFile
)

func SetNewFileFunc(f NewFileFunc) {
	newFileFuncInstance = f
}

func NewFile(task Task) File {
	if task == nil || newFileFuncInstance == nil {
		return nil
	}
	return newFileFuncInstance(task)
}

// ------------ DefaultFile ------------ //

type DefaultFile struct {
	task    Task
	content string
}

// file description params
var (
	DefaultFileName    = "file.name"
	DefaultFileSuffix  = "file.suffix"
	DefaultFileSkipGen = "file.skip_generate"
)

func NewDefaultFile(task Task) File {
	dict := task.GetTempDict()
	dict.AddKeyValueMap(map[string]string{
		DefaultFileName:    task.GetTplNameWithoutSuffix(),
		DefaultFileSuffix:  "",
		DefaultFileSkipGen: "false",
	})

	return &DefaultFile{
		task: task,
	}
}

func (d *DefaultFile) Append(content string) {
	d.content += content
}

func (d *DefaultFile) Generate() {
	skip, err := strconv.ParseBool(d.task.FindInDict(DefaultFileSkipGen))
	if skip {
		return
	}

	filename := fmt.Sprintf("%s%s", d.task.FindInDict(DefaultFileName), d.task.FindInDict(DefaultFileSuffix))
	// Generate all upper dir
	index := strings.LastIndex(filename, "/")
	if index > 0 {
		err = os.MkdirAll(filename[:index], 0660)
	}
	if err != nil {
		log.Error("cannot mkdir all upper dir, error: %s", err.Error())
		return
	}

	d.content = deleteEOFFeed(d.content)

	err = os.WriteFile(
		fmt.Sprintf("%s/%s", d.task.GetOutputDir(), filename),
		[]byte(d.content),
		0666,
	)
	if err != nil {
		log.Error("cannot generate %s, error: %s", filename, err.Error())
	}
}
