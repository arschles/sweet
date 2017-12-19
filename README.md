# Sweet

Sweet is an easy testing tool for Go apps.

# Sample Usage

```go
package mine

import (
    "github.com/arschles/sweet"
)

func TestAll(t *testing.T) {
    ste := sweet.New(t)
    std.AddTest(func(ctx sweet.Context) {
        myInt := getInt()
        added := add(myInt, 1)
        ctx.Value(myInt+1).Equals(added, "integers that were added together")
    })

    ste.RunAll()
}
```
