package main

import (
	"fmt"

	"github.com/package/external"
)

func main() {
	fmt.Printf("This language is %v", external.ReturnLang())
}
