package handler

import (
	"encoding/binary"
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
		novperhour = append(novperhour, len(perhour))
	}

	c.String(http.StatusOK, "%d\n", novperhour)
}

func bytes2uint(bytes ...byte) uint64 {
    padding := make([]byte, 8-len(bytes))
    i := binary.BigEndian.Uint64(append(padding, bytes...))
    return i
}

func bytes2int(bytes ...byte) int64 {
    if 0x7f < bytes[0] {
        mask := uint64(1<<uint(len(bytes)*8-1) - 1)

        bytes[0] &= 0x7f
        i := bytes2uint(bytes...)
        i = (^i + 1) & mask
        return int64(-i)

    } else {
        i := bytes2uint(bytes...)
        return int64(i)
    }
}

func NewReception(c *gin.Context) {
	bage := c.PostForm("age")
	age := int(bytes2int(bage[0]))
	sex := c.PostForm("sex")

	err := domain.NewReception(age, sex)
	if err != nil {
		log.Fatal("fail insert")
	}

	c.Redirect(http.StatusFound, "/reception/")
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
