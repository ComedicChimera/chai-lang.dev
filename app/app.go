package app

import "github.com/gin-gonic/gin"

func Run(addr string) {
	router := gin.Default()

	// General Static File Handling
	router.Static("/static/images", "./static/images")
	router.Static("/static/components", "./static/components")

	// Sass Static Files
	createDistDir()
	router.GET("/static/scss/*file", loadSass)

	// Views
	router.LoadHTMLGlob("templates/*")

	router.GET("/", Index)
	router.GET("/docs", Docs)

	router.Run(addr)
}
