/*
 * The Mulato Library of Tests (MLT).
 *
 * Copyright (C) 2024 Alexandre Mulatinho <alex@mulatinho.net>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either  version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Mulato Library of Tests is a package with helper functions to Testing module
package mlt

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Assert tests an exp expression and returns error if not true
//
// Parameters:
//   - t: the
//   - exp: expression to be tested
func Assert(t *testing.T, expression bool) {
	_, file, line, _ := runtime.Caller(1)

	if !expression {
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
func Equals(t *testing.T, a, b any) {
	_, file, line, _ := runtime.Caller(1)

	if !reflect.DeepEqual(a, b) {
		fmt.Printf("\033[1mequals failed!\033[0m %s:%d (%v != %v)\n", filepath.Base(file), line, a, b)
		t.FailNow()
	}
}

// equals test if a is deep equal to b, returns error if not true
//
// Parameters:
//   - t: the
//   - a: first item to be tested
//   - b: second item to be tested
func StringEquals(t *testing.T, a, b string) {
	_, file, line, _ := runtime.Caller(1)

	if a != b {
		fmt.Printf("\033[1mstringEquals failed!\033[0m %s:%d (%v != %v)\n", filepath.Base(file), line, a, b)
		t.FailNow()
	}
}
