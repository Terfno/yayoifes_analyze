package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Terfno/yayoifes_analyze/app/handler"
)

func main() {
	r := gin.Default()

	r.GET("/hello", handler.Hello)
	r.GET("/dbtest", handler.DBtest)

	r.Run(":8080")
}
