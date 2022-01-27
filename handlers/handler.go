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

// AddTransaction handles requests for adding transactions
func (h Handler) AddTransaction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

  var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

  var request models.Transaction
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("error decoding transaction request: %+v\n", err)
    w.WriteHeader(http.StatusBadRequest)
		return
	}

  h.s.AddTransaction(context.Background(), request)

  resp := make(map[string]string)
  resp["message"] = "201 Created"

  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(resp)

  return
}


//SpendPoints handles requests for spending points
func (h Handler) SpendPoints(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

  var request models.SpendRequest
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("error decoding spend request: %+v\n", err)
		return
	}

  spentPoints, err := h.s.SpendPoints(context.Background(), request)

  if err != nil {

    resp := make(map[string]string)
    resp["status"] = "400 BAD REQUEST"
    resp["message"] = "Spend Points Exceeds Total Points"

    fmt.Println("here0")

    w.WriteHeader(http.StatusBadRequest)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(spentPoints)

  return
}

//GetPoints handles requests for getting the points balance for the account
func (h Handler) GetPoints(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  points := h.s.GetPoints(context.Background())

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(points)

  return
}
