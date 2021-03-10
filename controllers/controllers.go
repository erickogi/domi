package controllers

import (
	// "context"
	"errors"
	"fmt"
	"log"
	"net/http"
	// "time"

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
	// ctx := context.Background()
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
		push := payload.(github.PushPayload)
		githubProvider.InstallationID = int64(push.Installation.ID)
		owner := push.Repository.Owner.Login
		repo := push.Repository.Name
		sha := push.After
		githubClient, err := githubProvider.GitHubAuthenticator()
		if err != nil {
			log.Println(errors.New("GitHub Provider Authentication Failed"))
		}
		log.Println("GitHub Provider Authentication Succeeded")
		archiveLink, _, err := githubClient.Repositories.GetArchiveLink(c, owner, repo, "zipball", &ghclient.RepositoryContentGetOptions{Ref: sha}, true)
		if err != nil {
			log.Println(err)
		}
		archiveURL := archiveLink.String()
		fs := lib.OSFS{}
		domiID, err := lib.DownloadFile(fs, archiveURL)
		if err != nil {
			log.Println(err)
		}
		unzipErr := lib.UnZip(fmt.Sprintf("/tmp/%s.zip", domiID), fmt.Sprintf("/tmp/%s", domiID))
		if unzipErr != nil {
			log.Println(err)
		}
		foundFiles, e := lib.FindFiles(fs, fmt.Sprintf("/tmp/%s", domiID), ".*\\.(tf|yaml|yml)")
		if e != nil {
			log.Println(e)
		}
		detailsURL := "https://example.com"
		// status := "queued"
		conclusion := "neutral"
		// startedAt := time.Now()
		// title := "domi - Policy-as-Code Enforcer"
		// summary := "Please stand by while we scan your repository... :thumbs-up:"
		// text := "Something can go here"
		if len(foundFiles) > 0 {
			_, _, checkError := githubClient.Checks.CreateCheckRun(c, owner, repo, ghclient.CreateCheckRunOptions{
				Name:	"domi",
				HeadSHA: sha,
				DetailsURL: &detailsURL,
				// Status: &status,
				Conclusion: &conclusion,
				// StartedAt: &ghclient.Timestamp{
				// 	Time: startedAt,
				// },
				// Output: &ghclient.CheckRunOutput{
				// 	Title: &title,
				// 	Summary: &summary,
				// 	Text: &text,
				// },
			})
			if checkError != nil {
				log.Println(checkError)
			}
		}
		c.String(http.StatusOK, "Payload Received")
	default:
		c.String(http.StatusNotImplemented, "Event Not Implemented")
	}
}