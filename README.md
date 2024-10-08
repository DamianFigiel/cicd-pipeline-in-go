# CI/CD Pipeline in Go with GitHub, Docker, and Kubernetes Integration

This project implements a lightweight CI/CD pipeline in Go that automates builds and deployments by integrating with the GitHub API, Docker API, and Kubernetes API. The pipeline can fetch commits from a GitHub repository, clone GitHub repository, build Docker images, and deploy them to a Kubernetes cluster using rolling updates.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Usage](#usage)
  - [Build](#build)
  - [Deploy](#deploy)
- [Environment Variables](#environment-variables)
- [Extending the Project](#extending-the-project)
- [Contributing](#contributing)
- [License](#license)

## Overview

This project demonstrates a fully automated CI/CD pipeline that:

- Fetches the latest commits from a GitHub repository.
- Clones the repository and builds a Docker image using the repository's Dockerfile.
- Deploys the Docker image to a Kubernetes cluster using rolling updates.

The project is designed to be modular, allowing developers to extend it by adding new features or integrating with additional tools.

## Features

- **GitHub Integration**: Automatically fetches commits from a repository.
- **Docker API Integration**: Builds Docker images based on the repository's Dockerfile.
- **Kubernetes API Integration**: Deploys the Docker image to a Kubernetes cluster using rolling updates.
- **Rolling Updates**: Updates existing Kubernetes deployments without downtime.

## Prerequisites

Before you start, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.18+)
- [Docker](https://docs.docker.com/get-docker/)
- [Kubernetes CLI (kubectl)](https://kubernetes.io/docs/tasks/tools/)
- Access to a Kubernetes cluster (local or remote)
- GitHub Personal Access Token (PAT) with the necessary permissions

## Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/your-repo-name
   cd your-repo-name
2. **Install dependencies**:
The project uses Go modules. Run the following command to download dependencies:
    ```bash
    go mod tidy
3. **Set up environment variables**:
    ```bash
    export GITHUB_TOKEN=your_personal_access_token
## Usage
### Build
To fetch commits from a GitHub repository, clone the repository, and build a Docker image:
    ```bash
    go run main.go build
This will:
1. Fetch the latest commits from your GitHub repository.
2. Clone the repository into a temporary directory.
3. Build a Docker image using the Dockerfile in the cloned repository.

### Deploy
To deploy the newly built Docker image to a Kubernetes cluster:
    ```bash
    go run main.go deploy
This will:
1. Connect to your Kubernetes cluster.
2. Update an existing Kubernetes deployment with the new image.
3. Perform a rolling update, ensuring zero downtime.

### Environment Variables
Here is a list of all the environment variables used by the project:
- GITHUB_OWNER:
    - Required: Yes
    - Description: The GitHub username or organization that owns the repository.
    - Example: DamianFigiel
- GITHUB_REPO:
    - Required: Yes
    - Description: The name of the GitHub repository to clone.
    - Example: aws-ecs-terraform-python-api
- GITHUB_TOKEN:
    - Required: Yes
    - Description: The GitHub Personal Access Token (PAT) used for authenticating with the GitHub API to fetch commits and clone repositories. It must have the appropriate permissions (repo scope for private repositories, public_repo for public repositories).
    - Example: ghp_your_personal_access_token
- CLONE_DIR:
    - Required: No
    - Description: The directory where the repository will be cloned. If not provided, a default temporary directory will be used.
    - Example: /tmp/repo-clone
- DOCKERFILE_PATH:
    - Required: No
    - Description: The path to the Dockerfile inside the cloned repository. If not provided, the root directory of the cloned repo is assumed.
    - Example: /tmp/repo-clone/Dockerfile
- IMAGE_NAME:
    - Required: Yes
    - Description: The name to tag the Docker image with. 
    - Example: my-app-image:latest
- CLEANUP_REPO:
    - Required: No
    - Description: Boolean flag to determine if the cloned repository should be deleted after building the Docker image. Defaults to true.
    - Example: true
- K8S_NAMESPACE:
    - Required: Yes
    - Description: The Kubernetes namespace where the application will be deployed.
    - Example: default
- K8S_DEPLOYMENT_NAME:
    - Required: Yes
    - Description: The name of the Kubernetes deployment to update with the new Docker image.
    - Example: my-app-deployment

## Extending the Project
This project is designed to be modular and extendable. Here are a few ideas for how you can extend it:

- Add More Integrations: Integrate with additional CI/CD tools or services, such as Jenkins or GitLab CI.
- Environment-Specific Deployments: Extend the Kubernetes deployment functionality to handle multiple environments (e.g., development, staging, production).
- Automated Rollbacks: Implement automated rollback functionality in case a deployment fails.
- Slack Notifications: Integrate with Slack to send notifications on build or deployment success/failure.
- Health Checks: Add health checks for Kubernetes deployments to ensure that the new version is running smoothly.

Feel free to fork this project and extend it according to your needs!

## Contributing
Contributions are welcome! To contribute:

1. Fork the project.
2. Create a new branch for your feature or bugfix.
3. Commit your changes and submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for more information.