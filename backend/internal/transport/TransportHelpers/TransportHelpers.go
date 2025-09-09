package TransportHelpers

import (
	"log"
	"net/http"
)

func CheckGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func CheckPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func CheckPut(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func CheckDelete(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func CheckJSON(w http.ResponseWriter, r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")

	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return false
	}
	return true
}

func FormJsonHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ProcessErrorCS(w http.ResponseWriter, err error, message string) bool {
	if err != nil {
		http.Error(w, message, http.StatusBadRequest)
		return false
	}
	return true
}

func ProcessErrorSS(err error, message string) bool {
	if err != nil {
		log.Println(message)
		return false
	}
	return true
}
