package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	klog "k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	conf, err := config.GetConfig()

	if err != nil {
		logrus.Panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(conf)
	informerFactory := informers.NewSharedInformerFactory(clientSet, 30*time.Second)
	deploy := informerFactory.Apps().V1().Deployments()
	deployInformer := deploy.Informer()
	deployListener := deploy.Lister()
	// 资源事件
	deployInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})
	stopper := make(chan struct{})
	defer close(stopper)
	// 启动informer List & Watch
	informerFactory.Start(stopper)
	// 等待所有缓存同步
	informerFactory.WaitForCacheSync(stopper)
	deployments, err := deployListener.Deployments("default").List(labels.Everything())
	for index, deploy := range deployments {
		fmt.Printf("%d -> %s", index, deploy.Name)

	}
	<-stopper
}
func onAdd(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	klog.Infoln("add a deploy: ", deploy.Name)
}

func onUpdate(old, new interface{}) {
	oldDeploy := old.(*v1.Deployment)
	newDeploy := new.(*v1.Deployment)
	klog.Infoln("update deploy: ", oldDeploy.Status.Replicas, newDeploy.Status.Replicas)
}

func onDelete(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	klog.Infoln("delete a deploy:", deploy.Name)
}
