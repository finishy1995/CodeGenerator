package define

import (
	"github.com/finishy1995/codegenerator/library/log"
	"github.com/finishy1995/codegenerator/library/utils"
	"path"
	"strings"
)

// NewTaskFunc tplDir + "/" + tplName can find this tpl file
type NewTaskFunc func(tplDir string, tplName string, outputDir string) Task

type UseDict interface {
	// GetGlobalDict get global dictionary address
	GetGlobalDict() Dictionary
	// GetTempDict get template dictionary address
	GetTempDict() Dictionary
	// FindInDict find the value by key in dictionary
	FindInDict(key string) string
}

// Task every template file will create one task
//
//	input: task -> line(s) -> step(s)
//	output: task -> file(s)
type Task interface {
	UseDict
	SetGlobalDict(dict Dictionary)

	// GetTplDir get template files root dir
	GetTplDir() string
	// GetOutputDir get output dir path
	GetOutputDir() string
	// GetTplNameWithoutSuffix get tpl file name without .tpl suffix (related to tpl dir)
	GetTplNameWithoutSuffix() string
	// GetLineContent get one line content
	GetLineContent(lineIndex int) (string, bool)
	// GetAllLineContent get all line content address
	GetAllLineContent() []string

	// CreateFile create file to output
	CreateFile() File
	// EndFile end file, output right now
	EndFile() File
	// GetFile get file
	GetFile() File

	// Exec task
	Exec() string
}

var (
	newTaskFuncInstance = NewDefaultTask
)

func SetNewTaskFunc(f NewTaskFunc) {
	newTaskFuncInstance = f
}

func NewTask(tplDir string, tplName string, outputDir string) Task {
	if tplName == "" || outputDir == "" || newTaskFuncInstance == nil {
		return nil
	}
	return newTaskFuncInstance(tplDir, tplName, outputDir)
}

// ------------ DefaultTask ------------ //

type DefaultTask struct {
	globalDict  Dictionary
	tempDict    Dictionary
	lineList    []string
	tplName     string
	tplDir      string
	outputDir   string
	currentFile File
}

func NewDefaultTask(tplDir string, tplName string, outputDir string) Task {
	name := tplName
	if strings.HasSuffix(tplName, TplFileSuffix) {
		name = name[:len(name)-len(TplFileSuffix)]
	}
	content, err := utils.GetFileContent(path.Join(tplDir, tplName))
	if err != nil {
		log.Error("read template file %s failed, error: %s", tplName, err.Error())
		return nil
	}
	arr := strings.Split(content, TplLineFeed)

	task := &DefaultTask{
		tplName:    name,
		tplDir:     tplDir,
		outputDir:  outputDir,
		lineList:   append([]string{""}, arr...),
		globalDict: NewDictionary(),
		tempDict:   NewDictionary(),
	}
	return task
}

func (d *DefaultTask) GetTplDir() string {
	return d.tplDir
}

func (d *DefaultTask) SetGlobalDict(dict Dictionary) {
	d.globalDict = dict
}

func (d *DefaultTask) GetGlobalDict() Dictionary {
	return d.globalDict
}

func (d *DefaultTask) GetTempDict() Dictionary {
	return d.tempDict
}

func (d *DefaultTask) FindInDict(key string) string {
	if value, ok := d.tempDict.Find(key); ok {
		return value
	}
	return d.globalDict.FindOrReturnDefault(key, "")
}

func (d *DefaultTask) GetOutputDir() string {
	return d.outputDir
}

func (d *DefaultTask) GetTplNameWithoutSuffix() string {
	return d.tplName
}

func (d *DefaultTask) GetLineContent(lineIndex int) (string, bool) {
	if lineIndex >= len(d.lineList) {
		return "", false
	}
	return d.lineList[lineIndex], true
}

func (d *DefaultTask) GetAllLineContent() []string {
	return d.lineList
}

func (d *DefaultTask) CreateFile() File {
	d.currentFile = NewFile(d)
	return d.currentFile
}

func (d *DefaultTask) Exec() string {
	line := NewLine(d)
	overLine := uint32(len(d.lineList))
	result := ""

	for {
		if overLine <= line.NextIndex {
			break
		}

		content := d.process(line)
		log.Debug("line: %d, obj: %+v \n\t content: %s, len(content): %d", line.GetIndex(), line, content, len(content))
		if d.currentFile != nil {
			d.currentFile.Append(content)
		}
		result += content
	}

	result = deleteEOFFeed(result)

	if d.currentFile != nil {
		d.currentFile.Generate()
	}
	return result
}

func (d *DefaultTask) process(line *Line) string {
	line.Update()
	err := line.Exec()
	if err != nil {
		log.Error("cannot process line %d, content: %s", line.GetIndex(), line.GetOrigin())
	}

	return line.GetOutput()
}

func (d *DefaultTask) EndFile() File {
	if d.currentFile == nil {
		return nil
	}
	file := d.currentFile
	d.currentFile = nil
	file.Generate()
	return file
}

func (d *DefaultTask) GetFile() File {
	return d.currentFile
}

func deleteEOFFeed(content string) string {
	// 去除行末的换行符
	if strings.HasSuffix(content, OutLineFeed) {
		return content[:len(content)-len(OutLineFeed)]
	}
	return content
}
