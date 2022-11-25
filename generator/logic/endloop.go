package logic

import (
	"CodeGenerator/generator/define"
	"CodeGenerator/library/log"
)

// EndLoop end loop block and generate again
type EndLoop struct {
	loop *Loop
}

func (el *EndLoop) GetKey() string {
	return "EndLoop"
}

func (el *EndLoop) Exec(step define.Step, args ...string) {
	stackLength := len(el.loop.stack)
	if stackLength < 1 {
		log.Error("unsupported define EndLoop, please define Loop first")
		return
	}
	info := el.loop.stack[stackLength-1]
	step.GetLine().Skip = false
	for _, f := range info.endCallback {
		f()
	}
	if info.loopIndex < info.loopLength {
		step.GetLine().NextIndex = info.lineIndex + 1 // goto Loop block first line
		info.loopIndex++
		if info.loopIndexKey != "" {
			step.GetTempDict().AddKeyValuePair(info.loopIndexKey, Int2String(info.loopIndex))
		}
		for _, f := range info.startCallback {
			f(step)
		}
	} else {
		el.loop.stack = el.loop.stack[:stackLength-1]
	}
}

func NewLogicKeyEndLoop(loop *Loop) define.LogicKey {
	if loop == nil {
		return nil
	}

	return &EndLoop{
		loop: loop,
	}
}
