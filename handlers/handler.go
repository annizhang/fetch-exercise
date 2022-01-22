package handlers
import (
  memdb "github.com/hashicorp/go-memdb"
)

type handler struct {
  db memdb
}

func NewHandler(db memdb) handler {
  return handler {
    db: db,
  }
}

func (h handler) AddTransaction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

  var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

  var request models.TransactionRequest
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("error decoding transaction request: %+v\n", err)
		return
	}

  //no error decoding
  // add interaction to table


  w.WriteHeader(http:StatusOK)
  return
}
