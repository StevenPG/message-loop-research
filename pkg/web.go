package pkg

import "github.com/gin-gonic/gin"

// RunWebServer - starts the alpha web server
func RunWebServer() {
	r := gin.Default()
	r.GET("/start", func(c *gin.Context) {
		ProduceMessage("Http-In")
		c.JSON(200, gin.H{
			"message": "HTTP-In",
			"ordinal": 0,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
