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
    // pretty simple - add a test :)
    //
    // All the tests will run serially, not concurrently.
    // But if you want to run tests concurrently, you can create a group.
    // Check out the AddGroup call below
    std.AddTest("my top-level tests", func(ctx sweet.Context) {
        myInt := 1
        added := add(myInt, 1)
        ctx.Value(added).Equals(added-1, "integers that were added together")
    })

    // Now, let's get fancy. This will create a group of tests that to run concurrently
    // (that's what the true is all about).
    //
    // You can always pass false to make these tests run serially too
    group := ste.Group("my other group of tests", true)
    group.AddTest("test the thing!", func(ctx sweet.Context) {
        myString := "mystring!"
        added := "mystring!!"
        ctx.Value(added).Equals(myString + "!", "the strings")
    })

    // Now, run the whole enchilada
    ste.Run()
}
```
