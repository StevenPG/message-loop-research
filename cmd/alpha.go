package main

import (
	"fmt"
	"message-loop/pkg"
)

func main() {
	fmt.Println("mainTest")
	pkg.Function()
	// Start sending kafka messages in below function
	pkg.RunWebServer()
	// Listen for gRPC messages from beta
	// Send message to beta with protobuf
}
