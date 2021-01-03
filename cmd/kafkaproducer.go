package main

import (
	"message-loop/pkg"
)

func main() {
	// Start sending kafka messages in below function
	pkg.RunWebServerSendKafka()
	// Listen for gRPC messages from beta
	// Send message to beta with protobuf
}
