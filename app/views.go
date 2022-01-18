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
	Title      string
	Content    template.HTML
	PageStyle  string
	Components []string
	CodeMirror bool
}

func loadViewContent(viewPath string) (string, error) {
	// open the view file
	viewFullPath := filepath.Join(htmlContentPath, viewPath)

	viewFile, err := os.Open(viewFullPath)
	if err != nil {
		return "", err
	}
	defer viewFile.Close()

	// read the view content as a string
	viewContent, err := ioutil.ReadAll(viewFile)
	if err != nil {
		return "", err
	}

	return string(viewContent), nil
}

func renderBase(c *gin.Context, title, viewPath, pageStyle string, components []string) {
	// read the view content
	viewContent, err := loadViewContent(viewPath)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	// serve it into the base template
	c.HTML(http.StatusOK, "base.html", renderContent{
		Title:      title,
		Content:    template.HTML(viewContent),
		PageStyle:  pageStyle,
		Components: components,
	})
}

func Index(c *gin.Context) {
	renderBase(c, "chai-lang.dev", "index.html", "index.scss", []string{"section-title"})
}

func Docs(c *gin.Context) {
	renderBase(c, "docs | chai-lang.dev", "docs.html", "docs.scss", []string{"section-title", "doc-card"})
}
