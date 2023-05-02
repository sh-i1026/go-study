package controller

import (
	"log"
	"net/http"
	"path"

	uc "github.com/study/go-study/clean-architecture/usecase"
)

const ZipCodeLength = 7

type AddressController struct {
	AddressInput uc.AddressInput
}

func NewAddressController(input uc.AddressInput) *AddressController {
	return &AddressController{AddressInput: input}
}

func (c *AddressController) GetAddress(w http.ResponseWriter, r *http.Request) {
	zipCode := path.Base(r.URL.Path)
	if len(zipCode) != ZipCodeLength {
		w.WriteHeader(400)
		log.Printf("Invalid zipcode length : %v", zipCode)
		return
	}

	err := c.AddressInput.GetAddress(w, zipCode)
	if err != nil {
		log.Println("Input port is error")
		w.WriteHeader(500)
	}
}
