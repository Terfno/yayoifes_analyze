package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Terfno/yayoifes_analyze/app/handler"
)

func main() {
	r := gin.Default()

	// try
	r.GET("/hello", handler.Hello)
	r.GET("/dbtest", handler.DBtest)

	// api
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/read24",handler.Read24)
	}

	r.Run(":8080")
}
