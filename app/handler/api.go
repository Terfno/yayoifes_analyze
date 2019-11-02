package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/Terfno/yayoifes_analyze/app/domain"
)

func Read24(c *gin.Context) {
	log24, err := domain.Read24()
	if err != nil {
		log.Fatal("fail get")
	}

	jsonBytes, err := json.Marshal(log24)
	if err != nil {
		log.Fatal("fail marshal")
	}

	c.String(http.StatusOK, string(jsonBytes)+"\n")
}
