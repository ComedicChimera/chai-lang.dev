package app

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const htmlContentPath string = "./views/"

type renderContent struct {
	Title   string
	Content template.HTML
}

func renderBase(c *gin.Context, title, viewPath string) {
	// open the view file
	viewFullPath := filepath.Join(htmlContentPath, viewPath)

	viewFile, err := os.Open(viewFullPath)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	defer viewFile.Close()

	// read the view content as a string
	viewContent, err := ioutil.ReadAll(viewFile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// serve it into the base template
	c.HTML(http.StatusOK, "base.html", renderContent{
		Title:   title,
		Content: template.HTML(viewContent),
	})
}

func Index(c *gin.Context) {
	renderBase(c, "chai-lang.dev", "index.html")
}
