package main

import (
	"fmt"
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"

	
	

)

func main(){

	kubeconfig := flag.String("kubeconfig","/home/deepak/.kube/config","location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err)
	}
	

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	

	pods := clientset.CoreV1().pods("default").List(Context.Background(),metav1.ListOptions{})
     
	fmt.Println(pods)





}