package github_commits

import (
	"context"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func GetCommits(owner, repo, token string) ([]*github.RepositoryCommit, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return commits, nil
}
