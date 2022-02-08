package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func APIGetGuideExercise(c *gin.Context) {
	label, ok := c.GetQuery("label")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, errors.New("request missing required parameter `label`"))
		return
	}

	var sectionNum, exNum int
	_, err := fmt.Sscanf(label, "%d.%d", &sectionNum, &exNum)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to parse exercise label: %s", err))
		return
	}

	file, err := os.Open(filepath.Join(contentPath, fmt.Sprintf("guide/exercises/e%d_%d.json", sectionNum, exNum)))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to open exercise file: %s", err))
		return
	}
	defer file.Close()

	fdata, err := ioutil.ReadAll(file)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to read exercise file: %s", err))
		return
	}

	jdata := make(map[string]interface{})
	if err := json.Unmarshal(fdata, &jdata); err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to decode exercise file: %s", err))
		return
	}

	c.JSON(http.StatusOK, jdata)
}
