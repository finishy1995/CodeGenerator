package logic

import (
	"fmt"
	"github.com/finishy1995/codegenerator/generator/define"
	"github.com/finishy1995/codegenerator/library/log"
	"strings"
)

// GetKey get key using index
type GetKey struct {
}

func (gk *GetKey) GetKey() string {
	return "GetKey"
}

func (gk *GetKey) Exec(step define.Step, args ...string) {
	if len(args) < 2 {
		return
	}
	num, err := String2Int(args[1])
	if err != nil {
		log.Error("cannot handle GetKey cmd, line %d, key %s, index %s", step.GetLine().GetIndex(), args[0], args[1])
		return
	}

	keyList := step.FindInDict(fmt.Sprintf("%s.%s", args[0], define.DictKeyMark))
	if keyList != "" {
		arr := strings.Split(keyList, ",")
		if num > len(arr)+1 {
			return
		}
		step.SetOutput(step.FindInDict(fmt.Sprintf("%s.%s", args[0], arr[num-1])))
	}
}

func NewLogicKeyGetKey() define.LogicKey {
	return &GetKey{}
}
