package usecase

import (
	"net/http"

	"github.com/study/go-study/clean-architecture/usecase/dto"
)

type AddressOutput interface {
	GetAddress(w http.ResponseWriter, dto *dto.GetAddressOutputDto) error
}
