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

# What?

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

---

# How?

<!--
    Example 1
    Showing simple building, loading and running of a function 
-->

---

<style scoped>
section, ul, ol { text-align: start; margin: 0 auto auto 0; }
pre { width: 100%; }
</style>

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

