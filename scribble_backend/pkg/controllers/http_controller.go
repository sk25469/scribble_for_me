package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomePageRoute(c *gin.Context) {
	fmt.Println("serving index file")
	http.ServeFile(c.Writer, c.Request, "../../static/index.html")
}
