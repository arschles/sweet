package main

import (
	"github.com/arschles/sweet"
)

func getSimpleString() string {
	return "abc"
}

func appendToString(orig string, more string) string {
	return orig + "_" + more
}

func SimpleStringTest(ctx sweet.Context) {
	origStr := getSimpleString()
	newStr := appendToString(origStr, "another")
	ctx.Value(newStr).Equals(origStr+"_"+"another", "thing1 didn't match!")
}
