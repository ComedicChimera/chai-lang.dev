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
		prev, next := getBottomNav(-1, 0, aside)

		// load the markdown content
		mdContent, err := loadMarkdownDocsTemplate("doc-group.html", filepath.Join(groupPath, "index.md"), map[string]interface{}{
			"GroupTitle": docName,
			"Aside":      aside,
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

// DocGroupSection returns a handler for retrieving a specific section of a
// specific chapter of a documentation group.
func DocGroupSection(docName, groupPath string) func(*gin.Context) {
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
		chapterN, err := strconv.Atoi(chapterPath[7:])
		if err != nil {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("invalid `chapter-path`: %s", chapterPath))
			return
		}

		// convert the section path into a section number
		sectionPath := c.Param("section-path")
		if sectionPath == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("missing url parameter `section-path`"))
			return
		}
		sectionN, err := strconv.Atoi(sectionPath[7:])
		if err != nil {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("invalid `section-path`: %s`", sectionPath))
		}

		// get the bottom navigation
		prev, next := getBottomNav(chapterN-1, sectionN-1, aside)

		// load the markdown content
		mdContent, err := loadMarkdownDocsTemplate("doc-group.html", filepath.Join(groupPath, fmt.Sprintf("chapter%d/section%d.md", chapterN, sectionN)), map[string]interface{}{
			"GroupTitle": docName,
			"Aside":      aside,
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

type groupAside struct {
	IndexTitle string
	IndexHref  string
	Chapters   []groupAsideChapter
}

type groupAsideChapter struct {
	Title    string
	ChapterN int
	Sections []groupAsideSection
}

type groupAsideSection struct {
	Title    string
	SectionN int
	Href     string
}

// getAside builds an aside for a specific documentation group.
func getAside(groupPath string) (*groupAside, error) {
	// load the `index.md` file and get its title
	indexTitle, err := getMdFileTitle(filepath.Join(groupPath, "index.md"))
	if err != nil {
		return nil, err
	}

	// determine the absolute group path
	groupAbsPath := filepath.Join(contentPath, groupPath)

	// build the index into the aside
	aside := &groupAside{IndexTitle: indexTitle, IndexHref: "/docs/" + groupPath}

	// open the chapter titles files
	chapterTitlesFile, err := os.Open(filepath.Join(groupAbsPath, "titles.txt"))
	if err != nil {
		return nil, err
	}

	// read all the chapter titles into a slice
	var chapterTitles []string
	titlesScanner := bufio.NewScanner(chapterTitlesFile)
	for titlesScanner.Scan() {
		chapterTitles = append(chapterTitles, titlesScanner.Text())
	}

	// close the titles file
	chapterTitlesFile.Close()

	// walk the group directory
	chapFInfos, err := ioutil.ReadDir(groupAbsPath)
	if err != nil {
		return nil, err
	}

	for _, chapFInfo := range chapFInfos {
		if chapFInfo.IsDir() && strings.HasPrefix(chapFInfo.Name(), "chapter") {
			// determine the chapter title
			chapterN, err := strconv.Atoi(chapFInfo.Name()[7:])
			if err != nil {
				return nil, err
			}

			chapter := groupAsideChapter{Title: chapterTitles[chapterN], ChapterN: chapterN}

			// walk the chapter directory
			sectionFInfos, err := ioutil.ReadDir(filepath.Join(groupAbsPath, chapFInfo.Name()))
			if err != nil {
				return nil, err
			}

			for i, sectionFInfo := range sectionFInfos {
				// ensure that we are only parsing md files
				if sectionFInfo.IsDir() || filepath.Ext(sectionFInfo.Name()) != ".md" || !strings.HasPrefix(sectionFInfo.Name(), "section") {
					continue
				}

				// extract the section title
				sectionTitle, err := getMdFileTitle(filepath.Join(groupPath, chapFInfo.Name(), sectionFInfo.Name()))
				if err != nil {
					return nil, err
				}

				// add the section to the chapter
				chapter.Sections = append(chapter.Sections, groupAsideSection{
					Title:    sectionTitle,
					SectionN: i + 1,
					// trim the extension
					Href: path.Join("/docs", groupPath, chapFInfo.Name(), sectionFInfo.Name()[:len(sectionFInfo.Name())-3]),
				})
			}

			// add the chapter to the aside
			aside.Chapters = append(aside.Chapters, chapter)
		}
	}

	return aside, nil
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
func getBottomNav(chapterNdx, sectionNdx int, aside *groupAside) (*groupAsideSection, *groupAsideSection) {
	var prev, next *groupAsideSection

	if chapterNdx == -1 {
		return nil, &aside.Chapters[0].Sections[0]
	}

	if sectionNdx == 0 {
		if chapterNdx == 0 {
			prev = &groupAsideSection{
				Title: aside.IndexTitle,
				Href:  aside.IndexHref,
			}
		} else {
			prev = &aside.Chapters[chapterNdx-1].Sections[len(aside.Chapters[chapterNdx-1].Sections)-1]
		}
	} else {
		prev = &aside.Chapters[chapterNdx].Sections[sectionNdx-1]
	}

	if sectionNdx == len(aside.Chapters[chapterNdx].Sections)-1 {
		if chapterNdx == len(aside.Chapters)-1 {
			next = nil
		} else {
			next = &aside.Chapters[chapterNdx+1].Sections[0]
		}
	} else {
		next = &aside.Chapters[chapterNdx].Sections[sectionNdx+1]
	}

	return prev, next
}
