// Author: Ganesh Girase <ganeshgirase@gmail.com>
// All database related activities will be handled in this package
// including initiating new database connection, database query executions 
// and so on.

package lib

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  "github.com/golang/glog"
  "util"
)

const (
  // Json configuration file for database paramaters
  DatabaseConfigFile = "config/database.json"
)

// Package level database connector
var db *sql.DB

// New returns a DbConnection with the sql.DB set with the postgres
// DB connection string in the configuration
// Args: 
//   None
// Returns:
//   dbHandler: Pointer to database connection.
//   err: Error if any
func NewDbConnection() (dbHandler *sql.DB, err error) {
  glog.Info("Loading database configuration!!")
  config, err := util.LoadConfig(DatabaseConfigFile)
  util.ChkError(err, "Unable to load database config file")  
  glog.Info(config)
  
  // Open connection to postgres sql databases.
  dbHandler, err = sql.Open("postgres", fmt.Sprintf(
    "user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
    config["user"], config["password"], config["database"], config["host"], config["port"]))
  util.ChkError(err, "Couldn't open connection to postgres database!!")
  
  // Ping verifies if the connection to the database is alive or if a
  // new connection can be made.
  err = dbHandler.Ping()
  util.ChkError(err, "Couldn't ping posgres database!")
  db = dbHandler
  return
}

// Returns pointer database connector
func GetDB() *sql.DB {
  return db
}
