package yaml

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var testYaml = `
Test:
  - A:
      B:
        C
      D:
        E
  - OK:
    - 1
    - 2
`

func TestLoader_LoadFromBytes(t *testing.T) {
	r := require.New(t)
	l := NewYamlLoader()
	r.NotNil(l)
	m, err := l.LoadFromBytes("./test.yaml", []byte(testYaml))
	r.Nil(err)
	r.NotNil(m)
	value, ok := m["Test"]
	r.True(ok)
	r.Equal(reflect.TypeOf(value).Kind(), reflect.Slice)
}
