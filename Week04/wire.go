//+build wireinject
package main

import (
	"github.com/google/wire"
	"github.com/geekbosszhang/Go-001/Week04/internal/service"
)

func InitializeEvent() service.Event  {
	wire.Build(service.NewEvent, service.NewGreeter, service.NewMessage)
    return service.Event{}
}