package routes

import (
	"github.com/devops-kung-fu/domi/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter - Set up gin router
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	github := r.Group("/github/v1")
	{
		github.POST("webhook", controllers.ReceiveGitHubWebHook)
	}
	return r
}