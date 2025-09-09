package main

import (
	"fmt"

	"backend/internal/transport"
)

func main() {
	fmt.Println("Starting...")

	transport.Listener()
}
