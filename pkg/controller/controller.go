package controller

import (
	log "github.com/sirupsen/logrus"
	"github.com/wilbeibi/kubeye/pkg/client"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type Controller struct {
	logger    *log.Entry
	clientset kubernetes.Interface
	// queue	workqueue.RateLimitingInterface
	informer cache.SharedIndexInformer
	// handler   handlers.Handler
}

func NewController() *Controller {
	client, err := client.NewOutOfClusterClient()
	if err != nil {
		log.Fatal(err)
	}
	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return client.CoreV1().Pods("").List(options) // "" == metav1.NamespaceAll
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return client.CoreV1().Pods("").Watch(options)
			},
		},
		&v1.Pod{},
		0, //Skip resync
		cache.Indexers{},
	)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// convert obj into key as "namespace/name"
			key, err := cache.MetaNamespaceKeyFunc(obj)
			log.Infof("Add %s", key)
			if err != nil {
				log.Error(err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(old)
			log.Infof("Update %s", key)
			if err != nil {
				log.Error(err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			log.Infof("Delete %s", key)
			if err != nil {
				log.Error(err)
			}
		},
	})
	return &Controller{
		logger:    log.WithField("pkg", "kubeye"),
		clientset: client,
		informer:  informer,
	}
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	c.logger.Info("Start running controller")
	go c.informer.Run(stopCh)
	<-stopCh
}
