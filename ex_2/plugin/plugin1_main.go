package main

import (
	"net/http"

	"ex_2/host/runtime"
)

func helloPlugin1(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hello from plugin 1"))
}

func byePlugin1(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Bye bye from plugin 1"))
}

func Register(init runtime.Initializer) {
	init.RegisterEndpoint("plugin1/hello", helloPlugin1)
	init.RegisterEndpoint("plugin1/bye", byePlugin1)
}
