package utils

import (
	"strings"
	"testing"

	"github.com/mulatinho/golabs/mlt"
)

func TestRandomNumber(t *testing.T) {
	inputs := []int{5, 25, 125, -1}
	for _, value := range inputs {
		var result int
		if value == -1 {
			result = GetRandomNumber()
			mlt.Assert(t, result < 10_000)
		} else {
			result = GetRandomNumber(value)
			mlt.Assert(t, result < value)
		}
		LogMessage(LOG_TYPE_DEBUG, result)
	}

}

func TestGetManName(t *testing.T) {
	manName := GetManName()
	mlt.Assert(t, manName != "")
	LogMessage(LOG_TYPE_DEBUG, manName)
}

func TestGetWomanName(t *testing.T) {
	womanName := GetWomanName()
	mlt.Assert(t, womanName != "")
	LogMessage(LOG_TYPE_DEBUG, womanName)
}

func TestGenerateName(t *testing.T) {
	fullNameOne := GenerateName(NAME_TYPE_MAN)
	fullNameTwo := GenerateName(NAME_TYPE_WOMAN)
	mlt.Assert(t, fullNameOne != fullNameTwo)
	mlt.Assert(t, strings.Contains(fullNameOne, " "))
	mlt.Assert(t, strings.Contains(fullNameTwo, " "))
	LogMessage(LOG_TYPE_DEBUG, fullNameOne)
	LogMessage(LOG_TYPE_DEBUG, fullNameTwo)
}

func TestTurnOrder(t *testing.T) {
	testsInput := []int{4, 8}

	for _, value := range testsInput {
		result := ApplyRandomIndexesToArray(value)
		mlt.Equals(t, len(result), value)
	}
}

func TestTurnOrderPreBuilt(t *testing.T) {
	testsInput := []int{4, 8}
	prebuilt := []int{3, 1}

	for _, value := range testsInput {
		result := ApplyRandomIndexesToArrayPreBuilt(value, prebuilt)
		mlt.Equals(t, len(result), value)
		mlt.Equals(t, result[0], prebuilt[0])
		mlt.Equals(t, result[1], prebuilt[1])
	}
}

func BenchmarkTurnOrderPreBuilt(b *testing.B) {
	testsInput := []int{8, 12, 26}
	prebuilt := []int{3, 1, 4}

	for _, value := range testsInput {
		_ = ApplyRandomIndexesToArrayPreBuilt(value, prebuilt)
	}
}
