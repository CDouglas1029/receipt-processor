package models

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseData string `json:"purchaseData"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
