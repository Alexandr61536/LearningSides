package transport

import (
	"fmt"
	"net/http"
)

func Listener() {
	fmt.Println(">>> Listener started")
	http.HandleFunc("/", Root)
	http.HandleFunc("/echo", Echo)
	http.HandleFunc("/parsing", Parsing)

	fmt.Println("Listentig port 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting listening:", err)
	}
}
