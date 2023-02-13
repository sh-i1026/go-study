package main

import (
	"fmt"
	"log"

	"github.com/init/init_test"
)

func main() {
	log.Println("This is main func")
	fmt.Print(init_test.ReturnHello() + "World")
}
