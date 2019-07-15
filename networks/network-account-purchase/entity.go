package networkaccountpurchase

import (
	"fmt"
	"net/url"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/serializers"
)

type Head struct {
	CreatedAt        string `json:"created-at"`
	Etag             string `json:"etag"`
	NetworkAccountID string `json:"network-account-id"`
	NetworkID        string `json:"network-id"`
	PurchaseID       string `json:"purchase-id"`
	VariationVersion string `json:"variation-version"`
}

type Body struct {
	ClickID          string      `json:"click-id"`          // our click-id returned by network
	CommissionAmount string      `json:"commission-amount"` // transaction-commission-amount
	Currency         string      `json:"currency"`          // transaction-currency converted to our currency
	Date             string      `json:"date"`              // transaction-date
	OrderID          string      `json:"order-id"`          // merchant-transaction-id
	SaleAmount       string      `json:"sale-amount"`       // transaction-sale-amount
	Status           string      `json:"status"`            // transaction-status converted to our status
	Variant          string      `json:"variant"`
	Variations       interface{} `json:"variations,omitempty"`
}

type Entity struct {
	Head *Head `json:"head"`
	Body *Body `json:"body"`
}

func NewEntity() *Entity {
	return &Entity{
		Head: &Head{},
		Body: &Body{},
	}
}

func (e *Entity) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return flaw.New("wrong type").Wrap("cannot Scan")
	}

	flawError := serializers.JSONUnmarshalTag("json", source, e)
	if flawError != nil {
		return flawError.Wrap("cannot Scan")
	}

	return nil
}

// Prefix implements Entity interface
func (e *Entity) Prefix() string {
	return fmt.Sprintf("/networks/%s/accounts/%s/purchases", url.PathEscape(e.Head.NetworkID), url.PathEscape(e.Head.NetworkAccountID))
}

// ID implements Entity interface
func (e *Entity) ID() string {
	return e.Head.PurchaseID
}

// Etag implements Entity interface
func (e *Entity) Etag() string {
	return e.Head.Etag
}

// CreatedAt implements Entity interface
func (e *Entity) CreatedAt() string {
	return e.Head.CreatedAt
}

// DBJSON implements Entity interface
func (e *Entity) DBJSON() []byte {
	return serializers.JSONCompactSerializer(e)
}

// BodyJSON implements Entity interface
func (e *Entity) BodyJSON() []byte {
	return serializers.JSONCompactSerializer(e.Body)
}
