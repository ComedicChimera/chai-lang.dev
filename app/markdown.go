package app

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

const contentPath string = "./content"

type mdSection struct {
	SectionTitle, SectionTag string
}

func loadMarkdownTemplate(htmlTemplatePath, markdownPath string, contextVars map[string]interface{}) (string, error) {
	// load the markdown source
	mdFile, err := os.Open(filepath.Join(contentPath, markdownPath))
	if err != nil {
		return "", err
	}
	defer mdFile.Close()

	mdSrc, err := ioutil.ReadAll(mdFile)
	if err != nil {
		return "", err
	}

	// isolate the sections
	var sections []*mdSection
	re := regexp.MustCompile(`\n##[^#\n]+\n`)
	for i, match := range re.FindAllString(string(mdSrc), -1) {
		sections = append(sections, &mdSection{
			SectionTitle: strings.TrimRight(match[4:], "\n"),
			SectionTag:   fmt.Sprintf("section%d", i),
		})
	}

	// convert the markdown
	mdParser := goldmark.New(goldmark.WithExtensions(extension.Table), goldmark.WithRendererOptions(html.WithUnsafe()))
	var mdBuff bytes.Buffer
	if err := mdParser.Convert(mdSrc, &mdBuff); err != nil {
		return "", err
	}
	mdHtml := mdBuff.String()

	// replace/update HTML tags
	mdHtml = strings.ReplaceAll(mdHtml, "h1", "section-title")

	// load the html template
	templ := template.Must(template.ParseFiles(filepath.Join(templateDir, htmlTemplatePath)))
	var htmlBuff bytes.Buffer
	contextVars["Content"] = template.HTML(mdHtml)
	contextVars["Sections"] = sections
	templ.Execute(&htmlBuff, contextVars)

	// return the template
	return htmlBuff.String(), nil
}

type bookUnit struct {
	UnitTitle string
	Chapters  []*bookChapter
}

type bookChapter struct {
	ChapterTitle                 string
	ChapterUnitHref, ChapterHref string
}

// getAside loads the units of the book
func getAside() ([]*bookUnit, error) {
	// load the units file
	unitsFile, err := os.Open(filepath.Join(contentPath, "book/units.txt"))
	if err != nil {
		return nil, err
	}
	defer unitsFile.Close()

	// convert the units to a list
	unitContent, err := ioutil.ReadAll(unitsFile)
	if err != nil {
		return nil, err
	}
	unitTitles := strings.Split(string(unitContent), "\n")

	// walk the book directory and collect the units
	finfos, err := ioutil.ReadDir(filepath.Join(contentPath, "book"))
	if err != nil {
		return nil, err
	}

	var units []*bookUnit
	for _, finfo := range finfos {
		// we are going to naively assume there are less than 10 units
		if finfo.IsDir() && strings.HasPrefix(finfo.Name(), "unit") {
			// collect the chapters
			var chapters []*bookChapter
			chapterfinfos, err := ioutil.ReadDir(filepath.Join(contentPath, "book", finfo.Name()))
			if err != nil {
				return nil, err
			}

			for _, chapterfinfo := range chapterfinfos {
				// open the md file
				chaptermdFile, err := os.Open(filepath.Join(contentPath, "book", finfo.Name(), chapterfinfo.Name()))
				if err != nil {
					return nil, err
				}
				defer chaptermdFile.Close()

				// read the chapter title
				sc := bufio.NewScanner(chaptermdFile)
				var chapterTitleLine string
				if sc.Scan() {
					chapterTitleLine = sc.Text()
				} else {
					return nil, errors.New("chapter file is empty")
				}

				// trim the leading `#` to get the title
				chapters = append(chapters, &bookChapter{
					ChapterTitle:    chapterTitleLine[2:],
					ChapterUnitHref: finfo.Name(),
					ChapterHref:     strings.TrimSuffix(chapterfinfo.Name(), ".md"),
				})
			}

			// create the final unit
			units = append(units, &bookUnit{
				UnitTitle: unitTitles[len(units)],
				Chapters:  chapters,
			})

		}
	}

	return units, nil
}
