package main

import (
	"plugin"
)

func main() {

	p, err := plugin.Open("plugin_main.so")
	if err != nil {
		panic(err)
	}

	s, err := p.Lookup("PrintHello")
	if err != nil {
		panic(err)
	}

	counterSymbol, err := p.Lookup("callCounter")
	if err != nil {
		panic(err)
	}

	*counterSymbol.(*int) = 13

	s.(func())()
	s.(func())()
	s.(func())()
}
