package rakuten

import "encoding/xml"

type response struct {
	XMLName      xml.Name      `xml:"XRoot"                                         json:"-"`
	Filename     string        `xml:"PRJ_LOYALTY_METADATA_HDR>FILENAME"             json:"filename,omitempty"`
	SID          string        `xml:"PRJ_LOYALTY_METADATA_HDR>SID"                  json:"sid,omitempty"`
	StartDate    string        `xml:"PRJ_LOYALTY_METADATA_HDR>START_DATE"           json:"start-date,omitempty"`
	EndDate      string        `xml:"PRJ_LOYALTY_METADATA_HDR>END_DATE"             json:"end-date,omitempty"`
	ReportType   string        `xml:"PRJ_LOYALTY_METADATA_HDR>REPORT_TYPE"          json:"report-type,omitempty"`
	Transactions []transaction `xml:"PRJ_LOYALTY_METADATA_HDR>PRJ_LOYALTY_U1_DAILY" json:"-"`
}

type transaction struct {
	MemberID            string `xml:"MEMBER_ID"             json:"member-id,omitempty"`
	MatchDate           string `xml:"MATCH_DATE"            json:"match-date,omitempty"`
	MatchTime           string `xml:"MATCH_TIME"            json:"match-time,omitempty"`
	MerchantPartnerID   string `xml:"MERCHANT_PARTNER_ID"   json:"merchant-partner-id,omitempty"`
	MerchantPartnerName string `xml:"MERCHANT_PARTNER_NAME" json:"merchant-partner-name,omitempty"`
	ProductSkuNumber    string `xml:"PRODUCT_SKU_NUMBER"    json:"product-sku-number,omitempty"`
	ProductName         string `xml:"PRODUCT_NAME"          json:"product-name,omitempty"`
	TransactionID       string `xml:"TRANSACTION_ID"        json:"transaction-id,omitempty"`
	OrderID             string `xml:"ORDER_ID"              json:"order-id,omitempty"`
	TransactionDate     string `xml:"TRANSACTION_DATE"      json:"transaction-date,omitempty"`
	TransactionTime     string `xml:"TRANSACTION_TIME"      json:"transaction-time,omitempty"`
	SalesAmount         string `xml:"SALES_AMOUNT"          json:"sales-amount,omitempty"`
	Quantity            string `xml:"QUANTITY"              json:"quantity,omitempty"`
	BaselineCommission  string `xml:"BASELINE_COMMISSION"   json:"baseline-commission,omitempty"`
	AdjustedCommission  string `xml:"ADJUSTED_COMMISSION"   json:"adjusted-commission,omitempty"`
	TotalCommission     string `xml:"TOTAL_COMMISSION"      json:"total-commission,omitempty"`
	ProcessDueDate      string `xml:"PROCESS_DUE_DATE"      json:"process-due-date,omitempty"`
	ProcessDueTime      string `xml:"PROCESS_DUE_TIME"      json:"process-due-time,omitempty"`
}
