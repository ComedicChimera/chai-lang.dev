package app

import (
	"fmt"
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

type bottomNav struct {
	DestName string
	Href     string
}

func Book(c *gin.Context) {
	// build the aside
	aside, err := getAside()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// load the markdown content
	mdContent, err := loadMarkdownTemplate("book.html", "book/index.md", map[string]interface{}{
		"BookChapters": aside,
		"Prev":         nil,
		"Next": &bottomNav{
			DestName: "Hello World",
			Href:     "/chapter1/section1",
		},
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// render it into the base template
	c.HTML(http.StatusOK, "base.html", renderContent{
		Title:      "book | chai-lang.dev",
		Content:    template.HTML(mdContent),
		PageStyle:  "book.scss",
		Components: []string{"section-title"},
	})
}

func Section(c *gin.Context) {
	// build the aside
	aside, err := getAside()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// calculate the bottom page navigation
	sectionPath := c.Param("section-path")
	prevNav, nextNav, err := getBottomNav(aside, sectionPath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// load the markdown content
	bookPath := fmt.Sprintf("book/%s.md", sectionPath)
	mdContent, err := loadMarkdownTemplate("book.html", bookPath, map[string]interface{}{
		"BookChapters": aside,
		"Prev":         prevNav,
		"Next":         nextNav,
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// render it into the base template
	c.HTML(http.StatusOK, "base.html", renderContent{
		Title:      "book | chai-lang.dev",
		Content:    template.HTML(mdContent),
		PageStyle:  "book.scss",
		Components: []string{"section-title"},
	})
}

// -----------------------------------------------------------------------------

func docsPage(mdPath string) func(*gin.Context) {
	return func(c *gin.Context) {
		// load the markdown content
		mdContent, err := loadMarkdownTemplate("doc-page.html", mdPath, make(map[string]interface{}))

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// render it into the base content
		c.HTML(http.StatusOK, "base.html", renderContent{
			Title:      "docs | chai-lang.dev",
			Content:    template.HTML(mdContent),
			PageStyle:  "doc-page.scss",
			Components: []string{"section-title"},
		})
	}
}
