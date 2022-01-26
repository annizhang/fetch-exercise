package service

import (
  "context"
  "fmt"
  "sort"
  "github.com/annizhang/fetch-exercise/models"
)

//Service is initiated with an empty slice/array to serve as a database to store transactions sorted from oldest to newest
type Service struct {
  db *models.AllTransactions
}

func New(db *models.AllTransactions) Service {
  return Service{
    db: db,
  }
}


//AddTransaction appends a new transaction to our in-memory and maintains order from oldest to newest
func (s Service) AddTransaction(ctx context.Context, transaction models.Transaction) {
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

//printTransactions
func (s Service) printTransactions() {
  for _, t := range s.db.Transactions {
    fmt.Printf("payer: %s, points: %d, timestamp: %s\n", t.Payer, t.Points, t.TimeStamp)
  }
}


//SpendPoints validates spend request and spends points from oldest to newest Transactions
//it also handles negative interactions
func (s Service) SpendPoints(ctx context.Context, spendRequest models.SpendRequest) (*models.SpendResponse, error) {
  points := spendRequest.Points

  //map to keep track of payers' points spent
  spentPoints := make(map[string]int)

  totalPoints, _ := s.getTotalPoints()
  if points > totalPoints {
    return nil, fmt.Errorf("Not enough points to spend")
  }

  for _, transaction := range s.db.Transactions {
    if points == 0 {
      break
    }

    //handling positive transaction points
    if transaction.Points > 0 {
      if transaction.Points >= points {

        //if payer is already in the map
        if _, ok := spentPoints[transaction.Payer]; ok {
          spentPoints[transaction.Payer] -= points
          transaction.Points -= points
          break
        }
        spentPoints[transaction.Payer] = -1 * points
        transaction.Points -= points
        break
      }
      // if transaction's points does not cover all of points to be spent
      points -= transaction.Points
      if _, ok := spentPoints[transaction.Payer]; ok {
        spentPoints[transaction.Payer] -= transaction.Points
        transaction.Points = 0
        continue
      }
      spentPoints[transaction.Payer] = -1 * transaction.Points
      transaction.Points = 0
    }

    //handling negative transaction points
    if transaction.Points < 0 {
      if _, ok := spentPoints[transaction.Payer]; ok {
        spentPoints[transaction.Payer] -= transaction.Points
        points -= transaction.Points
        transaction.Points = 0
      }
    }
  }

  return formatSpendResponse(spentPoints), nil
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

//getTotalPoints
func (s Service) getTotalPoints() (int, map[string]int) {
  totalPoints := 0

  //map of each payer to their total points
  pointsPerPayer := make(map[string]int)

  for _, transaction := range s.db.Transactions {
    totalPoints += transaction.Points

    //check if payer is already in map
    if _, ok := pointsPerPayer[transaction.Payer]; ok {
      pointsPerPayer[transaction.Payer] += transaction.Points
      continue
    }
    pointsPerPayer[transaction.Payer] = transaction.Points
  }

  return totalPoints, pointsPerPayer
}


//GetPoints calls internal function getTotalPoints to create a map of each payer to their total amount of points
func (s Service) GetPoints(ctx context.Context) *models.AllPoints {
  _, pointsPerPayer := s.getTotalPoints()

  return &models.AllPoints{
    Points: pointsPerPayer,
  }
}
