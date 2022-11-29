package pkg

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	core "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	apimachinery "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreInformer "k8s.io/client-go/informers/core/v1"
	networkingInformers "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	coreListers "k8s.io/client-go/listers/core/v1"
	networkingListers "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	"time"
)

const (
	workNum = 5
)

type controller struct {
	client kubernetes.Interface
	//ingress的listener
	ingressLister networkingListers.IngressLister
	//service的listener
	serviceLister coreListers.ServiceLister
	queue         workqueue.RateLimitingInterface
}

func (c *controller) addService(obj interface{}) {

}

// 更新service
func (c *controller) updateService(obj interface{}, obj2 interface{}) {
	//	todo 比较annotation
	//如果两个对象完全相等，那么退出更新流程
	if reflect.DeepEqual(obj, obj2) {
		return
	}
	c.enqueue(obj)
}

// 加入workqueue
func (c *controller) enqueue(obj interface{}) {
	//从metadata里获取namespace和name
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		logrus.Error(err)
		runtime.HandleError(err)
	}
	c.queue.Add(key)
}

func (c *controller) deleteService(obj interface{}) {

}

func (c *controller) deleteIngress(obj interface{}) {
	ig := obj.(*networking.Ingress)
	ownerReference := apimachinery.GetControllerOf(ig)
	if ownerReference == nil {
		return
	}
	if ownerReference.Kind != "Service" {
		return
	}
	c.queue.Add(ig.Namespace + "/" + ig.Name)
}

func (c *controller) Run(ch chan struct{}) {
	for i := 0; i < workNum; i++ {
		go wait.Until(c.worker, time.Minute, ch)
	}
	<-ch

}

func (c *controller) worker() {
	for c.processNextItem() {
		logrus.Info("start work queue process")
	}
}

func (c *controller) processNextItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(item)
	key := item.(string)
	fmt.Println("processNextItem : ", key)

	err := c.syncService(key)
	if err != nil {
		c.queue.AddRateLimited(item)
		logrus.Error("syncService err: ", err)
	}
	return true
}

func (c *controller) syncService(item string) error {
	namespaceKey, name, err := cache.SplitMetaNamespaceKey(item)
	if err != nil {
		return err
	}
	service, err := c.serviceLister.Services(namespaceKey).Get(name)
	if !errors.IsNotFound(err) {
		logrus.Error("Get service error :", err)
	}
	if err != nil {
		return err
	}
	//	add and delete
	_, ok := service.GetAnnotations()["ingress/http"]
	igName, err := c.ingressLister.Ingresses(namespaceKey).Get(name)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	if ok && errors.IsNotFound(err) {
		//	create ingress
		igConfig := c.constructIngress(service)
		_, err := c.client.NetworkingV1().Ingresses(namespaceKey).Create(context.Background(), igConfig, apimachinery.CreateOptions{})
		if err != nil {
			return err
		}

	} else if !ok && igName != nil {
		//	delete ingress
		err := c.client.NetworkingV1().Ingresses(namespaceKey).Delete(context.Background(), name, apimachinery.DeleteOptions{})
		if err != nil {
			return err
		}

	}
	return nil
}

func (c *controller) constructIngress(svc *core.Service) *networking.Ingress {
	cn := "nginx"
	pathType := networking.PathTypePrefix
	ingress := &networking.Ingress{
		ObjectMeta: apimachinery.ObjectMeta{
			Name:              svc.Name,
			GenerateName:      "",
			Namespace:         svc.Namespace,
			CreationTimestamp: apimachinery.Time{},
			Labels:            nil,
			Annotations:       nil,
			OwnerReferences: []apimachinery.OwnerReference{
				*apimachinery.NewControllerRef(svc, apimachinery.SchemeGroupVersion.WithKind("Service")),
			},
		},
		Spec: networking.IngressSpec{
			IngressClassName: &cn,
			DefaultBackend:   nil,
			TLS:              nil,
			Rules: []networking.IngressRule{
				{
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathType,
									Backend: networking.IngressBackend{
										Service: &networking.IngressServiceBackend{
											Name: svc.Name,
											Port: networking.ServiceBackendPort{
												Name: svc.Name,
											},
										},
									},
								},
							},
						},
					},
					Host: "example.com",
				},
			},
		},
	}
	return ingress
}

func NewController(client kubernetes.Interface, serviceInformer coreInformer.ServiceInformer,
	ingressInformer networkingInformers.IngressInformer) controller {
	c := controller{
		client:        client,
		ingressLister: ingressInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "limitQueue"),
	}

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
		DeleteFunc: c.deleteService,
	})
	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})
	return c
}
