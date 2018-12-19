// Author: Ganesh Girase <ganeshgirase@gmail.com>

// This package contains all common utlities which
// can be used across all modules of hospital projects.
//

package util

import (
  "encoding/json"
  "io/ioutil"
  "github.com/golang/glog"
  "github.com/pkg/errors"
  "net/http"
)

// Method to handle error if any and exit in graceful way.
// Args:
//   err: Error of any type
//   msg: List of messages from caller.
// Returns:
//   None
// Raises:
//   Panic error if any error found.
func ChkError(err error, msg ...string) {
  if err != nil {
    err = errors.Wrapf(err, msg[0])
    glog.Error(err)
    panic(err)
  }
}

// This is generic method used to load any 
// single level (time being) json data to 
// generic map data structure.
// Args:
//   filePath: Configuration file absolute path
// Returns:
//   config: Map contains data loaded from config file.
func LoadConfig(configFile string) (config map[string]string, err error) {
  glog.Info("Lodig configuration")
   
  // Read configuration file content
  configFileContent, err := ioutil.ReadFile(configFile)
  ChkError(err, "Error reading config")

  // Unmarshal event consumer file content
  err = json.Unmarshal([]byte(configFileContent), &config)
  ChkError(err, "Failed to unmarshal config")
  return
}

// This will accomodate any process/function status into success 
// or failure map structure.
// Args:
//   status(bool): True or False
//   message: String messgage want to bind with structure.
// Returns:
//   map: a map data structure contaiing proces/function status.
func Message(status bool, message string) (map[string]interface{}) {
  return map[string]interface{} {"status" : status, "message" : message}
}

// Repond function builds Http response with json data
// and write to Http Response writer.
// Args:
//   w: Http Response writer
//   data: Data contains in map which gets converted in json format.
// Returns:
//   None
func Respond(w http.ResponseWriter, data map[string] interface{})  {
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode(data)
}

