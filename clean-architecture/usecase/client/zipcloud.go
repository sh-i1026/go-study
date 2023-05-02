package client

import "github.com/study/go-study/clean-architecture/domain/entity"

type ZipCloudClient interface {
	GetAddressInfo(zipCode string) (*entity.AddressInfo, error)
}
