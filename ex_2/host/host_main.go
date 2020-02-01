package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"ex_2/host/runtime"
)

var pluginDir = flag.String("plugins", "plugins", "The directory containing all them plugins")

type localInitializer struct {
	endpoints []string
}

func (i *localInitializer) RegisterEndpoint(path string, handler func(http.ResponseWriter, *http.Request)) {

	endpoint := "/api/" + path

	http.HandleFunc(endpoint, handler)

	log.Printf("plugin registered endpoint `%s`", endpoint)

	i.endpoints = append(i.endpoints, endpoint)
}

func main() {

	plugins := []string{}

	if err := filepath.Walk(*pluginDir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() || filepath.Ext(path) != ".so" {
			return nil
		}

		plugins = append(plugins, path)

		return nil

	}); err != nil {
		panic(err)
	}

	log.Printf("located plugin files: [%s]", strings.Join(plugins, ", "))

	init := &localInitializer{}

	for _, path := range plugins {

		p, err := plugin.Open(path)
		if err != nil {
			panic(err)
		}

		symbol, err := p.Lookup("Register")
		if err != nil {
			panic(err)
		}

		if regFunc, is := symbol.(func(runtime.Initializer)); is {

			regFunc(init)

		} else {
			panic("incorrect signature for plugin `Register` function")
		}
	}

	http.HandleFunc("/", func(resp http.ResponseWriter, _ *http.Request) {
		resp.Write([]byte(strings.Join(init.endpoints, "\n")))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
