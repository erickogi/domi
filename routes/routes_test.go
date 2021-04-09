package routes

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	SetupRouter()
}
