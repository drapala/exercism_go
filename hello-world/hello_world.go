package greeting

import (
    "fmt"
)

// HelloWorld greets the world.
func HelloWorld() string {
    var s = "Hello, World!"
    fmt.Println(s)
	return s
}
