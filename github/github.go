package github

import (
	"context"
	"fmt"
	"os"
	"os/exec"

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

func CloneRepo(owner, repo, token string, targetDir string) error {
	repoUrl := fmt.Sprintf("https://%s@github.com/%s/%s.git", token, owner, repo)

	cmd := exec.Command("git", "clone", repoUrl, targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
