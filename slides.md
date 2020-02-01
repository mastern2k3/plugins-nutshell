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

# Hazards

- Plugin code must be compiled against the same go version as the host code
- That includes the standard library and shared dependencies
- No way to unload loaded binaries

---


<!--
    Example 2
    Show an example of a folder being scanned and register fun best practice
-->


# Benefits

- Separation of concerns
- Smaller docker image layer changes

---

# Good Practices

-  initialization method

