package main

import (
  "log"
  "net/http"
  "github.com/hashicorp/go-memdb"
  "github.com/julienschmidt/httprouter"
)

// func init() {
//
// }

func main() {

  router := httprouter.New()

  schema := &memdb.DBSchema{
    Tables: map[string]*memdb.TableSchema{
      "points": &memdb.TableSchema{
        Name: "points",
        Indexes: map[string]*memdb.IndexSchema{
          "payer": &memdb.IndexSchema{
            Name:    "payer",
            Unique:  true,
            Indexer: &memdb.StringFieldIndex{Field: "Payer"},
          },
          "points": &memdb.IndexSchema{
            Name:    "points",
            Unique:  false,
            Indexer: &memdb.IntFieldIndex{Field: "Points"},
          },
          "timestamp": &memdb.IndexSchema{
            Name:    "timestamp",
            Unique:  false,
            Indexer: &memdb.IntFieldIndex{Field: "Timestamp"},
          },
        },
      },
    },
  }

  if db, err := memdb.NewMemDB(schema); err != nil {
    log.Fatal("DBCreatiion:", err)
  }

  h := handlers.NewHandler(db)
  router.POST("/add_transaction", h.AddTransaction)
	// router.POST("/spend_points", h.SpendPoints)
  // router.GET("/points", h.GetPoints)

  if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	log.Printf("...server stopping\n")

}
