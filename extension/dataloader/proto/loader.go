package proto

import (
	"CodeGenerator/dataloader/define"
	"github.com/jhump/protoreflect/desc/protoparse"
	"google.golang.org/protobuf/types/descriptorpb"
	"strconv"
	"strings"
)

type Loader struct {
}

func NewLoader() define.Loader {
	return &Loader{}
}

func (l Loader) LoadFromFile(filePath string) (map[string]interface{}, error) {
	parser := &protoparse.Parser{}
	des, err := parser.ParseFiles(filePath)
	if err != nil {
		return nil, err
	}
	if len(des) != 1 {
		return nil, ErrFileCannotLoad
	}
	fileDes := des[0].AsFileDescriptorProto()

	// env setting
	result := make(map[string]interface{})
	result["name"] = fileDes.GetName()
	result["shortname"] = getShortName(fileDes.GetName())
	result["syntax"] = fileDes.GetSyntax()
	if pack := fileDes.GetPackage(); pack != "" {
		result["package"] = pack
	}
	opts := fileDes.GetOptions()
	if opts != nil {
		optsPack := opts.GetGoPackage()
		if optsPack != "" {
			result["option"] = map[string]interface{}{"go_package": optsPack}
		}
	}

	// extension setting
	extend := fileDes.GetExtension()
	if extend != nil {
		extendMap := make(map[string]interface{})
		getFieldSlice(extendMap, extend)
		result["extend"] = extendMap
	}

	// message setting
	messages := fileDes.GetMessageType()
	if messages != nil {
		messageSlice := make([]interface{}, 0, len(messages))
		for _, message := range messages {
			messageMap := make(map[string]interface{})
			messageMap["name"] = message.GetName()
			getFieldSlice(messageMap, message.GetField())
			messageSlice = append(messageSlice, messageMap)
		}
		result["message"] = messageSlice
	}

	// enum setting
	enumDes := fileDes.GetEnumType()
	if enumDes != nil {
		enumSlice := make([]interface{}, 0, len(enumDes))
		enumHelperMap := make(map[string]interface{})
		for _, enumItem := range enumDes {
			enumMap := make(map[string]interface{})
			enumItemName := enumItem.GetName()
			enumMap["name"] = enumItemName
			enumHelperMapInner := make(map[string]interface{})
			enumHelperMap[enumItemName] = enumHelperMapInner
			enumValueDes := enumItem.GetValue()
			if enumValueDes != nil {
				enumValueSlice := make([]interface{}, 0, len(enumValueDes))
				for _, enumValueItem := range enumValueDes {
					enumValueMap := make(map[string]interface{})
					enumValueItemName := enumValueItem.GetName()
					enumValueItemNumber := strconv.Itoa(int(enumValueItem.GetNumber()))
					enumValueMap["name"] = enumValueItemName
					enumValueMap["number"] = enumValueItemNumber
					enumHelperMapInner[enumValueItemNumber] = enumValueItemName
					enumHelperMapInner[enumValueItemName] = enumValueItemNumber
					enumValueSlice = append(enumValueSlice, enumValueMap)
				}
				enumMap["value"] = enumValueSlice
			}
			enumSlice = append(enumSlice, enumMap)
		}
		result["enum"] = enumSlice
		result["enum_helper"] = enumHelperMap
	}

	// service setting
	serviceDes := fileDes.GetService()
	if serviceDes != nil {
		serviceSlice := make([]interface{}, 0, len(serviceDes))
		for _, service := range serviceDes {
			serviceMap := make(map[string]interface{})
			serviceMap["name"] = service.GetName()
			methodDes := service.GetMethod()
			if methodDes != nil {
				methodSlice := make([]interface{}, 0, len(methodDes))
				for _, method := range methodDes {
					methodMap := make(map[string]interface{})
					methodMap["name"] = method.GetName()

					methodMessageType := method.GetInputType()
					methodMap["input_type"] = methodMessageType
					index := strings.LastIndex(methodMessageType, ".")
					if index > -1 && index < len(methodMessageType)-1 {
						methodMessageType = methodMessageType[index+1:]
					}
					methodMap["input_type_short"] = methodMessageType

					methodMessageType = method.GetOutputType()
					methodMap["output_type"] = methodMessageType
					index = strings.LastIndex(methodMessageType, ".")
					if index > -1 && index < len(methodMessageType)-1 {
						methodMessageType = methodMessageType[index+1:]
					}
					methodMap["output_type_short"] = methodMessageType

					methodOpts := method.GetOptions()
					if methodOpts != nil {
						methodOptionsMap := map[string]interface{}{}
						methodOptionsMap["deprecated"] = strconv.FormatBool(methodOpts.GetDeprecated())
						methodOptionsMap["idempotency_level"] = methodOpts.GetIdempotencyLevel().String()
						handleOtherOptions(methodOpts.String(), result, methodOptionsMap)

						methodMap["options"] = methodOptionsMap
					}
					methodSlice = append(methodSlice, methodMap)
				}
				serviceMap["method"] = methodSlice
			}
			serviceSlice = append(serviceSlice, serviceMap)
		}
		result["service"] = serviceSlice
	}

	return result, nil
}

func handleOtherOptions(s string, dict map[string]interface{}, out map[string]interface{}) {
	opts := strings.Split(s, " ")
	if opts != nil {
		for _, opt := range opts {
			if opt == "" {
				continue
			}
			// TODO: 完善这段逻辑
			keyValuePair := strings.Split(opt, ":")
			if len(keyValuePair) != 2 {
				continue
			}
			m := getDeeperMap(dict, "extend")
			if m == nil {
				continue
			}
			m = getDeeperMap(m, "field_helper")
			if m == nil {
				continue
			}
			m = getDeeperMap(m, keyValuePair[0])
			if m == nil {
				continue
			}
			shortNameInterface, ok := m["type_short_name"]
			if !ok {
				continue
			}
			shortName, ok := shortNameInterface.(string)
			if !ok {
				continue
			}
			nameInterface, ok := m["name"]
			if !ok {
				continue
			}
			name, ok := nameInterface.(string)
			if !ok {
				continue
			}
			m = getDeeperMap(dict, "enum_helper")
			if m == nil {
				continue
			}
			m = getDeeperMap(m, shortName)
			if m == nil {
				continue
			}
			if realValue, ok := m[keyValuePair[1]]; ok {
				out[name] = realValue
			}
		}
	}
}

func getDeeperMap(m map[string]interface{}, key string) map[string]interface{} {
	inter, ok := m[key]
	if !ok {
		return nil
	}
	interMap, ok := inter.(map[string]interface{})
	if !ok {
		return nil
	}
	return interMap
}

func getFieldSlice(result map[string]interface{}, fields []*descriptorpb.FieldDescriptorProto) {
	if fields != nil {
		fieldSlice := make([]interface{}, 0, len(fields))
		fieldHelperMap := make(map[string]interface{})
		for _, field := range fields {
			fieldMap := make(map[string]interface{})
			fieldName := field.GetName()
			fieldNumber := strconv.Itoa(int(field.GetNumber()))
			fieldMap["name"] = fieldName
			fieldMap["number"] = fieldNumber
			fieldNumberHelperMap := make(map[string]interface{})
			fieldHelperMap[fieldNumber] = fieldNumberHelperMap
			fieldNumberHelperMap["name"] = fieldName

			fieldMap["label"] = field.GetLabel().String()
			fieldMap["type"] = field.GetType().String()
			fieldNumberHelperMap["label"] = field.GetLabel().String()
			fieldNumberHelperMap["type"] = field.GetType().String()
			if typeName := field.GetTypeName(); typeName != "" {
				fieldMap["type_name"] = typeName
				fieldNumberHelperMap["type_name"] = typeName
				splitIndex := strings.LastIndex(typeName, ".")
				shotTypeName := typeName
				if splitIndex >= 0 {
					shotTypeName = shotTypeName[splitIndex+1:]
				}
				fieldNumberHelperMap["type_short_name"] = shotTypeName
			}
			fieldMap["json_name"] = field.GetJsonName()
			fieldSlice = append(fieldSlice, fieldMap)
		}
		result["field"] = fieldSlice
		result["field_helper"] = fieldHelperMap
	}
}

func getShortName(name string) string {
	short := name

	index := strings.LastIndex(short, "/")
	if index >= 0 {
		short = short[index+1:]
	}
	index = strings.LastIndex(short, ".proto")
	if index >= 0 {
		short = short[:index]
	}
	return short
}

func (l Loader) LoadFromBytes(_ string, _ []byte) (map[string]interface{}, error) {
	return nil, ErrUnsupportedMethod
}
