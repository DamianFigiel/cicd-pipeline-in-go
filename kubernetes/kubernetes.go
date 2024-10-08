package kubernetes

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeClient() (*kubernetes.Clientset, error) {
	kubeconfig := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("Error to load kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Kubernetes client: %v", err)
	}

	return clientset, nil
}

func UpdateDeployment(clientset *kubernetes.Clientset, namespace, deploymentName, newImage string) error {
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Failed to get deployment: %v", err)
	}

	deployment.Spec.Template.Spec.Containers[0].Image = newImage

	_, err = clientset.AppsV1().Deployments(namespace).Update(context.Background(), deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("Failed to update deployment: %v", err)
	}

	fmt.Printf("Deployment %s in namespace %s updated to image %s\n", deploymentName, namespace, newImage)
	return nil
}
