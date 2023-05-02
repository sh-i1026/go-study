package dto

import "github.com/study/go-study/clean-architecture/domain/entity"

type GetAddressInputDto struct {
	ZipCode string
}

type GetAddressOutputDto struct {
	AddresInfo *entity.AddressInfo
}
