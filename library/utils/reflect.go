package utils

import "encoding/json"

func DeepCopy(source interface{}, target interface{}) error {
	b, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, target)
	return err
}
