package entity

type AddressInfo struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	Kana1    string `json:"kana1"`
	Kana2    string `json:"kana2"`
	Kana3    string `json:"kana3"`
	PrefCode string `json:"prefcode"`
	ZipCode  string `json:"zipcode"`
}

type Response struct {
	Message string        `json:"message"`
	Results []AddressInfo `json:"results"`
	Status  int           `json:"status"`
}
