package printer

import (
	"fmt"
	"runtime"
)

// Hello is an exported function
func Hello() {
	fmt.Println("exported Hello")
}

// Version is an exported function
func Version() string {
	return runtime.Version()
}
