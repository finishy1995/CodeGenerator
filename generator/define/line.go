package define

import (
	"errors"
)

// Line every line will create several steps, cannot think any situation to rewrite the logic
//
//	task -> line(s) -> step(s)
type Line struct {
	origin    string
	index     uint32
	NextIndex uint32
	Feed      bool
	Skip      bool
	rootStep  Step
	task      Task
}

func NewLine(task Task) *Line {
	if task == nil {
		return nil
	}

	return &Line{
		origin:    "",
		index:     0,
		NextIndex: 1, // start from line 1
		Feed:      true,
		Skip:      false,
		rootStep:  nil,
		task:      task,
	}
}

func (l *Line) Update() {
	l.index = l.NextIndex
	l.origin, _ = l.task.GetLineContent(int(l.index))
	l.NextIndex++
	l.Feed = true
}

func (l *Line) GetOrigin() string {
	return l.origin
}

func (l *Line) GetIndex() uint32 {
	return l.index
}

func (l *Line) GetRootStep() Step {
	return l.rootStep
}

func (l *Line) CreateRootStep() {
	l.rootStep = NewStep(l)
}

func (l *Line) GetTask() Task {
	return l.task
}

func (l *Line) Exec() error {
	l.CreateRootStep()
	if l.rootStep == nil {
		return errors.New("cannot create step")
	}
	l.rootStep.Exec()
	return nil
}

func (l *Line) GetOutput() string {
	if l.rootStep == nil {
		return ""
	}
	if l.Skip {
		return ""
	}
	content := l.rootStep.GetOutput()
	// if you have output, must add feed
	if !l.Feed && content != "" {
		l.Feed = true
	}
	if l.Feed {
		content += OutLineFeed
	}
	return content
}
