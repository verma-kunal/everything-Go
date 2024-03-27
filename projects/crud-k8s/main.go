package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
)

func main() {

	// get kubeconfig
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	// create a new client
	client := kubernetes.NewForConfigOrDie(config)

	// define the namespace
	namespace := "default"

	// define the pods client (easy for later use)
	podsClient := client.CoreV1().Pods(namespace)

	// read all pods
	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for i, pod := range pods.Items {
		fmt.Printf("Name of %dth pod: %s\n", i, pod.Name)
	}

	// create a pod defintion
	podDefintion := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "go-api-",
			Namespace:    "default",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "nginx-container",
					Image: "nginx:latest",
				},
			},
		},
	}
	// create a new pod
	newPod, err := podsClient.Create(context.TODO(), podDefintion, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Pod '%s' is created!", newPod.Name)

	// update a pod
	fmt.Println("Updating pod...")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

		// retrive the latest pod
		currentPod, updateErr := podsClient.Get(context.TODO(), "go-api-2mwpl", metav1.GetOptions{})
		if updateErr != nil {
			panic(updateErr.Error())
		}

		// change container image
		currentPod.Spec.Containers[0].Image = "nginx:1.25.4"

		// update pod
		updatedPod, updateErr := podsClient.Update(context.TODO(), currentPod, metav1.UpdateOptions{})
		fmt.Printf("Updated pod: %s", updatedPod.Name)
		return updateErr
	})
	if retryErr != nil {
		panic(retryErr.Error())
	}

	// delete a pod
	deleteErr := podsClient.Delete(context.TODO(), "demo-crud55wwk", metav1.DeleteOptions{})
	if deleteErr != nil {
		panic(deleteErr.Error())
	}

}
