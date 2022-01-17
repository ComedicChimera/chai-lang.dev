package app

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var tourChapterTitles []string

func Tour(c *gin.Context) {
	// get the chapter number parameter
	chapterNStr := c.Param("chapter")
	if chapterNStr == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("missing url parameter `chapter`"))
		return
	}

	// convert it to a number
	chapterN, err := strconv.Atoi(chapterNStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("url parameter `chapter` must be an integer"))
		return
	}

	// test that it is a valid chapter number
	if 1 > chapterN || chapterN > len(tourChapterTitles) {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid chapter number: `%d`", chapterN))
		return
	}

	// fetch the appropriate tour chapter
	chapter, err := loadTourChapter(chapterN)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// build the tour HTML template
	templ := template.Must(template.ParseFiles(filepath.Join(templateDir, "tour.html")))

	var htmlBuff bytes.Buffer
	templ.Execute(&htmlBuff, gin.H{
		"Chapter":       chapter,
		"ChapterIndex":  chapterN - 1, // index in the titles array
		"ChapterTitles": tourChapterTitles,
	})

	// compile that template into the base template
	c.HTML(http.StatusOK, "base.html", renderContent{
		Title:      "Tour | chai-lang.dev",
		Content:    template.HTML(htmlBuff.String()),
		PageStyle:  "tour.scss",
		Components: []string{"section-title"},
	})
}

// -----------------------------------------------------------------------------

type tourChapter struct {
	LessonContent              template.HTML
	TemplateCode, SolutionCode string
	ExpectedOutput             string
}

// loadTourTitles compiles the global list of tour titles to be given to the
// front end on every call to `Tour` so that an appropriate selection dropdown
// can be generated.
func loadTourTitles() {
	// calculate the tour path
	tourPath := filepath.Join(contentPath, "tour")

	// load the contents of the tour directory
	finfos, err := ioutil.ReadDir(tourPath)
	if err != nil {
		log.Fatalf("failed to read tour directory: %s\n", err.Error())
	}

	// walk each file in the tour directory and determine its chapter name
	for _, finfo := range finfos {
		// check to make sure we are only reading chapter directories
		if finfo.IsDir() && strings.HasPrefix(finfo.Name(), "chapter") {
			// determine the path to the lesson markdown file
			lessonMdFilePath := filepath.Join(tourPath, finfo.Name(), "lesson.md")

			// open the lesson markdown file
			lessonMdFile, err := os.Open(lessonMdFilePath)
			if err != nil {
				log.Fatalf("failed to open lesson markdown file: %s\n", err.Error())
			}
			defer lessonMdFile.Close()

			// read the first (title) line from the markdown file
			sc := bufio.NewScanner(lessonMdFile)
			sc.Split(bufio.ScanLines)

			firstLine := ""
			for sc.Scan() {
				firstLine = sc.Text()
			}

			if firstLine == "" {
				log.Fatalf("failed to read first line from `lesson.md` in chapter %s: %s\n", finfo.Name(), err.Error())
			}

			// extract the title from the first by trimming the leading `# ` and
			// add it to the slice of chapter titles
			tourChapterTitles = append(tourChapterTitles, firstLine[2:])
		}
	}
}

func loadTourChapter(chapN int) (*tourChapter, error) {
	// calculate the chapter path
	chapterPath := filepath.Join("tour", fmt.Sprintf("chapter%d", chapN))

	// load the markdown content from the chapter
	mdHtml, err := loadMarkdownContent(filepath.Join(chapterPath, "lesson.md"))
	if err != nil {
		return nil, err
	}

	// TODO: load the code and output

	// return the built tour chapter
	return &tourChapter{
		LessonContent: template.HTML(mdHtml),
		// TODO: rest
	}, nil
}
