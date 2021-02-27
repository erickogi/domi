package lib

import (
	"context"
	"net/http"
	"os"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

// GitHubTokenAuthenticator - Authenticates Personal Access Token
func GitHubTokenAuthenticator() *http.Client {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	return oauth2.NewClient(ctx, tokenSource)	
}