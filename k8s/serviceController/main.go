package main

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"serviceController/pkg"
)

func main() {
	//k8s config
	//client
	//informer
	//event handler
	//informer start
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedConfigPathFlag)
	if err != nil {
		panic(err)
	}
	//	create client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	//informer
	factory := informers.NewSharedInformerFactory(clientSet, 0)
	serviceInformer := factory.Core().V1().Services()
	ingressInformer := factory.Networking().V1().Ingresses()
	testCon := pkg.NewController(clientSet, serviceInformer, ingressInformer)

	stopCh := make(chan struct{})
	// 启动informer List & Watch
	factory.Start(stopCh)
	// 等待所有缓存同步
	factory.WaitForCacheSync(stopCh)

	testCon.Run(stopCh)
}
