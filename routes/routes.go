package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/plapi/controllers"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.POST("/subject", controllers.InsertSubject)
	router.GET("/subjects", controllers.IndexSubjects)
}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Welcome To API"})
	return
}
