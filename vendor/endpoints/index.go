package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//Index View
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "main/index", gin.H{
		"title": "index",
	})
}
