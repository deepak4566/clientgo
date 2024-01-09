package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	//solve terminal error

	kubeconfig := flag.String("kubeconfig", "/c/Users/91911/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		// Handle error here.
		// TODO: Add error handling code.
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		// Handle error here.
		// TODO: Add error handling code.
	}
	fmt.Println(clientset)

	ctx := context.Background()

	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})

	if err != nil {
		//null
	}

	fmt.Println(pods)

}
