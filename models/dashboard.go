package models

type TrxSummary struct {
	TrashCategory    string `json:"trashCategory"`
	TransactionCount uint   `json:"transactionCount"`
}
