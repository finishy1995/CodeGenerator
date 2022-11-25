package logic

import (
	"CodeGenerator/generator/define"
	"CodeGenerator/library/log"
	"CodeGenerator/library/utils"
	"strings"
)

const (
	IndexMarkName   = "index"
	IndexValueStart = 1
)

type loopInfo struct {
	loopIndex     int
	lineIndex     uint32
	loopLength    int
	loopIndexKey  string
	startCallback []LoopStartCallback
	endCallback   []LoopEndCallback
}

type LoopStartCallback func(step define.Step)
type LoopEndCallback func()

// Loop exec the following lines
//
//	Example: Loop 3 index=param.i
type Loop struct {
	stack []*loopInfo
}

func (l *Loop) GetKey() string {
	return "Loop"
}

func (l *Loop) InLoop() bool {
	if len(l.stack) > 0 {
		return true
	}
	return false
}

func (l *Loop) SetStartCallback(f LoopStartCallback) {
	length := len(l.stack)
	if length < 1 {
		return
	}
	l.stack[length-1].startCallback = append(l.stack[length-1].startCallback, f)
}

func (l *Loop) SetEndCallback(f LoopEndCallback) {
	length := len(l.stack)
	if length < 1 {
		return
	}
	l.stack[length-1].endCallback = append(l.stack[length-1].endCallback, f)
}

func (l *Loop) Exec(step define.Step, args ...string) {
	if len(args) < 1 {
		log.Error("loop need at least one args, step: %+v", step)
		return
	}

	arr := strings.Split(args[0], ",")
	if len(arr) < 1 {
		log.Error("loop args wrong, args: %+v, step: %+v", args, step)
		return
	}
	if arr[0] == "" {
		arr[0] = "0" // if we cannot read length, set default to 0
	}

	count, err := String2Int(arr[0])
	if err != nil {
		log.Error("loop args wrong, args: %+v, step: %+v, error: %s", args, step, err.Error())
		return
	}
	argsMap := make(map[string]string)
	if len(arr) > 1 {
		str := MergeArgs(arr[1:]...)
		utils.MergeMap(argsMap, GetArgsMap(str))
	}
	if len(args) > 1 {
		str := MergeArgs(args[1:]...)
		utils.MergeMap(argsMap, GetArgsMap(str))
	}
	indexName := ""
	if newIndexName, ok := argsMap[IndexMarkName]; ok {
		indexName = newIndexName
	}

	l.updateToStack(count, step, indexName)
}

func (l *Loop) updateToStack(count int, step define.Step, index string) {
	if count == 0 {
		step.GetLine().Skip = true
	} else if index != "" {
		step.GetTempDict().AddKeyValuePair(index, Int2String(IndexValueStart))
	}

	l.stack = append(l.stack, &loopInfo{
		loopIndex:     IndexValueStart,
		lineIndex:     step.GetLine().GetIndex(),
		loopLength:    count,
		loopIndexKey:  index,
		startCallback: make([]LoopStartCallback, 0, 0),
		endCallback:   make([]LoopEndCallback, 0, 0),
	})
}

func NewLogicKeyLoop() *Loop {
	return &Loop{
		stack: make([]*loopInfo, 0, 3),
	}
}
