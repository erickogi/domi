package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/devops-kung-fu/domi/lib"
	"github.com/gin-gonic/gin"
	ghclient "github.com/google/go-github/v33/github"
	"gopkg.in/go-playground/webhooks.v5/github"
)

// CanYouHearMeNow - Responds to liveliness checks
func CanYouHearMeNow(c *gin.Context) {
	c.Status(200)
}

// ReceiveGitHubWebHook - Receives and processes GitHub WebHook Events
func ReceiveGitHubWebHook(c *gin.Context) {
	ctx := context.Background()
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
	payload, err := hook.Parse(c.Request, github.PushEvent, github.CheckRunEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			c.String(http.StatusNotImplemented, "This event has not been implemented.")
		}
	}
	
	switch payload.(type) {
	case github.PushPayload:
		cCopy := c.Copy()
		go func() {
			push := payload.(github.PushPayload)
			log.Println("Anyone in here!?")
			// if check.InstallationID != "" {
			// 	installationID := check.InstallationID
			// 	githubProvider.InstallationID = installationID
			// }
			owner := push.Repository.Owner.Login
			repo := push.Repository.Name
			sha := push.Before
			githubClient, err := githubProvider.GitHubAuthenticator()
			if err != nil {
				log.Println("GitHub Provider Authentication Failed")
				cCopy.Error(errors.New("GitHub Provider Authentication Failed"))
			}
			log.Println("GitHub Provider Authentication Succeeded")
			archiveLink, _, err := githubClient.Repositories.GetArchiveLink(ctx, owner, repo, "zipball", &ghclient.RepositoryContentGetOptions{Ref: sha}, true)
			if err != nil {
				cCopy.Error(err)
			}
			archiveURL := archiveLink.String()
			fs := lib.OSFS{}
			domiID, err := lib.DownloadFile(fs, archiveURL)
			if err != nil {
				cCopy.Error(err)
			}
			unzipErr := lib.UnZip(fmt.Sprintf("/tmp/%s.zip", domiID), fmt.Sprintf("/tmp/%s", domiID))
			if unzipErr != nil {
				cCopy.Error(err)
			}				
		}()
		c.String(http.StatusOK, "Payload Received")
	default:
		c.String(http.StatusNotImplemented, "Event Not Implemented")
	}
}