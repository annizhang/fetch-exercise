# Fetch Rewards Exercise

## Requirements
go 1.17.6

## How to Run
    $ go build
    $ ./fetch-exercise

    Starting server at port 8080


## Implemented Endpoints:

#### Add Transactions
Add transaction for specific payer, transaction amount, and date
* **URL**

  `/add_transaction`

* **Method**

  `PUSH`

* **Request Body**
  - `payer` string
  - `points` int
  - `timestamp` string


* **Success Response**

  - Code: 201

#### Spend Points
Spend points using oldest payer points first and returns which points were spent from which payer
* **URL**

  `/spend_points`

* **Method**

  `POST`

* **Request Body**

  - `points` int


* **Success Response**

  - Code: 204 No Content

* **Error Response**

  - 400 Bad Request

#### Get Points
Get Points returns total payer points
* **URL**

  `/get_points`

* **Method**

  `GET`

* **Success Response**

  - Code: 200
  - Content:
```
  {
  "spent_points": [
    {
      "payer": "DANNON",
      "points": -100
    },
    {
      "payer": "UNILEVER",
      "points": -200
    },
    {
      "payer": "MILLER COORS",
      "points": -4700
    }
  ]
}
```

### Example Calls

### First add transactions

##### Requests:

```
curl -H "Content-Type: application/json" -X POST -d '{ "payer": "DANNON", "points": 1000, "timestamp": "2020-11-02T14:00:00Z" }' http://localhost:8080/add_transaction

curl -H "Content-Type: application/json" -X POST -d '{ "payer": "UNILEVER", "points": 200, "timestamp": "2020-10-31T11:00:00Z" }' http://localhost:8080/add_transaction

curl -H "Content-Type: application/json" -X POST -d '{ "payer": "DANNON", "points": -200, "timestamp": "2020-10-31T15:00:00Z" }' http://localhost:8080/add_transaction

curl -H "Content-Type: application/json" -X POST -d '{ "payer": "MILLER COORS", "points": 10000, "timestamp": "2020-11-01T14:00:00Z" }' http://localhost:8080/add_transaction

curl -H "Content-Type: application/json" -X POST -d '{ "payer": "DANNON", "points": 300, "timestamp": "2020-10-31T10:00:00Z" }' http://localhost:8080/add_transaction
```

##### Response:
```
{"message":"201 Created"}
{"message":"201 Created"}
{"message":"201 Created"}
{"message":"201 Created"}
{"message":"201 Created"}
```

### Spend points

##### Request:
```
curl -H "Content-Type: application/json" -X POST -d '{ "points": 5000 }' http://localhost:8080/spend_points
```

##### Response:

```
{
  "spent_points": [
    {
      "payer": "DANNON",
      "points": -100
    },
    {
      "payer": "UNILEVER",
      "points": -200
    },
    {
      "payer": "MILLER COORS",
      "points": -4700
    }
  ]
}
```

### Get points

##### Request:

```
curl -H "Content-Type: application/json" http://localhost:8080/points
```

##### Response:

```
{
  "points": {
    "DANNON": 1000,
    "MILLER COORS": 5300,
    "UNILEVER": 0
  }
}
```

## Notes for Improvement
- Add unit tests
- Wrap responses with status codes and messages
- Add persistent storage
