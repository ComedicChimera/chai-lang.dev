package app

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wellington/go-libsass"
)

func loadSass(c *gin.Context) {
	// compute the file paths
	scssfilepath := filepath.Join("./static/scss/", c.Param("file"))
	distfilePath := filepath.Join("./static/dist/", strings.Replace(c.Param("file"), ".scss", ".css", 1))

	// check the ages of the files to determine if recompilation is necessary
	distfinfo, err := os.Stat(distfilePath)
	if err == nil {
		scssfinfo, err := os.Stat(scssfilepath)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		// if the dist file is newer than the scss file, just return the compile
		// file (save us the compilation)
		if distfinfo.ModTime().Unix() >= scssfinfo.ModTime().Unix() {
			c.File(distfilePath)
		}
	}

	// open the SCSS file
	f, err := os.Open(scssfilepath)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	defer f.Close()

	// compile it to CSS
	distfile, err := os.Create(distfilePath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer distfile.Close()

	comp, err := libsass.New(distfile, f)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err = comp.Run(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// serve the compiled file
	c.File(distfilePath)
}

func createDistDir() {
	_, err := os.Stat("./static/dist")
	if os.IsNotExist(err) {
		os.Mkdir("./static/dist", os.ModeDir)
	}
}
