// Author: Ganesh Girase <ganesh.girase@gmail.com>
//
// This is an API package which contains operation related to API
// which includes API registration and handling various web request
// like GET, POST, PUT, DELTE

package api

import (
  "api/schema"
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "github.com/golang/glog"
  "io/ioutil"
  "lib"
  "net/http"
  u "util"
)


// Register set of API's to http router
// These API's may be picked from config.json
// Args:
//   router: Http Router type helps to route the
//           web request to it's receiver point.
//   Returns: None
func RegisterAPI(router *mux.Router) {
  // TODO: Load API's from config/api.json instead of
  // calling just one by one
  glog.Info("Registering API's")

  router.HandleFunc("/api/login", Login).Methods("GET")
  router.HandleFunc("/api/entertainment", GetEntertainmentData).Methods("GET")
  router.HandleFunc("/api/entertainment", SaveEntertainmentData).Methods("POST")
  glog.Info("Registering API's completed")

  return
}

// Gets the list of entertainment packages
// Args:
//   w: Write the response to writer
//   req: Http request
// Returns:
//   None
func GetEntertainmentData(w http.ResponseWriter, r *http.Request) {
  glog.Info("Receive request for GET entertainment data")

  hospital_id := r.FormValue("hospital_id")
  language := r.FormValue("lang")

  // Check for mandetory parameter hospital id.
  if hospital_id == "" {
    response := u.Message(false, "Query did not pass hospital_id param.")
    response["data"] = "error"
    u.Respond(w, response)
    return
  }
  // If Language option isn't passed from request, default would be english.
  if language == "" {
    language = "en"
  }
  glog.Info(fmt.Sprintf("Request Parameters: hospital_id: %s, Language: %s", hospital_id, language))

  sql := `SELECT
            a.category, a.title, a.provider_email, a.description
          FROM
            entertainment_data a, hospital_category b
          WHERE
            a.category = b.category AND b.hospital_id = $1
          ORDER BY
            a.category, a.title`
  glog.Info(sql)

  // Select entertainment data from database for given hospital
  records, err := lib.GetDB().Query(sql, hospital_id)
  u.ChkError(err)
  defer records.Close()

  var rows []schema.EntertainmentData //Struct schema to handle database rows
  var playDescription []schema.PlayDescription

  for records.Next() { // For every record
    // Associate every record to schema struct
    var row = schema.EntertainmentData{}
    err = records.Scan(&row.Category, &row.Title, &row.ProviderEmail, &row.PlayDescription)
    u.ChkError(err, "Unable to scan database records to schema")

    // Extract list of descriptions available from database record.
    description_list, err := json.Marshal(row.PlayDescription)
    u.ChkError(err, "Unable to convert json to schema format")

    err = json.Unmarshal([]byte(description_list), &playDescription)
    u.ChkError(err, "Unable to convert schema to json format")

    // Check if description is available in given language.
    row.Description = ""
    for _, playDesc := range playDescription {
      if playDesc.Language == language  {
        row.Description = playDesc.Description
      }
    }

    // This will ensure that PlayDescription isn't going in output data.
    row.PlayDescription = nil

    // Accumulate every record to final output.
    rows = append(rows, row)
  }
  glog.Info(rows)
  // Write data to output and send in response.
  response := u.Message(true, "Success")
  response["data"] = map[string]interface{}{"entertainment": rows}
  u.Respond(w, response)
}

// Save the data of entertainment package
// Args:
//   w: Write the response to writer
//   req: Http request
// Returns:
//   None
func SaveEntertainmentData(w http.ResponseWriter, req *http.Request) {
  glog.Info("Received SaveEntertainmentData request")

  // Read request data
  req_data, err := ioutil.ReadAll(req.Body)
  u.ChkError(err, "Unable to read http request body")
  glog.Info(fmt.Sprintf("Request Body: ", string(req_data)))

  // Associate request body data with database schema.
  var e schema.EntertainmentData
  err = json.Unmarshal(req_data, &e)

  sql := "INSERT INTO entertainment_data(category, title, provider_email, description) VALUES($1, $2, $3, $4)"
  glog.Info(sql)

  // Execute insert query to database to insert data.
  res, err := lib.GetDB().Exec(sql, e.Category, e.Title, e.ProviderEmail, e.PlayDescription)
  u.ChkError(err, "Unable to insert record in table entertainment_data")
  row_inserted, err := res.RowsAffected()
  glog.Info(fmt.Sprintf("Records inserted: %d", row_inserted))

  // Raise error if unable to insert record.
  if row_inserted < 1 {
    msg := "Cannot insert record in table entertainment_data"
    u.ChkError(err, msg)
  }

  // Write data to output and send in response.
  response := u.Message(true, "Success")
  u.Respond(w, response)
}

// Handle user login request and perform
// basic user authrization
// Args:
//   w: Write the response to writer
//   req: Http request
// Returns:
//   None

func Login(w http.ResponseWriter, req *http.Request) {
  // TODO: We can implemnt basic JWT authrization which will assign a uniqe
  // Session token id for every user session.
  // That will ensure basic security mechanisms.
  // Also password we can store as Base64 encoded.
  glog.Info("Received login request")
}

