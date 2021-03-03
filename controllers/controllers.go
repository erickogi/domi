package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/devops-kung-fu/domi/lib"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/webhooks.v5/github"
)

// CanYouHearMeNow - Responds to liveliness checks
func CanYouHearMeNow(c *gin.Context) {
	c.Status(200)
}

// ReceiveGitHubWebHook - Receives and processes GitHub WebHook Events
func ReceiveGitHubWebHook(c *gin.Context) {
	githubProvider, err := lib.NewGitHubProvider()
	if err != nil {
		http.Error(c.Writer, "Could not get a provider.", 500)
		return
	}
	var hook *github.Webhook
	if githubProvider.GithubWebhookSecret != "" {
		hook, _ = github.New(github.Options.Secret(githubProvider.GithubWebhookSecret))
	} else {
		hook, _ = github.New()
	}	
	payload, err := hook.Parse(c.Request, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			c.String(http.StatusNotImplemented, "This event has not been implemented.")
		}
	}
	switch payload.(type) {
	case github.PushPayload:
		push := payload.(github.PushPayload)
		installationID := push.Installation.ID
		githubProvider.InstallationID = installationID
		_, err := githubProvider.GitHubAuthenticator()
		if err != nil {
			log.Println("GitHub Provider Authentication Failed")
			c.Error(errors.New("GitHub Provider Authentication Failed"))
		}
		log.Println("GitHub Provider Authentication Succeeded")
	}
	c.JSON(http.StatusOK, payload.(github.PushPayload))
}