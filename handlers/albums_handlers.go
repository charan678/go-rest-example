package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlbums(c *gin.Context) {
	//
	//if err := c.ShouldBind(&query); err != nil {
	//	for _, fieldErr := range err.(validator.ValidationErrors) {
	//		c.JSON(http.StatusBadRequest, fmt.Sprint(fieldErr))
	//		return // exit on first error
	//
	//	}
	//}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
