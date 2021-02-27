package lib

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/oauth2"
	"github.com/bradleyfalzon/ghinstallation"
)

// GitHubProvider - Contains the components necessary for GitHub
type GitHubProvider struct {
	GithubWebhookSecret		string
	githubToken				string
	githubPrivateKey		string
	appID					int
	installationID			int
	oauthClient				*http.Client
	githubClient			*http.Client
	installationTransport	*ghinstallation.Transport
}

// NewGitHubProvider - Starts a new GitHub Provider
func NewGitHubProvider() (*GitHubProvider, error) {
	gitHubProvider := GitHubProvider{}
	ghPrivateKey := os.Getenv("GITHUB_PRIVATE_KEY")
	ghToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	ghAppID, _ := strconv.Atoi(os.Getenv("GITHUB_APP_ID"))
	gitHubProvider.githubWebhookSecret = os.Getenv("GITHUB_WEBHOOK_SECRET")
	if ghPrivateKey != "" {
		gitHubProvider.githubPrivateKey = ghPrivateKey
		gitHubProvider.appID = ghAppID
	} else if ghToken != "" {
		gitHubProvider.githubToken = ghToken
	} else {
		return nil, errors.New("domi has no way to authenticate into GitHub")
	}
	return &gitHubProvider, nil
}

// GitHubAuthenticator - Authenticates Personal Access Token
func (githubProvider *GitHubProvider) GitHubAuthenticator() *http.Client {
	if githubProvider.githubToken != "" {
		ctx := context.Background()
		tokenSource := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubProvider.githubToken},
		)
		githubProvider.oauthClient = oauth2.NewClient(ctx, tokenSource)
		
	} else {

	}	
}