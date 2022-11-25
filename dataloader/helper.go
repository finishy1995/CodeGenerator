package dataloader

import (
	"CodeGenerator/dataloader/define"
	"CodeGenerator/library/log"
	"CodeGenerator/library/utils"
	"os"
	"strings"
)

var (
	loaderMap = map[string]define.Loader{}
)

// SetDataLoader set many data loader using map
// use filename suffix as map key, such as ".yaml", "yaml"
func SetDataLoader(m map[string]define.Loader) {
	loaderMap = m
}

// AddDataLoader add a data loader by filename suffix and its loader
// will delete the same key key-value pair
func AddDataLoader(fileSuffix string, loader define.Loader) {
	if loaderMap == nil {
		loaderMap = make(map[string]define.Loader)
	}
	loaderMap[fileSuffix] = loader
}

// AddDataLoaders add many data loader using map
// will delete the same key key-value pair
func AddDataLoaders(m map[string]define.Loader) {
	if loaderMap == nil {
		loaderMap = make(map[string]define.Loader)
	}
	for key, value := range m {
		loaderMap[key] = value
	}
}

// LoadFromDir load data from os directory
func LoadFromDir(dir string) map[string]interface{} {
	files := utils.GetAllFilesInDir(dir)
	result := map[string]interface{}{}
	for _, file := range files {
		MergeData(result, LoadFromFile(file))
	}

	return result
}

func LoadFromFile(path string) map[string]interface{} {
	var result map[string]interface{}
	var err error
	for suffix, loader := range loaderMap {
		if strings.HasSuffix(path, suffix) {
			result, err = loader.LoadFromFile(path)
			if err != nil {
				if err != ErrorNoNeedLog {
					log.Error("load file: %s failed, method: LoadFromFile, error: %s", path, err.Error())
				}

				bytes, err := os.ReadFile(path)
				if err != nil {
					log.Error("load file: %s failed, method: ReadFile, error: %s", path, err.Error())
				}
				result, err = loader.LoadFromBytes(path, bytes)
				if err != nil && err != ErrorNoNeedLog {
					log.Error("load file: %s failed, method: LoadFromBytes, error: %s", path, err.Error())
				}
			}
			break
		}
	}
	return result
}

func MergeData(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	var result map[string]interface{}
	err := utils.DeepCopy(m1, result)
	if err != nil {
		log.Error("DataLoader merge data failed, error: %s", err.Error())
		return m1
	}

	for key, value := range m2 {
		// TODO: for now, change duplicate value to the new one
		result[key] = value
	}

	return result
}
