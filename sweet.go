package sweet

import (
	"testing"
)

// Tests is the top level object to hold a collection of tests
type Tests interface {
	// Add a single test to this Suite. Tests will be executed serially.
	// If you'd like to execute tests concurrently, create a new group
	// with NewGroup, and pass true for the concurrent parameter.
	AddTest(t Test)
	// Group creates a test group called name. You can choose whether to run all the tests
	// inside this group concurrently. This function returns a new Suite object,
	// so you can create large trees of tests if you want.
	//
	// Groups will be executed serially in the top level
	Group(name string, concurrent bool) Tests
	// Run all the tests inside this Suite
	Run()
	// Name returns the name of these tests. Since you can arrange tests as a tree,
	// test names are fully-qualified. If you run this:
	//
	//	myGroupName := sweet.New("mytests", t).Group("myothertests", false).Name()
	//
	// then myGroupName will be "mytests.myothertests"
	Name() string
}

// New creates a new test suite
func New(name string, t *testing.T) Tests {
	return &tests{nameChain: name, t: t, concurrent: false}
}

type tests struct {
	nameChain     string
	t             *testing.T
	concurrent    bool
	topLevelTests []Test
	subGroups     []Tests
}

func (t *tests) Group(name string, concurrent bool) Tests {
	newGroup := &tests{
		nameChain:     addToNameChain(t.nameChain, name),
		t:             t.t,
		concurrent:    concurrent,
		topLevelTests: nil,
		subGroups:     nil,
	}
	t.subGroups = append(t.subGroups, newGroup)
	return newGroup
}

func (t *tests) AddTest(tst Test) {
	t.topLevelTests = append(t.topLevelTests, tst)
}

func (t *tests) Name() string {
	return t.nameChain
}

func (t *tests) Run() {
	// first run all individual tests serially
	for _, test := range t.topLevelTests {
		ctx := newContext(i, t.nameChain)
		test(ctx)
	}
	// then run groups serially
	for _, group := range t.subGroups {
		group.Run()
	}
}
