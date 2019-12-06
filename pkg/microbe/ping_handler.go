package microbe

import "github.com/gin-gonic/gin"

func (app *App) pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
