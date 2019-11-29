package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Terfno/yayoifes_analyze/app/domain"
)

func Graph(c *gin.Context) {
	// number of visitor
	log24, err := domain.GetNumberOfVisitor()
	if err != nil {
		log.Fatal("fail get")
	}
	nov := len(log24)

	// number of visitor yesterday
	novyi, err := domain.Read1103()
	if err != nil {
		log.Fatal("fail get")
	}
	novy := len(novyi)

	// number of visitor today
	novti, err := domain.Read1104()
	if err != nil {
		log.Fatal("fail get")
	}
	novt := len(novti)

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
		"novt":  novt,
		"novy":  novy,
		"datas": novph,
		"avg":   avg,
	})
}
