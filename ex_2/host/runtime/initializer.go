package runtime

import (
	"net/http"
)

type Initializer interface {
	RegisterEndpoint(path string, handler func(http.ResponseWriter, *http.Request))
}
