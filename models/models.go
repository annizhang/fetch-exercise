package models


type Transaction struct {
  Payer string `json:"payer"`
  Points int `json:"points"`
  TimeStamp int `json:"timestamp"`
}

type AllTransactions struct {
  Transactions []Transaction `json:"transactions"`
}
