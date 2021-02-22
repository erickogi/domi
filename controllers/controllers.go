package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// CanYouHearMeNow - Responds to liveliness checks
func CanYouHearMeNow(c *gin.Context) {
	c.Status(200)
}

// ReceiveGitHubWebHook - Receives and processes GitHub WebHook Events
func ReceiveGitHubWebHook(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}