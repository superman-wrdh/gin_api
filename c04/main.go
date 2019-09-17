package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// Upload the file to specific dst.
		dst := "./upload/" + file.Filename
		//dst := file.Filename
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, map[string]string{"fileName": file.Filename})
		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	fmt.Print("start server http://127.0.0.1:8080")
	router.Run(":8080")
}
