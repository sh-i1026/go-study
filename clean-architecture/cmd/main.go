package main

import (
	"log"
	"net/http"

	"github.com/study/go-study/clean-architecture/registry"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	controller := registry.InitAddress()
	http.HandleFunc("/sample/", controller.GetAddress)

	log.Printf("server listen at %v", server.Addr)
	server.ListenAndServe()
}
