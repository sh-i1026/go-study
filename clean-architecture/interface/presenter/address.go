package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/study/go-study/clean-architecture/usecase/dto"
)

type AddressPresenter struct{}

func NewAddressPresenter() *AddressPresenter {
	return &AddressPresenter{}
}

func (p *AddressPresenter) GetAddress(w http.ResponseWriter, in *dto.GetAddressOutputDto) error {
	output, err := json.MarshalIndent(in, "", "\t\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return nil
}
