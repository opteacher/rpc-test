package service

import (
	"fmt"
	"rpcserver/domain"
)

type helloSvc struct {}

func (bs *baseService) NewHelloSvc() *helloSvc {
	return &helloSvc{}
}

func (hs *helloSvc) SayHello(name string) domain.Resp {
	return domain.Resp{
		Status: 200,
		Message: fmt.Sprintf("Hello, %s", name),
		Data: make(map[string]interface{}),
	}
}