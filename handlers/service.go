package handlers

import (
  "context"
  "github.com/annizhang/fetch-exercise/models"
)

type Service interface {
  AddTransaction(ctx context.Context, transaction models.Transaction)
  SpendPoints(ctx context.Context, spendRequest models.SpendRequest) *models.SpendResponse
  GetPoints(ctx context.Context) *models.AllPoints
}
