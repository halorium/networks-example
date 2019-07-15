package affiliatewindow

import "encoding/json"

type rawTransaction struct {
	AdvertiserCountry            string               `json:"advertiserCountry,omitempty"`
	AdvertiserID                 json.Number          `json:"advertiserId,omitempty"`
	Amended                      bool                 `json:"amended,omitempty"`
	AmendReason                  string               `json:"amendReason,omitempty"`
	AwinCommissionAmount         rawAmount            `json:"commissionAmount,omitempty"`
	AwinSaleAmount               rawAmount            `json:"saleAmount,omitempty"`
	ClickDate                    string               `json:"clickDate,omitempty"`
	ClickDevice                  string               `json:"clickDevice,omitempty"`
	ClickRefs                    rawClickRefs         `json:"clickRefs,omitempty"`
	CommissionSharingPublisherID string               `json:"commissionSharingPublisherId,omitempty"`
	CommissionStatus             string               `json:"commissionStatus,omitempty"`
	CustomParameters             []rawCustomParameter `json:"customParameters,omitempty"`
	DeclineReason                string               `json:"declineReason,omitempty"`
	ID                           json.Number          `json:"id,omitempty"`
	LapseTime                    json.Number          `json:"lapseTime,omitempty"`
	OldCommissionAmount          rawAmount            `json:"oldCommissionAmount,omitempty"`
	OldSaleAmount                rawAmount            `json:"oldSaleAmount,omitempty"`
	OrderRef                     string               `json:"orderRef,omitempty"`
	OriginalSaleAmount           string               `json:"originalSaleAmount,omitempty"`
	PaidToPublisher              bool                 `json:"paidToPublisher,omitempty"`
	PaymentID                    json.Number          `json:"paymentId,omitempty"`
	PublisherID                  json.Number          `json:"publisherId,omitempty"`
	PublisherURL                 string               `json:"publisherUrl,omitempty"`
	SiteName                     string               `json:"siteName,omitempty"`
	TransactionDate              string               `json:"transactionDate,omitempty"`
	TransactionDevice            string               `json:"transactionDevice,omitempty"`
	TransactionParts             []rawTransactionPart `json:"transactionParts,omitempty"`
	TransactionQueryID           json.Number          `json:"transactionQueryId,omitempty"`
	Type                         string               `json:"type,omitempty"`
	URL                          string               `json:"url,omitempty"`
	ValidationDate               string               `json:"validationDate,omitempty"`
	VoucherCodeUsed              bool                 `json:"voucherCodeUsed,omitempty"`
}

type rawAmount struct {
	Amount   json.Number `json:"amount,omitempty"`
	Currency string      `json:"currency,omitempty"`
}

type rawClickRefs struct {
	ClickRef  string `json:"clickRef,omitempty"`
	ClickRef2 string `json:"clickRef2,omitempty"`
	ClickRef3 string `json:"clickRef3,omitempty"`
	ClickRef4 string `json:"clickRef4,omitempty"`
	ClickRef5 string `json:"clickRef5,omitempty"`
	ClickRef6 string `json:"clickRef6,omitempty"`
}

type rawTransactionPart struct {
	CommissionGroupID json.Number `json:"commissionGroupId,omitempty"`
	Amount            json.Number `json:"amount,omitempty"`
}

type rawCustomParameter struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type transaction struct {
	AdvertiserCountry            string            `json:"advertiser-country,omitempty"`
	AdvertiserID                 string            `json:"advertiser-id,omitempty"`
	Amended                      string            `json:"amended,omitempty"`
	AmendReason                  string            `json:"amend-reason,omitempty"`
	AwinCommissionAmount         amount            `json:"awin-commission-amount,omitempty"`
	AwinSaleAmount               amount            `json:"awin-sale-amount,omitempty"`
	ClickDate                    string            `json:"click-date,omitempty"`
	ClickDevice                  string            `json:"click-device,omitempty"`
	ClickRefs                    clickRefs         `json:"click-refs,omitempty"`
	CommissionSharingPublisherID string            `json:"commission-sharing-publisher-id,omitempty"`
	CommissionStatus             string            `json:"commission-status,omitempty"`
	CustomParameters             []customParameter `json:"custom-parameters,omitempty"`
	DeclineReason                string            `json:"decline-reason,omitempty"`
	ID                           string            `json:"id,omitempty"`
	LapseTime                    string            `json:"lapse-time,omitempty"`
	OldCommissionAmount          amount            `json:"old-commission-amount,omitempty"`
	OldSaleAmount                amount            `json:"old-sale-amount,omitempty"`
	OrderRef                     string            `json:"order-ref,omitempty"`
	OriginalSaleAmount           string            `json:"original-sale-amount,omitempty"`
	PaidToPublisher              string            `json:"paid-to-publisher,omitempty"`
	PaymentID                    string            `json:"payment-id,omitempty"`
	PublisherID                  string            `json:"publisher-id,omitempty"`
	PublisherURL                 string            `json:"publisher-url,omitempty"`
	SiteName                     string            `json:"site-name,omitempty"`
	TransactionDate              string            `json:"transaction-date,omitempty"`
	TransactionDevice            string            `json:"transaction-device,omitempty"`
	TransactionParts             []transactionPart `json:"transaction-parts,omitempty"`
	TransactionQueryID           string            `json:"transaction-query-id,omitempty"`
	Type                         string            `json:"type,omitempty"`
	URL                          string            `json:"url,omitempty"`
	ValidationDate               string            `json:"validation-date,omitempty"`
	VoucherCodeUsed              string            `json:"voucher-code-used,omitempty"`
}

type amount struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type clickRefs struct {
	ClickRef  string `json:"click-ref,omitempty"`
	ClickRef2 string `json:"click-ref-2,omitempty"`
	ClickRef3 string `json:"click-ref-3,omitempty"`
	ClickRef4 string `json:"click-ref-4,omitempty"`
	ClickRef5 string `json:"click-ref-5,omitempty"`
	ClickRef6 string `json:"click-ref-6,omitempty"`
}

type transactionPart struct {
	CommissionGroupID string `json:"commission-group-id,omitempty"`
	Amount            string `json:"amount,omitempty"`
}

type customParameter struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
