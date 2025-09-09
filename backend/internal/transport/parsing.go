package transport

import (
	"backend/internal/transport/TransportHelpers"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestStructure struct {
	Name string
	Age  int
}

type ResponseStructure struct {
	Value string
}

func Parsing(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if !TransportHelpers.CheckPost(w, r) {
		return
	}

	if !TransportHelpers.CheckJSON(w, r) {
		return
	}

	var body RequestStructure

	err := json.NewDecoder(r.Body).Decode((&body))

	if TransportHelpers.ProcessErrorCS(w, err, "Error reading request body") {
		return
	}

	TransportHelpers.FormJsonHeader(w, r)

	response_body := ResponseStructure{
		Value: fmt.Sprintf("Your name is %s and your age is %d", body.Name, body.Age),
	}

	response_body_marshaled, err := json.Marshal(response_body)

	if TransportHelpers.ProcessErrorSS(err, "Error marshaling response body") {
		return
	}

	devnull, err := w.Write(response_body_marshaled)

	devnull++

	TransportHelpers.ProcessErrorSS(err, fmt.Sprintf("Error writing response: %v", err))

}
