package sweet

import (
	"testing"
)

// Value contains some data from a variable that you can test against other expected
// values. You create this inside a Context like so:
//
//	ctx.Value(actualValue).Equals("something else", "something didn't equal something else!")
//
type Value interface {
	Equals(other interface{}, fmtString string, vals ...interface{})
	NotEqual(other interface{}, fmtString string, vals ...interface{})
	IsError(fmtString string, vals ...interface{})
	EqualsError(err error, fmtString string, vals ...interface{})
	IsNil(fmtString string, vals ...interface{})
	IsNotNil(fmtString string, vals ...interface{})
	IsString(fmtString string, vals ...interface{})
	EqualsString(str string, fmtString string, vals ...interface{})
	IsBool(fmtString string, vals ...interface{})
	EqualsBool(b bool, fmtString string, vals ...interface{})
	MatchesRegex(regex string, fmtString string, vals ...interface{})
}

func newValue(t *testing.T, path *nameChain, val interface{}) Value {
	// TODO: implement this!
	return nil
}
