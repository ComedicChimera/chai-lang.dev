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
