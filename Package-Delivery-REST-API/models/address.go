package models

type AddressData struct {
	StreetNumber int    `json:"streetnumber"`
	Street       string `json:"street"`
	State        string `json:"state"`
	City         string `json:"city"`
	Country      string `json:"country"`
	PostalCode   string `json:"postalcode"`
}
