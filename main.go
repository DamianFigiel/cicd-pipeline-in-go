package main

import (
	"cicd_pipeline_in_go/github_commits"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cicd-pipeline",
	Short: "A liteweight CI/CD pipeline tool written in Go",
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Triggert build based on new commits",
	Run: func(cmd *cobra.Command, args []string) {
		owner := os.Getenv("GITHUB_OWNER")
		repo := os.Getenv("GITHUB_REPO")
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" || owner == "" || repo == "" {
			fmt.Println("Please set GITHUB_OWNER, GITHUB_REPO and GITHUB_TOKEN")
			os.Exit(1)
		}

		commits, err := github_commits.GetCommits(owner, repo, token)
		if err != nil {
			fmt.Println("Error fetching commits: ", err)
			os.Exit(1)
		}

		fmt.Printf("Fetched %d commits from the reposotory %s/%s: \n", len(commits), owner, repo)
		for _, commit := range commits {
			fmt.Printf("- %s: %s\n", *commit.SHA, *commit.Commit.Message)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
