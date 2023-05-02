package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/study/go-study/clean-architecture/domain/entity"
)

const ZipCloudURL = "https://zipcloud.ibsnet.co.jp/api/search"

type ZipCloudImplement struct{}

func NewZipCloudImplement() *ZipCloudImplement {
	return &ZipCloudImplement{}
}

func (i *ZipCloudImplement) GetAddressInfo(zipCode string) (*entity.AddressInfo, error) {

	var addressInfo *entity.AddressInfo
	var response *entity.Response

	resp, err := http.Get(ZipCloudURL + "?zipcode=" + zipCode)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed Unmarshal")
		panic(err)
	}

	addressInfo = &entity.AddressInfo{
		Address1: response.Results[0].Address1,
		Address2: response.Results[0].Address2,
		Address3: response.Results[0].Address3,
		Kana1:    response.Results[0].Kana1,
		Kana2:    response.Results[0].Kana2,
		Kana3:    response.Results[0].Kana3,
		PrefCode: response.Results[0].PrefCode,
		ZipCode:  response.Results[0].ZipCode,
	}

	log.Println(addressInfo)

	return addressInfo, nil
}
