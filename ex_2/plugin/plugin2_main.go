package main

import (
	"net/http"

	"ex_2/host/runtime"
)

func helloPlugin2(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hello from plugin 2"))
}

func byePlugin2(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Bye bye from plugin 2"))
}

func Register(init runtime.Initializer) {
	init.RegisterEndpoint("plugin2/hello", helloPlugin2)
	init.RegisterEndpoint("plugin2/bye", byePlugin2)
}
