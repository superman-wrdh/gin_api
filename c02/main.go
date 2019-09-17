package c02

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}
