---
marp: true
theme: uncover
class: invert
---

# Go Plugins in a Nutshell

A quick tour

by @nitzanzada

![bg left:40% 100%](assets/gopher.png)

---

# What for?

Load compiled go binaries and symbols at runtime

---

<style scoped>
pre { margin: auto 0 }
</style>

# When, where?

- Been around for a while, since **1.8**
- `import "plugin"`
- `$ go build -buildmode=plugin`

```go
func Open(path string) (*Plugin, error)

func (p *Plugin) Lookup(symName string) (Symbol, error)
```

Literally the only two exported members of the `"plugin"` package

---

# How?

<!--
    Example 1
    Showing simple building, loading and running of a function 
-->

---

# Why Though?

- Separation of concerns
- Smaller docker image layer changes
- Required modular behavior

---

# Gotchas

- Plugin code must be compiled against the same go version as the host code
- That includes the standard library and shared dependencies
- No way to unload loaded binaries

---

# Good Practices

- Registration method
- Plugin builder images reflecting plugin host versions

---

<style scoped>
pre { margin: auto 0; font-size: 0.6em }
section { text-align: start; }
</style>

## Good Practices
Registration method

```go
// host.go

func main() error {

}
```

```go
// plugin.go

func Register(initializer host.Initializer, logger *log.Logger) error {
    // do things with initializer
}
```

<!--
    Example 2
    Show an example of a folder being scanned and register fun best practice
-->

---

<style scoped>
pre { margin: auto 0; font-size: 0.6em }
section { text-align: start; }
</style>

## Good Practices
Plugin builder images reflecting plugin host versions

```dockerfile
# plugin builder Dockerfile

FROM golang:1.13.7-buster as builder

ENTRYPOINT ["go", "build", "-buildmode=plugin"]
```

```dockerfile
# host Dockerfile

FROM golang:1.13.7-buster

RUN mkdir /app/plugins

ENTRYPOINT ["app", "--plugins", "/app/plugins"]
```

<!--
    Example 2
    Show an example of a folder being scanned and register fun best practice
-->
