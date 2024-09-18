package inspect

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// assert tests an exp expression and returns error if not true
//
// Parameters:
//   - t: the
//   - exp: expression to be tested
func assert(t *testing.T, exp bool) {
	_, file, line, _ := runtime.Caller(1)

	if !exp {
		fmt.Printf("\033[1massert failed!\033[0m %s:%d\n", filepath.Base(file), line)
		t.FailNow()
	}
}

// equals test if a is deep equal to b, returns error if not true
//
// Parameters:
//   - t: the
//   - a: first item to be tested
//   - b: second item to be tested
func equals(t *testing.T, a, b any) {
	_, file, line, _ := runtime.Caller(1)

	if !reflect.DeepEqual(a, b) {
		fmt.Printf("\033[1equals failed!\033[0m %s:%d\n", filepath.Base(file), line)
		t.FailNow()
	}
}
