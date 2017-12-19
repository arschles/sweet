package sweet

type Group interface {
	AddTest(t Test)
	RunsConcurrently() bool
}
