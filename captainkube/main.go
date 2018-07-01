package main

import (
	"time"
	"bytes"
	"encoding/json"
	"net/http"
	"log"
	"reflect"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type Pod struct {
	Container      string
	ContainerImage string
	Name           string
	Namespace      string
	Status         string
	Action         string
}

func main() {
	log.Println("Starting up Captain Kube")

	// stop will be used by the informer to allow a clean shutdown
	// If the channel is closed, it communicates the informer that it needs to shutdown
	stop := make(chan struct{})

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Clear the cluster status, start with a blank slate
	/*
	req, err := http.NewRequest(http.MethodDelete, "http://parrot-parrot/api/ClusterStatus", bytes.NewBuffer([]byte(``)))
	httpclient := &http.Client{}
    _, err = httpclient.Do(req)
	if err != nil {
		log.Printf("The HTTP request failed with error %s", err)
	} else {
		log.Printf("\n\n**** Cleared parrot****\n\n")
	}
	*/

	watchList := cache.NewListWatchFromClient(client.Core().RESTClient(), "pods", v1.NamespaceAll, fields.Everything())

	// Setup the informer that will start watching for pod triggers
	informer := cache.NewSharedIndexInformer(
		watchList,
		&v1.Pod{},
		10*time.Second,
		cache.Indexers{},
	) // We only want `Pod`, force resync every 10 seconds
	
	  // Setup the trigger handlers that will receive triggers
	  informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		// This method is executed when a new pod is created
		AddFunc: func(obj interface{}) {
			pod, ok := obj.(*v1.Pod) // cast the object as a pod
			if !ok {
				//log.Printf("Couldn't cast object as pod: %s", obj)
				return
			}
			pingparrot(pod,"Added") // Ping the parrot
		},
		// This method is executed when an existing pod is updated
		UpdateFunc: func(oldObj, newObj interface{}) {
			newPod, ok := newObj.(*v1.Pod) // cast the object as a pod
			if !ok {
				//log.Printf("Couldn't cast object as pod: %s", newObj)
				return
			}
			// Deep compare objects and only notify if they are truly different
			if !reflect.DeepEqual(oldObj, newObj) {
				pingparrot(newPod,"Updated") // Ping the parrot
			}
		},
		// This method is executed when an existing pod is deleted
		DeleteFunc: func(obj interface{}) {
			pod, ok := obj.(*v1.Pod) // cast the object as a pod
			if !ok {
				//log.Printf("Couldn't cast object as pod: %s", obj)
				return
			}
			pingparrot(pod,"Deleted") // Ping the parrot
		},
	  })
	
	  // Start the informer, until `stop` is closed
	  informer.Run(stop)
}

func pingparrot(pod *v1.Pod, state string) {
	if pod.ObjectMeta.Namespace != "kube-system" {
		log.Printf("Pod %s: %s", state, pod.ObjectMeta.Name)
		log.Printf("namespace: %s", pod.ObjectMeta.Namespace)
		log.Printf("status: %s", pod.Status.Phase)
		log.Printf("startTime: %s", pod.Status.StartTime)
		log.Printf("conditions:")

		for _, condition := range pod.Status.Conditions {
			log.Printf("\ttype: %s", condition.Type)
			log.Printf("\tlastTransitionTime: %s", condition.LastTransitionTime)
		}

		// shrink the object we send over
		p := Pod{Action: state, Container: pod.Spec.Containers[0].Name, ContainerImage: pod.Spec.Containers[0].Image, Name: pod.ObjectMeta.Name, Namespace: pod.ObjectMeta.Namespace, Status: string(pod.Status.Phase)}

		jsonValue, _ := json.Marshal(p)
		//log.Printf("\n%s\n",jsonValue)

		_, err := http.Post("http://parrot-parrot/api/ClusterStatus", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			log.Printf("The HTTP request failed with error %s", err)
		} else {
			log.Printf("Notified parrot: %s", state)
		}
		log.Printf("\n\n")
	}
}