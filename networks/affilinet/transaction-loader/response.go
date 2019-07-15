package affilinet

import "encoding/xml"

type response struct {
	XMLName      xml.Name      `xml:"Envelope"                                                       json:"-"`
	TotalRecords int           `xml:"Body>GetTransactionsResponse>TotalRecords"                      json:"total-records"`
	Transactions []transaction `xml:"Body>GetTransactionsResponse>TransactionCollection>Transaction" json:"transactions"`
}

type transaction struct {
	BasketInfo          basketInfo   `xml:"BasketInfo"          json:"basket-info"`
	CancellationReason  string       `xml:"CancellationReason"  json:"cancellation-reason"`
	CheckDate           string       `xml:"CheckDate"           json:"check-date"`
	ClickDate           string       `xml:"ClickDate"           json:"click-date"`
	CreativeInfo        creativeInfo `xml:"CreativeInfo"        json:"creative-info"`
	NetPrice            string       `xml:"NetPrice"            json:"net-price"`
	ProgramID           string       `xml:"ProgramId"           json:"program-id"`
	ProgramTitle        string       `xml:"ProgramTitle"        json:"program-title"`
	PublisherCommission string       `xml:"PublisherCommission" json:"publisher-commission"`
	RateInfo            string       `xml:"RateInfo"            json:"rate-info"`
	RegistrationDate    string       `xml:"RegistrationDate"    json:"registration-date"`
	SubID               string       `xml:"SubId"               json:"sub-id"`
	TrackingMethod      string       `xml:"TrackingMethod"      json:"tracking-method"`
	TransactionID       string       `xml:"TransactionId"       json:"transaction-id"`
	TransactionStatus   string       `xml:"TransactionStatus"   json:"transaction-status"`
}

type basketInfo struct {
	BasketID             string `xml:"BasketId"             json:"basket-id"`
	OpenBasketItemCount  string `xml:"OpenBasketItemCount"  json:"open-basket-item-count"`
	RegisteredNetPrice   string `xml:"RegisteredNetPrice"   json:"registered-net-price"`
	TotalBasketItemCount string `xml:"TotalBasketItemCount" json:"total-basket-item-count"`
}

type creativeInfo struct {
	CreativeNumber string `xml:"CreativeNumber" json:"creative-number"`
	CreativeType   string `xml:"CreativeType"   json:"creative-type"`
}

type authentication struct {
	XMLName         xml.Name `xml:"Envelope"               json:"-"`
	CredentialToken string   `xml:"Body>CredentialToken" json:"credential-token"`
}
