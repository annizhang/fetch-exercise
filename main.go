package main

import (
  "log"
  "fmt"
  "net/http"
  // "github.com/hashicorp/go-memdb"
  "github.com/julienschmidt/httprouter"
  "github.com/annizhang/fetch-exercise/handlers"
  "github.com/annizhang/fetch-exercise/service"
  "github.com/annizhang/fetch-exercise/models"
)

// func init() {
//
// }

func main() {

  router := httprouter.New()

  var transactions []models.Transaction
  allTransactions := models.AllTransactions{
    Transactions:transactions,
  }
  s := service.New(&allTransactions)

  fmt.Printf("making a new handler\n")
  h := handlers.NewHandler(s)
  router.POST("/add_transaction", h.AddTransaction)
	router.POST("/spend_points", h.SpendPoints)
  router.GET("/points", h.GetPoints)

  fmt.Printf("Starting server at port 8080\n")
  if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	log.Printf("...server stopping\n")

}
