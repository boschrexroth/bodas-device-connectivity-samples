package main

import (
	"fmt"

	"boschrexroth.com/can-demo-golang/pkg/frontend"
)

func main() {
	fmt.Println("Serving frontend and REST API.")
	frontend.StartServer()
}
