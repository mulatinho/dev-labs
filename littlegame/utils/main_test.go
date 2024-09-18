package utils

import (
	"strings"
	"testing"

	"github.com/mulatinho/golabs/mlt"
)

func TestGetManName(t *testing.T) {
	manName := GetManName()
	mlt.Assert(t, manName != "")
}

func TestGetWomanName(t *testing.T) {
	womanName := GetWomanName()
	mlt.Assert(t, womanName != "")
}

func TestGenerateName(t *testing.T) {
	fullNameOne := GenerateName(NAME_TYPE_MAN)
	fullNameTwo := GenerateName(NAME_TYPE_WOMAN)
	mlt.Assert(t, fullNameOne != fullNameTwo)
	mlt.Assert(t, strings.Contains(fullNameOne, " "))
	mlt.Assert(t, strings.Contains(fullNameTwo, " "))
}
