package controller

import (
	"log" // will be replaces with logrus

	"github.com/wilbeibi/kubeye/pkg/client"
)

func Run() {
	client, err := client.NewOutOfClusterClient()
	if err != nil {
		log.Fatal(err)
	}
	// TODO:informer
}
