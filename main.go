package main

import (
	"cicd_pipeline_in_go/docker"
	"cicd_pipeline_in_go/github"
	"cicd_pipeline_in_go/kubernetes"
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
		// Github
		owner := os.Getenv("GITHUB_OWNER")
		repo := os.Getenv("GITHUB_REPO")
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" || owner == "" || repo == "" {
			fmt.Println("Please set GITHUB_OWNER, GITHUB_REPO and GITHUB_TOKEN")
			os.Exit(1)
		}

		commits, err := github.GetCommits(owner, repo, token)
		if err != nil {
			fmt.Println("Error fetching commits: ", err)
			os.Exit(1)
		}

		fmt.Printf("Fetched %d commits from the reposotory %s/%s: \n", len(commits), owner, repo)
		for _, commit := range commits {
			fmt.Printf("- %s: %s\n", *commit.SHA, *commit.Commit.Message)
		}

		cloneDir := os.Getenv("CLONE_DIR")
		if cloneDir == "" {
			cloneDir = "repoTmpDir"
		}
		fmt.Println("Cloning repository...")
		err = github.CloneRepo(owner, repo, token, cloneDir)
		if err != nil {
			fmt.Println("Error cloning repository: ", err)
			os.Exit(1)
		}

		// Docker
		fmt.Println("Building Docker image...")
		dockerfilePath := os.Getenv("DOCKERFILE_PATH")
		imageName := os.Getenv("IMAGE_NAME")
		if dockerfilePath == "" {
			dockerfilePath = fmt.Sprintf("%s/Dockerfile", cloneDir)
		}
		if imageName == "" {
			fmt.Println("Please set IMAGE_NAME in format <image>:<tag>")
			os.Exit(1)
		}

		err = docker.BuildDockerImage(dockerfilePath, imageName)
		if err != nil {
			fmt.Println("Error building Docker image: ", err)
			os.Exit(1)
		}

		fmt.Println("Docker image built successfully!")

		cleanUpRepo := os.Getenv("CLEANUP_REPO")
		if cleanUpRepo == "true" {
			fmt.Println("Cleaning up repository...")
			err = os.RemoveAll(cloneDir)
			if err != nil {
				fmt.Println("Error cleaning up repository: ", err)
				os.Exit(1)
			}
			fmt.Println("Repository cleaned up successfully!")
		}
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the Docker image to Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := os.Getenv("K8S_NAMESPACE")
		deploymentName := os.Getenv("K8S_DEPLOYMENT_NAME")
		newImage := os.Getenv("IMAGE_NAME")

		if namespace == "" || deploymentName == "" || newImage == "" {
			fmt.Println("Please set K8S_NAMESPACE, K8S_DEPLOYMENT_NAME and IMAGE_NAME")
			os.Exit(1)
		}

		clientset, err := kubernetes.GetKubeClient()
		if err != nil {
			fmt.Println("Error creating Kubernetes client: ", err)
			os.Exit(1)
		}

		err = kubernetes.UpdateDeployment(clientset, namespace, deploymentName, newImage)
		if err != nil {
			fmt.Println("Error updating Kubernetes deployment: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(deployCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
