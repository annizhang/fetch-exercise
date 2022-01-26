package models

import (

  "time"
)

type Transaction struct {
  Payer string `json:"payer"`
  Points int `json:"points"`
  TimeStamp time.Time `json:"timestamp"`
}

type TransactionRequest struct {
  Payer string `json:"payer"`
  Points int `json:"points"`
  TimeStamp string `json:"timestamp"`
}

type AllTransactions struct {
  Transactions []Transaction `json:"transactions"`
}

type SpendRequest struct {
  Points int `json:"points"`
}

type PaidPoints struct {
  Payer string `json:"payer"`
  Points int `json:"points"`
}

type SpendResponse struct {
  SpentPoints []PaidPoints `json:"spent_points"`
}

type AllPoints struct {
  Points map[string]int `json:"points"`
}
