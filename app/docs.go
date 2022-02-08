package app

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// DocPage returns a handler for a single-file documentation page.
func DocPage(mdPath string) func(*gin.Context) {
	return func(c *gin.Context) {
		// load the markdown content
		mdContent, err := loadMarkdownDocsTemplate("doc-page.html", mdPath, make(map[string]interface{}))

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

// -----------------------------------------------------------------------------

// DocGroupIndex returns a handler for the index of documentation group.
func DocGroupIndex(docName, groupPath string) func(*gin.Context) {
	return func(c *gin.Context) {
		// build the aside
		aside, err := getAside(groupPath)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// get the bottom navigation
		prev, next := getBottomNav(0, aside)

		// load the markdown content
		mdContent, err := loadMarkdownDocsTemplate("doc-group.html", filepath.Join(groupPath, "index.md"), map[string]interface{}{
			"GroupTitle": docName,
			"Chapters":   aside,
			"Prev":       prev,
			"Next":       next,
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// render it into the base template
		c.HTML(http.StatusOK, "base.html", renderContent{
			Title:      fmt.Sprintf("%s | chai-lang.dev", docName),
			Content:    template.HTML(mdContent),
			PageStyle:  "doc-group.scss",
			Components: []string{"section-title", "guide-exercise"},
		})
	}
}

// DocGroupChapter returns a handler for retrieving a specific chapter of a
// documentation group.
func DocGroupChapter(docName, groupPath string) func(*gin.Context) {
	// build the aside (before loading the chapters)
	aside, err := getAside(groupPath)
	if err != nil {
		log.Fatalf("failed to build aside for doc group `%s`: %s", docName, err.Error())
	}

	return func(c *gin.Context) {
		// convert the chapter path into a chapter number
		chapterPath := c.Param("chapter-path")
		if chapterPath == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("missing url parameter `chapter-path`"))
			return
		}
		chapterN, err := strconv.Atoi(chapterPath[8:])
		if err != nil {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("invalid `chapter-path`: %s", chapterPath))
			return
		}

		// get the bottom navigation; note that we do NOT need to subtract 1
		// from the chapterN because the "0th" chapter is always the index page.
		prev, next := getBottomNav(chapterN, aside)

		// load the markdown content
		mdContent, err := loadMarkdownDocsTemplate("doc-group.html", filepath.Join(groupPath, "chapters", fmt.Sprintf("chapter%d.md", chapterN)), map[string]interface{}{
			"GroupTitle": docName,
			"Chapters":   aside,
			"Prev":       prev,
			"Next":       next,
		})

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// render it into the base template
		c.HTML(http.StatusOK, "base.html", renderContent{
			Title:      fmt.Sprintf("%s | chai-lang.dev", docName),
			Content:    template.HTML(mdContent),
			PageStyle:  "doc-group.scss",
			Components: []string{"section-title", "guide-exercise"},
		})
	}
}

// -----------------------------------------------------------------------------

type groupAsideElem struct {
	Title string
	Href  string
}

// getAside builds an aside for a specific documentation group.
func getAside(groupPath string) ([]groupAsideElem, error) {
	// load the `index.md` file and get its title
	indexTitle, err := getMdFileTitle(filepath.Join(groupPath, "index.md"))
	if err != nil {
		return nil, err
	}

	// build the first element of the aside: the homepage
	elems := []groupAsideElem{{Title: indexTitle, Href: "/docs/" + groupPath}}

	// walk the group chapters
	finfos, err := ioutil.ReadDir(filepath.Join(contentPath, groupPath, "chapters"))
	if err != nil {
		return nil, err
	}

	// assumption: chapter files are in the right order when loaded from the OS
	// NOTE: we may have to revise this later if this turns out not to be
	// generally true.
	for _, finfo := range finfos {
		if finfo.IsDir() || filepath.Ext(finfo.Name()) != ".md" || !strings.HasPrefix(finfo.Name(), "chapter") {
			continue
		}

		chapterTitle, err := getMdFileTitle(filepath.Join(groupPath, "chapters", finfo.Name()))
		if err != nil {
			return nil, err
		}

		elems = append(elems, groupAsideElem{
			Title: chapterTitle,
			Href:  path.Join("/docs", groupPath, finfo.Name()[:len(finfo.Name())-3]), // trim the extension
		})
	}

	return elems, nil
}

// getMdFileTitle takes the path to an MD file relative to the content path and
// returns its title if possible.
func getMdFileTitle(path string) (string, error) {
	// open the file
	file, err := os.Open(filepath.Join(contentPath, path))
	if err != nil {
		return "", err
	}
	defer file.Close()

	// read the first line (must be title)
	sc := bufio.NewScanner(file)
	if !sc.Scan() {
		return "", errors.New("file ended without a title line")
	}
	titleLine := sc.Text()

	// trim off the `# ` before returning the title
	return titleLine[2:], nil
}

// -----------------------------------------------------------------------------

// getBottomNav returns the bottom navigation for a page in a documentation group.
func getBottomNav(i int, aside []groupAsideElem) (*groupAsideElem, *groupAsideElem) {
	var prev, next *groupAsideElem

	if i == 0 {
		prev = nil
	} else {
		prev = &aside[i-1]
	}

	if i == len(aside)-1 {
		next = nil
	} else {
		next = &aside[i+1]
	}

	return prev, next
}
