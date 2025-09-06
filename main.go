package main

import (
	"fmt"

	"github.com/z6wdc/go-escape-analysis/examples"
)

func main() {
	// This is the main entry point of the application.
	fmt.Println("returnValue:", examples.ReturnValue())
	fmt.Println("returnPointer:", examples.ReturnPointer())
}
