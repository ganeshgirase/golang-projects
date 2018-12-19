package api


import (
  "fmt"
  "lib"
  "testing"
  "net/http"
  "github.com/stretchr/testify/assert"
  "github.com/gorilla/mux"
)

func TestAddition(t *testing.T) {
   db, _ := lib.NewDbConnection()
   fmt.Println(db)
   router := mux.NewRouter()
   router.HandleFunc("/api/login", Login).Methods("POST")
   url := "http://127.0.0.1:1234/api/login"
   response, _ := http.Get(url)
   fmt.Println(response)
   assert.Equal(t, 4, 5, "not good")
}
