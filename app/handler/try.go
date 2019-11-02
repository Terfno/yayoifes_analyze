package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/Terfno/yayoifes_analyze/app/domain"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!\n")

	return
}

func DBtest(c *gin.Context) {
	loger, err := domain.GetNumberOfVisitor()
	if err != nil {
		log.Fatal("fail get")
	}

	jsonBytes, err := json.Marshal(loger)
	if err != nil {
		log.Fatal("fail marshal")
	}

	c.String(http.StatusOK, string(jsonBytes)+"\n")

	return
}
