package usecase

import (
	"log"
	"net/http"

	"github.com/study/go-study/clean-architecture/usecase/client"
	"github.com/study/go-study/clean-architecture/usecase/dto"
)

type AddressInteractor struct {
	addressOutPut  AddressOutput
	zipCloudClient client.ZipCloudClient
}

func NewAddressInteractor(
	addressOutPut AddressOutput,
	zipcloudClient client.ZipCloudClient,
) *AddressInteractor {
	return &AddressInteractor{
		addressOutPut:  addressOutPut,
		zipCloudClient: zipcloudClient,
	}
}

func (i *AddressInteractor) GetAddress(w http.ResponseWriter, zipCode string) error {
	addressInfo, err := i.zipCloudClient.GetAddressInfo(zipCode)
	if err != nil {
		log.Println("CLIENT ERROR")
		return err
	}
	err = i.addressOutPut.GetAddress(w, &dto.GetAddressOutputDto{
		AddresInfo: addressInfo,
	})
	return err
}
