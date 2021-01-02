package main

import (
	"fmt"
	"message-loop/pkg"
)

func main() {
	fmt.Println("mainTest")
	pkg.Function()
	pkg.ConsumeMessage()
	// TODO - consume and send back to alpha using gRPC
}
