package define

import (
	"CodeGenerator/library/stack"
	"sort"
	"strings"
)

func init() {
	SetLogicMark(logicMarkPrefix, logicMarkSuffix)
}

// SetLogicMark set logic mark format
func SetLogicMark(prefix, suffix string) {
	logicMarkHandler = &markHandler{
		prefix:       prefix,
		prefixLength: len(prefix),
		suffix:       suffix,
		suffixLength: len(suffix),
	}
}

type logicIndex struct {
	Start  int
	LStart int
	End    int
	LEnd   int
}

func GetLogicIndexArray(content string) []*logicIndex {
	result := make([]*logicIndex, 0, 0)
	indexIsPrefix := make(map[int]bool)

	prefixArr := GetStringIndexAll(content, logicMarkPrefix)
	arr := prefixArr
	for _, value := range prefixArr {
		indexIsPrefix[value] = true
	}
	suffixArr := GetStringIndexAll(content, logicMarkSuffix)
	arr = append(arr, suffixArr...)
	for _, value := range suffixArr {
		indexIsPrefix[value] = false
	}

	if len(arr) < 2 {
		return result
	}
	sort.Ints(arr)
	sta := stack.NewStack()
	for _, index := range arr {
		if indexIsPrefix[index] {
			sta.Push(index)
		} else {
			length := sta.Len()
			if length == 0 {
				continue
			}
			value := sta.Pop()
			if length == 1 {
				realValue := value.(int)
				result = append(result, &logicIndex{
					Start:  realValue,
					LStart: realValue + logicMarkHandler.prefixLength,
					End:    index + 1,
					LEnd:   index - logicMarkHandler.suffixLength + 1,
				})
			}
		}
	}

	return result
}

func GetStringIndexAll(str, substr string) []int {
	result := make([]int, 0, 0)
	content := str
	start := 0
	for {
		index := strings.Index(content, substr)
		if index < 0 {
			break
		}
		content = content[index+1:]
		start += index
		result = append(result, start)
		start += 1
	}
	return result
}

func SetLineFeed(new string) {
	TplLineFeed = new
	OutLineFeed = new
}

func SetLineFeedWindows() {
	TplLineFeed = "\r\n"
	OutLineFeed = "\r\n"
}

func SetLineFeedLinux() {
	TplLineFeed = "\n"
	OutLineFeed = "\n"
}
