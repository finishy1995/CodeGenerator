package define

import (
	"CodeGenerator/library/log"
	"CodeGenerator/library/utils"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NewDictionaryFunc func() Dictionary

type Dictionary interface {
	fmt.Stringer

	SetData(data interface{})
	Reset()
	ReGenerate()

	Find(key string) (value string, ok bool)
	FindOrReturnDefault(key string, defaultValue string) string
	AddKeyValueMap(keyValueMap map[string]string)
	AddKeyValuePair(key string, value string)
	GetAllKeyValueMap() map[string]string
}

var (
	newDictionaryFuncInstance = NewDefaultDictionary
)

func SetNewDictionaryFunc(f NewDictionaryFunc) {
	newDictionaryFuncInstance = f
}

func NewDictionary() Dictionary {
	if newDictionaryFuncInstance == nil {
		return nil
	}
	return newDictionaryFuncInstance()
}

// ------------ DefaultDictionary ------------ //

type DefaultDictionary struct {
	data   interface{}
	dict   map[string]string
	prefix string
}

func NewDefaultDictionary() Dictionary {
	dict := &DefaultDictionary{}
	dict.Reset()
	return dict
}

func (d *DefaultDictionary) Reset() {
	d.dict = make(map[string]string)
}

func (d *DefaultDictionary) ReGenerate() {
	if d.data == nil {
		d.Reset()
	} else {
		d.dict = d.generate(d.prefix, d.data)
	}
}

func (d *DefaultDictionary) AddKeyValuePair(key string, value string) {
	// skip special param add
	if strings.HasSuffix(key, DictLengthMark) || strings.HasSuffix(key, DictKeyMark) {
		return
	}

	duplicateKey := false
	if _, ok := d.dict[key]; ok {
		duplicateKey = true
	}
	d.dict[key] = value
	if duplicateKey {
		return
	}

	// Regenerate *Length and *Key
	checkStr := key
	for {
		index := strings.LastIndex(checkStr, ".")
		if index < 0 {
			return
		}

		lastKey := checkStr[index+1:]
		checkStr = checkStr[:index]
		checkStrLengthKey := fmt.Sprintf("%s.%s", checkStr, DictLengthMark)
		checkStrKeyKey := fmt.Sprintf("%s.%s", checkStr, DictKeyMark)
		if oldValue, ok := d.dict[checkStrLengthKey]; ok {
			length, err := strconv.Atoi(oldValue)
			if err != nil {
				log.Error("unknown error when try add a key value pair to dict, error: %s", err.Error())
				length = 0
			}
			d.dict[checkStrLengthKey] = strconv.Itoa(length + 1)
			d.dict[checkStrKeyKey] += "," + lastKey
			return
		} else {
			d.dict[checkStrLengthKey] = "1"
			d.dict[checkStrKeyKey] = lastKey
		}
	}
}

func (d *DefaultDictionary) AddKeyValueMap(keyValueMap map[string]string) {
	for key, value := range keyValueMap {
		d.AddKeyValuePair(key, value)
	}
}

func (d *DefaultDictionary) Find(key string) (value string, ok bool) {
	value, ok = d.dict[key]
	return
}

func (d *DefaultDictionary) FindOrReturnDefault(key string, defaultValue string) string {
	if value, ok := d.dict[key]; ok {
		return value
	} else {
		return ""
	}
}

func (d *DefaultDictionary) SetPrefix(prefix string) {
	d.prefix = prefix
}

func (d *DefaultDictionary) SetData(data interface{}) {
	d.data = data
	d.ReGenerate()
}

func (d *DefaultDictionary) generate(prefix string, input interface{}) map[string]string {
	k := reflect.TypeOf(input).Kind()
	switch k {
	case reflect.String:
		return map[string]string{prefix: input.(string)}
	case reflect.Slice:
		sli := input.([]interface{})
		result := make(map[string]string)
		length := 0
		keys := ""
		for index, item := range sli {
			if d.CheckType(item) {
				// index will start from 1, for human style reading
				realIndex := index + 1
				utils.MergeMap(result, d.generate(fmt.Sprintf("%s.%d", prefix, realIndex), item))
				length++
				keys += strconv.Itoa(realIndex) + ","
			}
		}
		result[fmt.Sprintf("%s.%s", prefix, DictLengthMark)] = strconv.Itoa(length)
		result[fmt.Sprintf("%s.%s", prefix, DictKeyMark)] = keys[:len(keys)-1]
		return result
	case reflect.Map:
		m := input.(map[string]interface{})
		result := make(map[string]string)
		length := 0
		keys := ""
		for key, item := range m {
			if d.CheckType(item) {
				utils.MergeMap(result, d.generate(fmt.Sprintf("%s.%s", prefix, key), item))
				length++
				keys += key + ","
			}
		}
		result[fmt.Sprintf("%s.%s", prefix, DictLengthMark)] = strconv.Itoa(length)
		result[fmt.Sprintf("%s.%s", prefix, DictKeyMark)] = keys[:len(keys)-1]
		return result
	default:
		return nil
	}
}

// CheckType check if value type is supported, will not recurse
func (d *DefaultDictionary) CheckType(value interface{}) bool {
	if value == nil {
		return false
	}
	k := reflect.TypeOf(value).Kind()
	if k == reflect.Slice || k == reflect.Map || k == reflect.String {
		return true
	}
	return false
}

func (d *DefaultDictionary) String() string {
	result := ""
	for key, value := range d.dict {
		result += fmt.Sprintf("%-50s : %s\n", key, value)
	}
	return result
}

func (d *DefaultDictionary) GetAllKeyValueMap() map[string]string {
	var target map[string]string
	err := utils.DeepCopy(d.dict, &target)
	if err != nil {
		log.Error("cannot get a new map now, error: %s", err.Error())
		return nil
	}
	return target
}
