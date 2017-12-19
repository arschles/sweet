package sweet

import (
	"fmt"
	"strings"
)

func addToNameChain(chain string, new string) string {
	return fmt.Sprintf("%s.%s", chain, sanitizeTestName(new))
}

func sanitizeTestName(name string) string {
	ret := strings.Replace(ret, ".", "_", -1)
	return ret
}
