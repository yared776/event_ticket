package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}
func Buy(c *gin.Context) {
	//TODO
	c.HTML(http.StatusOK, "buy.html", nil)
}
