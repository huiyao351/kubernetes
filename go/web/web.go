package web

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/go-web"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	"github.com/micro/go-plugins/selector/static"
)

// NewService returns a web service for kubernetes
func NewService(opts ...web.Option) web.Service {
	// setup
	k := kubernetes.NewRegistry()
	st := static.NewSelector()

	// create new service
	service := grpc.NewService(
		micro.Registry(k),
		micro.Selector(st),
	)

	// prepend option
	options := []web.Option{
		web.MicroService(service),
	}

	options = append(options, opts...)

	// return new service
	return web.NewService(options...)
}
