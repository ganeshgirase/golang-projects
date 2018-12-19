// Author: Ganesh Girase <ganeshgirase@gmail.com>
//
// This project is about web application to deal with 
// entertainment packages specifically for Hospital
// Using this API, you can get the lits of entertainment packages
// available for various Hospitals. Also you can edit the list 
// of entertainment packages being running for different hospitals.

package main

import (
  "api"
  "flag"
  "github.com/gorilla/mux"
  "lib"
  "net/http"
u "util"
)

// This is the entry point for Hospital Entertainment package.
// Args: None
// Returns: None
func main() {
  // Setup database connection
  db, err := lib.NewDbConnection()
  u.ChkError(err, "Failed to get database connection") 

  // Close database connection at the end.
  defer db.Close()

  // Register for all API's through router
  router := mux.NewRouter()
  api.RegisterAPI(router)
 
  // TODO: This port can be picked up from configuration json file instead of hardcoding. 
  port := "1234"
  err = http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:12345/api
  u.ChkError(err)
}

// Function to handle command line parameters if any.
// Args: None
func init() {
  flag.Parse()
}
