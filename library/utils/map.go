package utils

func MergeMap(target map[string]string, source map[string]string) {
	if source == nil {
		return
	}
	if target == nil {
		target = source
		return
	}

	for key, value := range source {
		target[key] = value
	}
}
