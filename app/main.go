package main

import (
	"fmt"
	"github.com/dudick123/toolkit"
)

func main() {
	// Create a new instance of the toolkit
	var tools toolkit.Tools

	s := tools.RandomString(10)
	fmt.Println("Random string:", s)
}
