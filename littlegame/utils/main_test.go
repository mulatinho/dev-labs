package utils

import (
	"strings"
	"testing"
)

func TestGetManName(t *testing.T) {
	manName := GetManName()
	mlt.assert(t, manName != "")
}

func TestGetWomanName(t *testing.T) {
	womanName := GetWomanName()
	mlt.assert(t, womanName != "")
}

func TestGenerateName(t *testing.T) {
	fullNameOne := GenerateName(NAME_TYPE_MAN)
	fullNameTwo := GenerateName(NAME_TYPE_WOMAN)
	mlt.assert(t, fullNameOne != fullNameTwo)
	mlt.assert(t, strings.Contains(fullNameOne, " "))
	mlt.assert(t, strings.Contains(fullNameTwo, " "))
}
