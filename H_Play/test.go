/// Testing package for Hospital entertainment project

package main

import (
    "bytes"
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
  u "util" 
)

// Function to test GET nature of entertainment package.
// This aims to verify whether api gives the correct data 
// for respective Hospital Id and language.
func TestGetEntertainmentData() {
    req, err := http.NewRequest("GET", "http://localhost:1234/api/entertainment?hospital_id=1&lang=de", nil)
    u.ChkError(err)
  
    // Requst client for hispital data for hospital id 
    client := &http.Client{}
    req.Header.Add("Accept", "application/json")

    resp, err := client.Do(req)
    if resp.StatusCode != 200 {
      u.ChkError(err, "GET operation failed!")
    }
    type Response struct {
      Status bool `json:"status"`
      Message string `json:"message"`
    }
    var response Response
    data, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal([]byte(data), &response)
    if (response.Status == true && response.Message == "Success") {
      fmt.Println("Test passed for GET operation")
    } else {
      fmt.Println("Test failed for GET operation")
      fmt.Println(response)
      fmt.Println(string(data))
    }
}

// Function to test POST nature of entertainment package.
// This test aims to verify if api is able to successfully
// store the data to database and retun success status in return

func TestPostEntertainmentData() {
 // Json data for testing 
 // test_json_1 := `{"title": "ARM", "providerEmail": "ard@mail.com", "category": "magzine_2", "i18n":[{ "lang":"en", "description":"ARM is a joint organisation of Germany’s regional public-service broadcasters"}, { "lang":"de", "description":"Wir sind eins."}]}`
 test_json_1 := `{"title": "WDR4", "providerEmail": "wdr2@mail.com", "category": "radio"}`

  url := "http://localhost:1234/api/entertainment"

  var jsonStr = []byte(test_json_1)
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  u.ChkError(err)
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    u.ChkError(err, "GET operation failed!")
  }
  type Response struct {
    Status bool `json:"status"`
    Message string `json:"message"`
  }
  var response Response
  data, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(data), &response)
  if (response.Status == true && response.Message == "Success") {
    fmt.Println("Test passed for POST operation")
  } else {
    fmt.Println("Test failed for POST operation")
    fmt.Println(response)
    fmt.Println(string(data))
  }
}

func main() {
  TestPostEntertainmentData()
  TestGetEntertainmentData()
}
