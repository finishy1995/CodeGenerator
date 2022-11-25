package logic

import (
	"strconv"
	"strings"
)

type KeyValuePair struct {
	Key   string
	Value string
}

// GetArgsKeyValuePair format "a=1"
func GetArgsKeyValuePair(args string) *KeyValuePair {
	index := strings.Index(args, "=")
	if index < 0 {
		return nil
	}
	return &KeyValuePair{
		Key:   args[:index],
		Value: args[index+1:],
	}
}

// GetArgsMap format "a=1,b=s,c=1..."
func GetArgsMap(args string) map[string]string {
	arr := strings.Split(args, ",")
	result := make(map[string]string, len(arr))
	for _, str := range arr {
		pair := GetArgsKeyValuePair(str)
		if pair == nil {
			continue
		}
		result[pair.Key] = pair.Value
	}

	return result
}

// MergeArgs merge several args to a string ["a", " ", "=", " ", "1"] => "a=1"
func MergeArgs(args ...string) string {
	result := ""
	for _, str := range args {
		result += str
	}
	return result
}

func Int2String(num int) string {
	return strconv.Itoa(num)
}

func String2Int(str string) (int, error) {
	return strconv.Atoi(str)
}

func String2Bool(str string) (bool, error) {
	return strconv.ParseBool(str)
}
