package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}