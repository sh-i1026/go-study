package usecase

import "net/http"

type AddressInput interface {
	GetAddress(w http.ResponseWriter, zipcode string) error
}
