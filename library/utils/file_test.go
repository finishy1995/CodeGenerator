package utils

import (
	"testing"
)

func TestGetAllFilesInDir(t *testing.T) {
	// TODO:
	result := GetAllFilesInDir("../")
	if len(result) < 5 {
		t.Fail()
	}
}
