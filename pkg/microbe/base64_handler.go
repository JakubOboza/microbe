package microbe

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	How to play with this API using curl
	encode:
		curl -H "Content-Type: application/json" -X POST -d '{"data":"some content to encode"}' http://localhost:8080/api/base64/encode

	decode:
	  curl -H "Content-Type: application/json" -X POST -d '{"data":"c29tZSBjb250ZW50IHRvIGVuY29kZQ=="}' http://localhost:8080/api/base64/decode
*/

type base64data struct {
	Data string `form:"data" json:"data" xml:"data"  binding:"required"`
}

func (app *App) encodeBase64Handler(c *gin.Context) {

	var input base64data

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sEnc := base64.StdEncoding.EncodeToString([]byte(input.Data))

	c.JSON(http.StatusOK, gin.H{
		"data": sEnc,
	})
}

func (app *App) decodeBase64Handler(c *gin.Context) {

	var input base64data

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sDec, err := base64.StdEncoding.DecodeString(input.Data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": string(sDec),
	})
}
