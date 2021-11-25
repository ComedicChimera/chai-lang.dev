package app

import "github.com/gin-gonic/gin"

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
	router.GET("/docs/book", Book)
	router.GET("/docs/book/*section-path", Section)
	router.GET("/docs/module-schema", docsPage("module_schema.md"))

	router.Run(addr)
}
