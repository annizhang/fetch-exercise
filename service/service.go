package service

import (
  "context"
  "fmt"
  "sort"
  "github.com/annizhang/fetch-exercise/models"
  // memdb "github.com/hashicorp/go-memdb"
)

type Service struct {
  db *models.AllTransactions
}

func New(db *models.AllTransactions) Service {
  return Service{
    db: db,
  }
}

func (s Service) AddTransaction(ctx context.Context, transaction models.Transaction){
  fmt.Printf("in service addtransaction len %d\n", len(s.db.Transactions))

  s.db.Transactions = append(s.db.Transactions, transaction)
  // sort transactions by time stamp
  sort.Slice(s.db.Transactions, func(i, j int) bool {
    return s.db.Transactions[i].TimeStamp.Before(s.db.Transactions[j].TimeStamp)
  })

  fmt.Printf("transaction added len %d\n", len(s.db.Transactions))
  fmt.Printf("all transactions: \n")
  s.printTransactions()

  return
}

func (s Service) printTransactions() {
  for _, t := range s.db.Transactions {
    fmt.Printf("payer: %s, points: %d, timestamp: %s\n", t.Payer, t.Points, t.TimeStamp)
  }
}

func (s Service) SpendPoints(ctx context.Context, spendRequest models.SpendRequest) *models.SpendResponse {
  points := spendRequest.Points
  spentPoints := make(map[string]int)

  for _, transaction := range s.db.Transactions {
    if points == 0 {
      break
    }
    if transaction.Points > 0 {
      if transaction.Points >= points {
        if _, ok := spentPoints[transaction.Payer]; !ok {
          spentPoints[transaction.Payer] -= points
          transaction.Points -= points
          break
        }
        spentPoints[transaction.Payer] = -1 * points
        transaction.Points -= points
        break
      }
      // if t.Points < total poitns to be spentPoints
      points -= transaction.Points
      if _, ok := spentPoints[transaction.Payer]; !ok {
        spentPoints[transaction.Payer] = -1 * transaction.Points
        continue
      }
      spentPoints[transaction.Payer] -= transaction.Points
    }
    if transaction.Points < 0 {
      if _, ok := spentPoints[transaction.Payer]; ok {
        spentPoints[transaction.Payer] -= transaction.Points
        points += (-1 * transaction.Points)
        transaction.Points = 0
      }
    }
  }

  return formatSpendResponse(spentPoints)
}

func formatSpendResponse(spentPoints map[string]int) *models.SpendResponse {
  var paidPoints []models.PaidPoints
  for k,v := range spentPoints {
    paidPoints = append(paidPoints, models.PaidPoints{Payer: k, Points: v})
  }
  return &models.SpendResponse{
    SpentPoints: paidPoints,
  }
}

func (s Service) GetPoints(ctx context.Context) *models.AllPoints {
  allPoints := make(map[string]int)
  for _, transaction := range s.db.Transactions {
    if _, ok := allPoints[transaction.Payer]; ok {
      allPoints[transaction.Payer] += transaction.Points
      continue
    }
    allPoints[transaction.Payer] = transaction.Points
  }

  return &models.AllPoints{
    Points: allPoints,
  }
}
