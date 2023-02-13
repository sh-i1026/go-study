package init_test

import (
	"log"
)

func init() {
	log.Println("This is init func")
}

func ReturnHello() string {
	return "Hello"
}
