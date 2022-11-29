package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//get k8s  config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedConfigPathFlag)
	if err != nil {
		panic(err)
	}
	//config.GroupVersion = &schema.GroupVersion{}
	//config.NegotiatedSerializer = scheme.Codecs
	//config.APIPath = "/apis"
	////fmt.Println(config)
	////	client
	//restClient, err := rest.RESTClientFor(config)
	//if err != nil {
	//	panic(err)
	//}
	//pod := v1.Pod{}
	//err = restClient.Get().Namespace("default").Resource("pod").Name("coffee-6ddc6f7584-f8vgb").Do(context.TODO()).Into(&pod)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(pod.Name)
	//fmt.Println(restClient.APIVersion().String())

	clientset, err := kubernetes.NewForConfig(config)
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "busybox", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pod.Name)
}
