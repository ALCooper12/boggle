# boggle-web-service

This is a web service for solving Boggle boards. It allows users to submit a Boggle board via an API call and returns all valid words found on the board.

## Features

- Accepts a Boggle board as input via a POST request
- Solves the Boggle board and returns all valid words
- Uses a predefined dictionary to validate words

## Getting Started 

### Prerequisites
Make sure the following are installed first:

- [Go](https://golang.org/doc/install) (I used version 1.21)
- [Gin](https://github.com/gin-gonic/gin) package `go get -u github.com/gin-gonic/gin`

### Running the service
Start the web service by running the following command:
```
go run main.go
```
Make POST requests in a different terminal like the following example input and response below (taken from [wikapedia](https://en.wikipedia.org/wiki/Boggle)):
```
arianna@Ariannas-MBP % curl -X POST http://localhost:8080/solveBoggleBoard -d '{"board":[["r","h","r","e"],["y","p","c","s"],["w","n","s","n"],["t","e","g","o"]]}' -H "Content-Type: application/json"

{"words":["chypres","cress","cresset","set","song"]}% 
```
