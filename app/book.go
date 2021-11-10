package app

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

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

// getBottomNav gets the page-bottom navigation for the book.
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
			lastChapter := len(units[unitN-2].Chapters)
			prevNav = &bottomNav{
				Href:     fmt.Sprintf("/unit%d/chapter%d", unitN-1, lastChapter),
				DestName: units[unitN-2].Chapters[lastChapter-1].ChapterTitle,
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
				DestName: units[unitN].Chapters[0].ChapterTitle,
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
