package lib

import (
	"context"
	"errors"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"github.com/bradleyfalzon/ghinstallation"
)

// GitHubProvider - Contains the components necessary for GitHub
type GitHubProvider struct {
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
	if ghPrivateKey != "" {
		gitHubProvider.githubPrivateKey = ghPrivateKey
	} else if ghToken != "" {
		gitHubProvider.githubToken = ghToken
	} else {
		return nil, errors.New("domi has no way to authenticate into GitHub")
	}
	return &gitHubProvider, nil
}

// GitHubTokenAuthenticator - Authenticates Personal Access Token
func GitHubTokenAuthenticator() *http.Client {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	return oauth2.NewClient(ctx, tokenSource)	
}