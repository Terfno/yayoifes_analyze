package main

import (
	"net/http"
	"path/filepath"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"

	"github.com/Terfno/yayoifes_analyze/app/handler"
)

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")

	// try
	r.GET("/hello", handler.Hello)
	r.GET("/dbtest", handler.DBtest)

	// ui
	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// api
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/nov",handler.GetNumberOfVisitor)
		apiRouter.GET("/novperhour",handler.GetNumberOfVisitorPerHour)
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
