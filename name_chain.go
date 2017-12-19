package sweet

import (
	"fmt"
	"strings"
)

type nameChain struct {
	fmt.Stringer
	names []string
}

func newNameChain(names ...string) *nameChain {
	return &nameChain{
		names: names,
	}
}

func (n *nameChain) String() string {
	return strings.Join(n.names, " => ")
}

func (n *nameChain) append(name string) *nameChain {
	newArr := append(n.names, sanitizeTestName(name))
	return newNameChain(newArr...)
}

func sanitizeTestName(name string) string {
	ret := strings.Replace(name, "=>", "_", -1)
	return ret
}
