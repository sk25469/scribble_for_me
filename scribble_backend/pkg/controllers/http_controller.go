package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomePageRoute(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "../../static/index.html")
}
