package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const templateDir = "./templates/"

func Run(addr string) {
	router := gin.Default()

	// General Static File Handling
	router.Static("/static/images", "./static/images")
	router.Static("/static/components", "./static/components")
	router.Static("/static/vendor", "./static/vendor")
	router.Static("/static/js", "./static/js")
	router.Static("/static/demos", "./static/demos")

	// Sass Static Files
	createDistDir()
	router.GET("/static/scss/*file", loadSass)

	// Views
	router.LoadHTMLGlob(templateDir + "*")

	router.GET("/", Index)
	router.GET("/docs", Docs)
	router.GET("/docs/spec", DocGroupIndex("Language Specification", "spec"))
	router.GET("/docs/spec/*chapter-path", DocGroupChapter("Language Specification", "spec"))
	router.GET("/docs/tour", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/docs/tour/chapter1") })
	router.GET("/docs/tour/chapter:chapter", Tour)

	// Pre-run data loading (eg. tour chapter titles)
	loadTourTitles()

	router.Run(addr)
}
