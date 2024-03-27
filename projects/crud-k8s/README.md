# CRUD Operations With Kubernetes API

## Packages Used

- k8s.io/client-go
- k8s.io/apimachinery
- k8s.io/api

## CRUD Operations

### 1. Create
```golang
newPod, err := podsClient.Create(context.TODO(), podDefintion, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Pod '%s' is created!", newPod.Name)
```

### 2. Read
```golang
pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for i, pod := range pods.Items {
		fmt.Printf("Name of %dth pod: %s\n", i, pod.Name)
	}
```

### 3. Update
```golang
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
```

### 4. Delete
```golang
deleteErr := podsClient.Delete(context.TODO(), "demo-crud55wwk", metav1.DeleteOptions{})
	if deleteErr != nil {
		panic(deleteErr.Error())
	}
```
## Resource

- https://medium.com/cloud-native-daily/working-with-kubernetes-using-golang-a3069d51dfd6
- https://dev.to/docker/kubernetes-crud-operation-using-go-on-docker-desktop-3l6a
- https://github.com/iximiuz/client-go-examples/tree/main/crud-typed-simple