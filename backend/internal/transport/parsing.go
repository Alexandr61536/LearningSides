package transport

import (
	"encoding/json"
	"fmt"
	"log"
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

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var body RequestStructure

	err := json.NewDecoder(r.Body).Decode((&body))
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response_body := ResponseStructure{
		Value: fmt.Sprintf("Your name is %s and your age is %d", body.Name, body.Age),
	}

	response_body_marshaled, err := json.Marshal(response_body)

	if err != nil {
		fmt.Printf("Error marshaling response body")
		return
	}

	if _, err := w.Write(response_body_marshaled); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
