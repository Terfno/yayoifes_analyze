package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Terfno/yayoifes_analyze/app/domain"
)

func GetNumberOfVisitor(c *gin.Context) {
	log24, err := domain.GetNumberOfVisitor()
	if err != nil {
		log.Fatal("fail get")
	}

	c.String(http.StatusOK, strconv.Itoa(len(log24)))
}

func GetNumberOfVisitorPerHour(c *gin.Context) {
	var novperhour []int

	for i := 0; i < 24; i++ {
		perhour, err := domain.GetNumberOfVisitorByHour(i, i+1)
		if err != nil {
			log.Fatal("fail get")
		}
		fmt.Println("%d,%d,%d\n", perhour, i+1, i)
		novperhour = append(novperhour, len(perhour))
	}

	c.String(http.StatusOK, "%d\n", novperhour)
}

func Graph(c *gin.Context) {
	// number of visitor
	novi, err := domain.GetNumberOfVisitor()
	if err != nil {
		log.Fatal("fail get")
	}
	nov := strconv.Itoa(len(novi))

	// per hour slice
	var novph []int
	for i := 0; i < 24; i++ {
		perhour, err := domain.GetNumberOfVisitorByHour(i, i+1)
		if err != nil {
			log.Fatal("fail get")
		}
		novph = append(novph, len(perhour))
	}

	// avg
	var avg float64
	for _, value := range novph {
		avg += float64(value)
	}
	avg = avg / float64(len(novph))

	c.HTML(http.StatusOK, "index.html", gin.H{
		"nov":   nov,
		"datas": novph,
		"avg":   avg,
	})
}
