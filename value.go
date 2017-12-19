package sweet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	IsBool(fmtString string, vals ...interface{})
}

type value struct {
	t    *testing.T
	orig interface{}
	path *nameChain
}

func newValue(t *testing.T, path *nameChain, val interface{}) Value {
	return &value{t: t, path: path, orig: val}
}

func (v *value) Equals(other interface{}, fmtString string, vals ...interface{}) {
	assert.Equal(v.t, v.orig, other, fmt.Sprintf(fmtString, vals...))
}

func (v *value) NotEqual(other interface{}, fmtString string, vals ...interface{}) {
	assert.NotEqual(v.t, v.orig, other, fmt.Sprintf(fmtString, vals...))
}
func (v *value) IsError(fmtString string, vals ...interface{}) {
	assert.Error(v.t, v.orig.(error), fmt.Sprintf(fmtString, vals...))
}
func (v *value) EqualsError(err error, fmtString string, vals ...interface{}) {
	assert.EqualError(v.t, err, err.Error(), fmt.Sprintf(fmtString, vals...))
}
func (v *value) IsNil(fmtString string, vals ...interface{}) {
	assert.Nil(v.t, v.orig, fmt.Sprintf(fmtString, vals...))
}
func (v *value) IsNotNil(fmtString string, vals ...interface{}) {
	assert.NotNil(v.t, v.orig, fmt.Sprintf(fmtString, vals...))
}
func (v *value) IsString(fmtString string, vals ...interface{}) {
	assert.IsType(v.t, "", v.orig, fmt.Sprintf(fmtString, vals...))
}
func (v *value) IsBool(fmtString string, vals ...interface{}) {
	assert.IsType(v.t, true, v.orig, fmt.Sprintf(fmtString, vals...))
}
