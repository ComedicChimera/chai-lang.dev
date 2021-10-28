package app

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	renderBase(c, "chai-lang.dev", "index.html", "index.scss", nil)
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
		"BookUnits": aside,
		"Prev":      nil,
		"Next": &bottomNav{
			DestName: "Hello World",
			Href:     "/unit1/chapter1",
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

func Chapter(c *gin.Context) {
	// build the aside
	aside, err := getAside()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// calculate the bottom page navigation
	chapterPath := c.Param("chapter-path")
	prevNav, nextNav, err := getBottomNav(aside, chapterPath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// load the markdown content
	bookPath := fmt.Sprintf("book/%s.md", chapterPath)
	mdContent, err := loadMarkdownTemplate("book.html", bookPath, map[string]interface{}{
		"BookUnits": aside,
		"Prev":      prevNav,
		"Next":      nextNav,
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

func getBottomNav(units []*bookUnit, chapterPath string) (*bottomNav, *bottomNav, error) {
	// split up the chapter path
	chapterPathElems := strings.Split(chapterPath, "/")[1:]
	unit, chapter := chapterPathElems[0], chapterPathElems[1]

	// get the unit number and chapter number
	unitN, err := strconv.Atoi(unit[4:])
	if err != nil {
		return nil, nil, err
	}

	chapterN, err := strconv.Atoi(chapter[7:])
	if err != nil {
		return nil, nil, err
	}

	// prev navigation
	var prevNav *bottomNav
	if chapterN == 1 {
		// if we are unit 1 => previous points to `Introduction`
		if unitN == 1 {
			prevNav = &bottomNav{
				Href:     "",
				DestName: "The Book",
			}
		} else {
			// otherwise, just go back a unit
			lastChapter := len(units[unitN-1].Chapters)
			prevNav = &bottomNav{
				Href:     fmt.Sprintf("/unit%d/chapter%d", unitN-1, lastChapter),
				DestName: units[unitN-2].Chapters[lastChapter].ChapterTitle,
			}
		}
	} else {
		// otherwise, just go back one chapter
		prevNav = &bottomNav{
			Href:     fmt.Sprintf("/unit%d/chapter%d", unitN, chapterN-1),
			DestName: units[unitN-1].Chapters[chapterN-2].ChapterTitle,
		}
	}

	// next navigation
	var nextNav *bottomNav
	if chapterN == len(units[unitN-1].Chapters) {
		// we only supply a next if we aren't at the end of the book
		if unitN < len(units) {
			nextNav = &bottomNav{
				Href:     fmt.Sprintf("/unit%d/chapter1", unitN+1),
				DestName: units[unitN-1].Chapters[0].ChapterTitle,
			}
		} else {
			nextNav = nil
		}
	} else {
		// otherwise just go forward one chapter
		nextNav = &bottomNav{
			Href:     fmt.Sprintf("/unit%d/chapter%d", unitN, chapterN+1),
			DestName: units[unitN-1].Chapters[chapterN].ChapterTitle,
		}
	}

	return prevNav, nextNav, nil
}

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
