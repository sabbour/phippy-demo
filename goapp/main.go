package main

import (
	"fmt"
	"time"
	
	"k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// start watching
	watchlist := cache.NewListWatchFromClient(
        clientset.CoreV1().RESTClient(),
        string(v1.ResourcePods), // Kubernetes object to watch
        v1.NamespaceAll, // namespace to watch
        fields.Everything(), // fields
    )
    _, controller := cache.NewInformer(
        watchlist,
        &v1.Pod{}, // Kubernetes object to watch
        0, // resync period if non-zero, will re-list this often (you will get OnUpdate
			// calls, even if nothing changed). Otherwise, re-list will be delayed as
			// long as possible (until the upstream source closes the watch or times out,
			// or you stop the controller).
        cache.ResourceEventHandlerFuncs{
			// called when an object is added
            AddFunc: func(obj interface{}) {
				// cast the object as a pod
				pod, ok := obj.(*v1.Pod)
				if !ok {
					fmt.Printf("couldn't cast object as pod: %s \n", obj)
					return
				}
                fmt.Printf("pod added: %s \n", pod.ObjectMeta.Name)
                fmt.Printf("\tnamespace: %s \n", pod.ObjectMeta.Namespace)
                fmt.Printf("\tlabels: %s \n", pod.ObjectMeta.Labels)
                fmt.Printf("\tstatus: %s \n", pod.Status.Phase)
			},
			// called when an object is modified. Note that oldObj is the
    		// last known state of the object-- it is possible that several changes
        	// were combined together, so you can't use this to see every single
    		// change. OnUpdate is also called when a re-list happens, and it will
    		// get called even if nothing changed. This is useful for periodically
    		// evaluating or syncing something.
            UpdateFunc: func(oldObj, newObj interface{}) {
				// cast the object as a pod
				//oldPod, ok := oldObj.(*v1.Pod)
				//if !ok {
				//	fmt.Printf("couldn't cast object as pod: %s \n", oldObj)
				//	return
				//}
				// cast the object as a pod
				newPod, ok := newObj.(*v1.Pod)
				if !ok {
					fmt.Printf("couldn't cast object as pod: %s \n", newObj)
					return
				}
				fmt.Printf("pod changed: %s \n", newPod.ObjectMeta.Name)
                fmt.Printf("\tnamespace: %s \n", newPod.ObjectMeta.Namespace)
                fmt.Printf("\tlabels: %s \n", newPod.ObjectMeta.Labels)
                fmt.Printf("\tstatus: %s \n", newPod.Status.Phase)
			},
			// will get the final state of the item if it is known, otherwise
			// it will get an object of type DeletedFinalStateUnknown. This can
			// happen if the watch is closed and misses the delete event and we don't
			// notice the deletion until the subsequent re-list.
            DeleteFunc: func(obj interface{}) {
				// cast the object as a pod
				pod, ok := obj.(*v1.Pod)
				if !ok {
					fmt.Printf("couldn't cast object as pod: %s \n", obj)
					return
				}
				fmt.Printf("pod deleted: %s \n", pod.ObjectMeta.Name)
                fmt.Printf("\tnamespace: %s \n", pod.ObjectMeta.Namespace)
                fmt.Printf("\tlabels: %s \n", pod.ObjectMeta.Labels)
                fmt.Printf("\tstatus: %s \n", pod.Status.Phase)
            },
         },
	 )
	 
	stop := make(chan struct{})
    defer close(stop)
    go controller.Run(stop)
    for {
        time.Sleep(time.Second*10)
    }
}