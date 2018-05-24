package main

import (
	"fmt"
	"net/url"
)

//Builder Interface
type ServiceBuilder interface {
	SetName(name string)
	SetAddress(addr string)
	Service() (*Service, error)
}

//main product
type Service struct {
	name string
	url  *url.URL
}

//concrete implementation
type authentication struct {
	name    string
	address string
}

func (a *authentication) SetName(name string) {
	a.name = name
}

func (a *authentication) SetAddress(addr string) {
	a.address = addr
}

func (a *authentication) Service() (*Service, error) {
	svc := &Service{name: a.name}
	url, err := url.Parse(a.address)
	if err == nil {
		svc.url = url
	}
	return svc, err
}

type nlu struct {
	name    string
	address string
}

func (a *nlu) SetName(name string) {
	a.name = name
}

func (a *nlu) SetAddress(addr string) {
	a.address = addr
}

func (a *nlu) Service() (*Service, error) {
	svc := &Service{name: a.name}
	url, err := url.Parse(a.address)
	if err == nil {
		svc.url = url
	}
	return svc, err
}

type Director struct {
	builder ServiceBuilder
}

func (d *Director) Build(name, address string) (*Service, error) {
	d.builder.SetAddress(address)
	d.builder.SetName(name)
	return d.builder.Service()
}

func main() {
	director := &Director{builder: new(authentication)}
	svc, _ := director.Build("auth", "https://auth.i.am")
	fmt.Println(svc.name, svc.url)
	director = &Director{builder: new(nlu)}
	svc, _ = director.Build("nlu", "https://nlu.com")
	fmt.Println(svc.name, svc.url)
}
