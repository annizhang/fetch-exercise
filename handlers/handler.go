package handlers
import (
  "bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

  "github.com/annizhang/fetch-exercise/models"
  "github.com/julienschmidt/httprouter"
)

type Handler struct {
  s Service
}

func NewHandler(s Service) Handler {
  return Handler {
    s: s,
  }
}


func (h Handler) AddTransaction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

  fmt.Printf("In Handler AddTransaction\n")

  var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

  var request models.Transaction
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("error decoding transaction request: %+v\n", err)
		return
	}

  h.s.AddTransaction(context.Background(), request)

  w.WriteHeader(http.StatusOK)
  return
}

func (h Handler) SpendPoints(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Printf("in Handler SpendPoints\n")

  var request models.SpendRequest
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("error decoding spend request: %+v\n", err)
		return
	}

  spentPoints := h.s.SpendPoints(context.Background(), request)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(spentPoints)

  return
}

func (h Handler) GetPoints(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  points := h.s.GetPoints(context.Background())
  fmt.Println("all points: ", points)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(points)

  return
}
