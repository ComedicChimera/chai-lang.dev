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

type bookChapter struct {
	ChapterTitle string
	Sections     []*bookSection
}

type bookSection struct {
	SectionTitle                    string
	SectionChapterHref, SectionHref string
}

// getAside loads the chapters of the book for constructing the side bar.
func getAside() ([]*bookChapter, error) {
	// load the chapters file
	chaptersFile, err := os.Open(filepath.Join(contentPath, "book/chapters.txt"))
	if err != nil {
		return nil, err
	}
	defer chaptersFile.Close()

	// convert the chapters to a list
	chapterContent, err := ioutil.ReadAll(chaptersFile)
	if err != nil {
		return nil, err
	}
	chapterTitles := strings.Split(string(chapterContent), "\n")

	// walk the book directory and collect the chapters
	finfos, err := ioutil.ReadDir(filepath.Join(contentPath, "book"))
	if err != nil {
		return nil, err
	}

	var chapters []*bookChapter
	for _, finfo := range finfos {
		// we are going to naively assume there are less than 10 chapters
		if finfo.IsDir() && strings.HasPrefix(finfo.Name(), "chapter") {
			// collect the sections
			var sections []*bookSection
			sectionfinfos, err := ioutil.ReadDir(filepath.Join(contentPath, "book", finfo.Name()))
			if err != nil {
				return nil, err
			}

			for _, sectionfinfo := range sectionfinfos {
				// open the md file
				sectionmdFile, err := os.Open(filepath.Join(contentPath, "book", finfo.Name(), sectionfinfo.Name()))
				if err != nil {
					return nil, err
				}
				defer sectionmdFile.Close()

				// read the section title
				sc := bufio.NewScanner(sectionmdFile)
				var sectionTitleLine string
				if sc.Scan() {
					sectionTitleLine = sc.Text()
				} else {
					return nil, errors.New("section file is empty")
				}

				// trim the leading `#` to get the title
				sections = append(sections, &bookSection{
					SectionTitle:       sectionTitleLine[2:],
					SectionChapterHref: finfo.Name(),
					SectionHref:        strings.TrimSuffix(sectionfinfo.Name(), ".md"),
				})
			}

			// create the final chapter
			chapters = append(chapters, &bookChapter{
				ChapterTitle: chapterTitles[len(chapters)],
				Sections:     sections,
			})

		}
	}

	return chapters, nil
}

// getBottomNav gets the page-bottom navigation for the book.
func getBottomNav(chapters []*bookChapter, sectionPath string) (*bottomNav, *bottomNav, error) {
	// split up the section path
	sectionPathElems := strings.Split(sectionPath, "/")[1:]
	chapter, section := sectionPathElems[0], sectionPathElems[1]

	// get the chapter and section number
	chapterN, err := strconv.Atoi(chapter[7:])
	if err != nil {
		return nil, nil, err
	}

	sectionN, err := strconv.Atoi(section[7:])
	if err != nil {
		return nil, nil, err
	}

	// prev navigation
	var prevNav *bottomNav
	if sectionN == 1 {
		// if we are chapter 1 => previous points to `Introduction`
		if chapterN == 1 {
			prevNav = &bottomNav{
				Href:     "",
				DestName: "The Book",
			}
		} else {
			// otherwise, just go back a section
			lastSection := len(chapters[chapterN-2].Sections)
			prevNav = &bottomNav{
				Href:     fmt.Sprintf("/chapter%d/section%d", chapterN-1, lastSection),
				DestName: chapters[chapterN-2].Sections[lastSection-1].SectionTitle,
			}
		}
	} else {
		// otherwise, just go back one section
		prevNav = &bottomNav{
			Href:     fmt.Sprintf("/chapter%d/section%d", chapterN, sectionN-1),
			DestName: chapters[chapterN-1].Sections[sectionN-2].SectionTitle,
		}
	}

	// next navigation
	var nextNav *bottomNav
	if sectionN == len(chapters[chapterN-1].Sections) {
		// we only supply a next if we aren't at the end of the book
		if chapterN < len(chapters) {
			nextNav = &bottomNav{
				Href:     fmt.Sprintf("/chapter%d/section1", chapterN+1),
				DestName: chapters[chapterN].Sections[0].SectionTitle,
			}
		} else {
			nextNav = nil
		}
	} else {
		// otherwise just go forward one section
		nextNav = &bottomNav{
			Href:     fmt.Sprintf("/chapter%d/section%d", chapterN, sectionN+1),
			DestName: chapters[chapterN-1].Sections[sectionN].SectionTitle,
		}
	}

	return prevNav, nextNav, nil
}
