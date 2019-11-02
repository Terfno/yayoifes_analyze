package handler

import (
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"

	"github.com/Terfno/yayoifes_analyze/app/domain"
)

func GetNumberOfVisitor(c *gin.Context) {
	log24, err := domain.GetNumberOfVisitor()
	if err != nil {
		log.Fatal("fail get")
	}

	c.String(http.StatusOK, strconv.Itoa((len(log24))))
}

func GetNumberOfVisitorPerHour(c *gin.Context) {
	var novperhour []int

	for i := 0; i < 24; i++ {
		perhour, err := domain.GetNumberOfVisitorByHour(i, i+1)
		if err != nil {
			log.Fatal("fail get")
		}
		novperhour = append(novperhour, len(perhour))
	}

	c.String(http.StatusOK, "%d\n", novperhour)
}
