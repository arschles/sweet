package main

import (
	"testing"

	"github.com/arschles/sweet"
)

func TestSimpleStrings(t *testing.T) {
	ste := sweet.New("simple tests", t)
	ste.AddTest(TestASimpleString)
	ste.Run()
}
