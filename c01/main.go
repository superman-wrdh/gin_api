package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func index(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func getApi(c *gin.Context) {
	id := c.Query("id")
	if id == "0" {
		i, _ := strconv.Atoi(id)
		var x = 1 / i
		fmt.Println(x)
	}
	fmt.Println(id)
	c.String(http.StatusOK, "get api")
}

func postApi(c *gin.Context) {
	fmt.Println(c.PostForm("id"))
	c.String(http.StatusOK, "ok")
}

func postjson(c *gin.Context) {
	var data = &struct {
		Name string `json:"title"`
	}{}

	c.BindJSON(data)

	fmt.Println(data)
	c.String(http.StatusOK, "ok")

}

//全局中间件 允许跨域
func GlobalMiddleware(c *gin.Context) {
	fmt.Println("start middleware")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
	fmt.Println("finished middleware")
}

func AuthMiddleWare(c *gin.Context) {
	fmt.Println("AuthMiddleWare")
	fullPath := c.FullPath()
	fmt.Println("full path", fullPath)
	Authorization := c.GetHeader("Authorization")
	fmt.Println(Authorization)
	if Authorization == "Token 123456" {
		c.Next()
	} else {
		c.JSON(http.StatusForbidden, map[string]string{"message": "Authorization failed"})
		c.Abort()
	}
}

func ErrorMiddleware(c *gin.Context) {
	print("error middleware")
	c.Next()
}
func main() {
	r := gin.Default()
	r.Use(GlobalMiddleware, AuthMiddleWare)

	r.GET("/", index)
	r.GET("/getApi", getApi)
	r.POST("/postApi", postApi)
	r.POST("/postjson", postjson)
	r.Use(ErrorMiddleware)
	r.Run(":8080") // default listen and serve on 0.0.0.0:8080
}
