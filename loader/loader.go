package main

import (
	"plugin"

	"github.com/qaisjp/go-plugins-experiment/structs"
)

type Context = structs.Context

func main() {
	pl, err := plugin.Open("plugin1.so")
	if err != nil {
		panic(err)
	}

	c := &Context{
		IP:   "192.168.0.1",
		Port: 8080,
	}

	cSym, err := pl.Lookup("Context")
	if err != nil {
		panic(err)
	}

	theirC, ok := cSym.(**Context)
	if !ok {
		panic("Invalid type of Context")
	}

	*theirC = c

	mainSym, err := pl.Lookup("OnLoad")
	if err != nil {
		panic(err)
	}

	theirMain, ok := mainSym.(func())
	if !ok {
		panic("Invalid type of OnLoad")
	}

	theirMain()
}
