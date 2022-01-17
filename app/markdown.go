package app

import (
	"bytes"
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

type mdHeading struct {
	HeadingText, HeadingTag string
}

// loadMarkdownContent loads a markdown file and converts it into HTML.
func loadMarkdownContent(markdownPath string) (string, error) {
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

	// convert the markdown
	mdParser := goldmark.New(goldmark.WithExtensions(extension.Table), goldmark.WithRendererOptions(html.WithUnsafe()))
	var mdBuff bytes.Buffer
	if err := mdParser.Convert(mdSrc, &mdBuff); err != nil {
		return "", err
	}
	mdHtml := mdBuff.String()

	// replace/update HTML tags (h1 is always converted to `section-title`)
	mdHtml = strings.ReplaceAll(mdHtml, "h1", "section-title")

	return mdHtml, nil
}

// loadMarkdownDocsTemplate loads a page of standard markdown documentation and
// fills in the appropriate HTML template with its content.
func loadMarkdownDocsTemplate(htmlTemplatePath, markdownPath string, contextVars map[string]interface{}) (string, error) {
	// load the markdown content
	mdHtml, err := loadMarkdownContent(markdownPath)
	if err != nil {
		return "", err
	}

	// isolate the headings
	var headings []*mdHeading
	re := regexp.MustCompile(`<h2>[^<]+</h2>`)
	for i, match := range re.FindAllString(string(mdHtml), -1) {
		headings = append(headings, &mdHeading{
			HeadingText: strings.TrimRight(match[4:len(match)-5], "\n"),
			HeadingTag:  fmt.Sprintf("section%d", i),
		})
	}

	// load the HTML template
	templ := template.Must(template.ParseFiles(filepath.Join(templateDir, htmlTemplatePath)))
	var htmlBuff bytes.Buffer
	contextVars["Content"] = template.HTML(mdHtml)
	contextVars["Headings"] = headings
	templ.Execute(&htmlBuff, contextVars)

	// return the template
	return htmlBuff.String(), nil
}
