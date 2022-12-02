package generator

import (
	"CodeGenerator/generator/define"
	"CodeGenerator/library/log"
	"CodeGenerator/library/utils"
	"os"
	"strings"
)

// Mission 一次代码生成为一个 mission，每个tpl的生成为一个 task
type Mission struct {
	dict      define.Dictionary
	dir       string
	taskList  []string
	pathList  []string
	active    int
	outputDir string
}

func NewMission(dict define.Dictionary, dir string, output string) *Mission {
	err := os.MkdirAll(output, 0660)
	if err != nil {
		log.Error("cannot create output dir, please check permissions, error: %s", err.Error())
		return nil
	}

	return &Mission{
		dict:      dict,
		dir:       dir,
		pathList:  utils.GetAllFilesInDir(dir),
		active:    -1,
		outputDir: output,
	}
}

func (m *Mission) Run() {
	m.taskList = make([]string, 0, 0)

	for _, path := range m.pathList {
		if strings.HasSuffix(path, define.TplFileSuffix) {
			relativePath := utils.GetRelativePath(path, m.dir)
			m.taskList = append(m.taskList, relativePath)
		}
	}

	// 逐一生成文件
	for _, path := range m.taskList {
		m.process(path)
	}
}

func (m *Mission) process(path string) {
	t := define.NewTask(m.dir, path, m.outputDir)
	if t == nil {
		log.Error("skip file %s generate, cannot create task", path)
	} else {
		t.SetGlobalDict(m.dict)
		t.CreateFile()
		t.Exec()
	}
}
