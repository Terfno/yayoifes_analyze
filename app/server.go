package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"

	"github.com/Terfno/yayoifes_analyze/app/handler"
)

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")

	// ui
	r.GET("/", handler.Graph)
	r.GET("/reception", func(c *gin.Context) {
		c.HTML(http.StatusOK, "reception.html", gin.H{})
	})

	// api
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/nov", handler.GetNumberOfVisitor)
		apiRouter.GET("/novperhour", handler.GetNumberOfVisitorPerHour)
		apiRouter.GET("/1103", handler.Get1103)
		apiRouter.GET("/1104", handler.Get1104)
		apiRouter.POST("/newrecep/", handler.NewReception)
	}

	r.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		log.Fatal("html load err")
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		log.Fatal("html load err")
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
