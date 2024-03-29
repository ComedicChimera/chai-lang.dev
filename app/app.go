package app

import (
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
	// router.GET("/docs/spec", DocGroupIndex("Language Specification", "spec"))
	// router.GET("/docs/spec/*chapter-path", DocGroupSection("Language Specification", "spec"))
	router.GET("/docs/book", DocGroupIndex("The Chai Book", "book"))
	router.GET("/docs/book/:chapter-path/:section-path", DocGroupSection("The Book", "book"))

	// API
	router.GET("/api/get-solution-code/*url-path", APIGetSolutionCode)

	router.Run(addr)
}
