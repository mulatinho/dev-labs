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

package mlt

import (
	"testing"
)

var arrayOfInts = []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
var arrayOfStrings = []string{
	"one", "two", "three", "four",
	"four", "three", "two", "one",
}

func TestAssertFunction(t *testing.T) {
	assert(t, 1 != 2)
	assert(t, 20%2 == 0)
	assert(t, 21%2 != 0)
}

func TestEqualsFunction(t *testing.T) {
	equals(t, arrayOfInts[0:1], arrayOfInts[:1])
	equals(t, arrayOfStrings[0:1], arrayOfStrings[:1])
	equals(t, 10000, 10_000)
}

func TestStringEqualsFunction(t *testing.T) {
	stringEquals(t, arrayOfStrings[0], "one")
	stringEquals(t, arrayOfStrings[2], arrayOfStrings[5])
}
