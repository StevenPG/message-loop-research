package pkg

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

// RunWebServerSendKafka - starts the alpha web server
func RunWebServerSendKafka() {
	r := gin.Default()
	r.GET("/start", func(c *gin.Context) {
		value := rand.Float64()
		message := fmt.Sprintf("Http-In -> %f", value)
		ProduceMessage(message)
		c.JSON(200, gin.H{
			"message": message,
			"ordinal": 0,
		})
		fmt.Println(message)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
