package logic

import (
	"CodeGenerator/generator/define"
	"CodeGenerator/library/log"
	"strings"
)

// Insert another file into the following line
type Insert struct {
}

func (d *Insert) GetKey() string {
	return "Insert"
}

// circular insert check
var (
	checkMap   = map[string]bool{}
	insertPath = make([]string, 0, 0)
)

func (d *Insert) Exec(step define.Step, args ...string) {
	if len(args) < 1 {
		return
	}

	for _, str := range args {
		task := step.GetLine().GetTask()
		newTask := define.NewTask(
			task.GetTplDir(),
			str,
			task.GetOutputDir(),
		)
		if newTask != nil {
			// check if circular insert
			taskName := task.GetTplNameWithoutSuffix() + define.TplFileSuffix
			if _, ok := checkMap[str]; ok {
				insertPath = append(insertPath, str)
				log.Error("circular insert! insert path: %s", strings.Join(insertPath, " -> "))
				return
			} else {
				if len(insertPath) == 0 {
					checkMap[taskName] = true
					checkMap[str] = true
					insertPath = append(insertPath, taskName, str)
				} else {
					checkMap[str] = true
					insertPath = append(insertPath, str)
				}
			}

			// add father task global dict and temp dict to son task global dict
			globalDictMap := task.GetGlobalDict().GetAllKeyValueMap()
			if globalDictMap != nil {
				newTask.GetGlobalDict().AddKeyValueMap(globalDictMap)
			}
			tempDictMap := task.GetTempDict().GetAllKeyValueMap()
			if tempDictMap != nil {
				newTask.GetGlobalDict().AddKeyValueMap(tempDictMap)
			}

			result := newTask.Exec()
			if result != "" {
				step.SetOutput(result)
			}

			// add son temp dict to father temp dict
			tempDictMap = newTask.GetTempDict().GetAllKeyValueMap()
			task.GetTempDict().AddKeyValueMap(tempDictMap)

			delete(checkMap, str)
			if len(insertPath) == 2 {
				delete(checkMap, taskName)
				insertPath = make([]string, 0, 0)
			} else {
				insertPath = insertPath[:len(insertPath)-1]
			}
			return
		}
	}
}

func NewLogicKeyInsert() define.LogicKey {
	return &Insert{}
}
