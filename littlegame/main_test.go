package main

import (
	"testing"
	"time"

	"github.com/mulatinho/golabs/mlt"
)

func TestExampleGameTwo(t *testing.T) {
	timeStart := time.Now()
	ExampleGameTwo()
	timeEnd := time.Since(timeStart)
	mlt.Assert(t, timeEnd < (1_000_000_000*5))
}

func TestExampleGameOne(t *testing.T) {
	timeStart := time.Now()
	ExampleGameOne()
	timeEnd := time.Since(timeStart)
	mlt.Assert(t, timeEnd < (1_000_000_000*5))
}

func BenchmarExampleGameOne(b *testing.B) {
	ExampleGameOne()
}

func BenchmarExampleGameTwo(b *testing.B) {
	ExampleGameTwo()
}
