package define

// Loader data loader, implement at least one method
// 数据加载器，只需要实现其中一个接口即可使用
type Loader interface {
	// LoadFromFile load data from file, input file path(absolute or related)
	// will try LoadFromBytes if failed
	LoadFromFile(filePath string) (map[string]interface{}, error)
	
	// LoadFromBytes load data from byte slice
	LoadFromBytes(filePath string, bytes []byte) (map[string]interface{}, error)
}
