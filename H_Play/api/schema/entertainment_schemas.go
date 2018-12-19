//Author: Ganesh Girase<ganeshgirase@gmail.com>
// This schema package contains schema which
// serves dual purpose to accomodate web request data
// as well as entertainment_data table from database.

package schema

import "encoding/json"

type EntertainmentData struct {
  Title string `json:"title"`
  ProviderEmail string `json:"providerEmail"`
  Category string `json:"category"`
  PlayDescription *json.RawMessage `json:"i18n,omitempty"`
  Description string `json:"description"`// This field is used to marshal for specific language
}

type PlayDescription struct {
  Language string `json:"lang",omitempty`
  Description string `json:"description"`
}
