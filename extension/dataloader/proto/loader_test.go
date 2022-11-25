package proto

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	l := NewLoader()
	fmt.Println(l.LoadFromFile("./account.proto"))
	// TODO: testing
}
