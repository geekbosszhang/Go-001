//+build wireinject
package main

import (
	"github.com/google/wire"
	"service"
)

func InitializeEvent() service.Event  {
	wire.Build(service.NewEvent, service.NewGreeter, service.NewMessage)
    return service.Event{}
}