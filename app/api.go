package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const exerciseFileURLBase string = "https://raw.githubusercontent.com/ComedicChimera/chai/main/tests/suites/exercises/"

func APIGetSolutionCode(c *gin.Context) {
	urlPath, ok := c.Params.Get("url-path")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, errors.New("request missing required parameter `url-path`"))
		return
	}

	resp, err := http.Get(exerciseFileURLBase + urlPath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get file from github: %s", err))
		return
	}

	fileData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to read response from github: %s", err))
		return
	}

	c.String(http.StatusOK, string(fileData))
}
