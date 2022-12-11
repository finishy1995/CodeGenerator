package define

import (
	"github.com/finishy1995/codegenerator/library/log"
	"strings"
)

type NewStepFunc func(line *Line) Step

// Step minimum unit for exec logic
type Step interface {
	UseDict

	// GetOutput get logic output
	GetOutput() string
	// SetOutput set logic output
	SetOutput(new string)
	// GetLogic get logic string
	GetLogic() string
	// SetLogic set logic string
	SetLogic(new string)
	// GetOrigin get origin tpl text
	GetOrigin() string
	// GetLine get line
	GetLine() *Line

	// GetFather get father step, root step has nil father(so can use this to check if root step)
	GetFather() Step
	// GetSonList get son list, remember to check if nil
	GetSonList() []Step
	// CreateStep create a step
	CreateStep(logic string) Step

	// Exec logic
	Exec()
}

var (
	newStepFuncInstance = NewDefaultStep
)

func SetNewStepFunc(f NewStepFunc) {
	newStepFuncInstance = f
}

func NewStep(line *Line) Step {
	if line == nil || newStepFuncInstance == nil {
		return nil
	}
	return newStepFuncInstance(line)
}

// ------------ DefaultStep ------------ //

// DefaultStep default step implement
type DefaultStep struct {
	origin string
	output string
	logic  string
	father Step
	son    []Step
	line   *Line
}

func NewDefaultStep(line *Line) Step {
	return &DefaultStep{
		origin: line.GetOrigin(),
		line:   line,
	}
}

func (d *DefaultStep) Exec() {
	arr := GetLogicIndexArray(d.origin)
	startIndex := 0
	for _, index := range arr {
		d.logic += d.origin[startIndex:index.Start]
		startIndex = index.End
		logic := d.origin[index.LStart:index.LEnd]

		log.Debug("step logic words: %s", logic)
		son := d.CreateStep(logic)
		son.Exec()
		d.logic += son.GetOutput()
	}

	// due to root step not in logic block，so we cannot exec root step, output directly
	if d.father == nil {
		d.output += d.logic + d.origin[startIndex:]
	} else {
		d.logic += d.origin[startIndex:]
		d.exec()
	}
}

func (d *DefaultStep) exec() {
	arr := strings.Split(d.logic, " ")
	arrWithoutEmpty := make([]string, 0, len(arr))

	for _, str := range arr {
		if str != "" {
			arrWithoutEmpty = append(arrWithoutEmpty, str)
		}
	}

	// 判断是否第一个是 logic key
	if len(arrWithoutEmpty) > 0 {
		if logicKey, ok := LogicKeyMap[arrWithoutEmpty[0]]; ok {
			// if run logic, will not output any content and feed, can change this in logic code
			d.output = ""
			logicKey.Exec(d, arrWithoutEmpty[1:]...)
			if d.output == "" {
				d.line.Feed = false
			}
			return
		}
	}

	for _, str := range arr {
		if str == "" {
			continue
		}

		d.output += d.FindInDict(str)
	}

}

func (d *DefaultStep) GetGlobalDict() Dictionary {
	return d.line.task.GetGlobalDict()
}

func (d *DefaultStep) GetTempDict() Dictionary {
	return d.line.task.GetTempDict()
}

func (d *DefaultStep) FindInDict(key string) string {
	return d.line.task.FindInDict(key)
}

func (d *DefaultStep) GetOutput() string {
	return d.output
}

func (d *DefaultStep) SetOutput(new string) {
	d.output = new
}

func (d *DefaultStep) GetLogic() string {
	return d.logic
}

func (d *DefaultStep) SetLogic(new string) {
	d.logic = new
}

func (d *DefaultStep) GetOrigin() string {
	return d.origin
}

func (d *DefaultStep) GetFather() Step {
	return d.father
}

func (d *DefaultStep) GetSonList() []Step {
	return d.son
}

func (d *DefaultStep) CreateStep(logic string) Step {
	step := NewDefaultStep(d.line).(*DefaultStep)
	step.origin = logic
	step.father = d
	if d.son == nil {
		d.son = []Step{step}
	} else {
		d.son = append(d.son, step)
	}
	return step
}

func (d *DefaultStep) GetLine() *Line {
	return d.line
}
