package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"greeter/handler"
	"greeter/subscriber"

	example "greeter/proto/example"
)

func main() {
	// New Service
	service := k8s.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.greeter", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.greeter", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
