package sweet

import (
	"testing"
)

// Context holds all the information you need to run a test
type Context interface {
	// T() returns the internally held *testing.T. you can use this
	// value to integrate with other testing libraries
	T() *testing.T
	Value(val interface{}) Value
	// Name returns the fully-qualified name of the group.
	// for example, if you executed the following:
	//
	//	group := sweet.New("mytests").Group("subtests", false)
	//
	// then group.Name() returns "mytests.subtests"
	Name() string
}

type context struct {
	t    *testing.T
	path *nameChain
}

func newContext(t *testing.T, path *nameChain) *context {
	return &context{
		t:    t,
		path: path,
	}
}

func (c *context) T() *testing.T {
	return c.t
}

func (c *context) Value(val interface{}) Value {
	return newValue(c.t, c.path, val)
}

func (c *context) Name() string {
	return c.path.String()
}
