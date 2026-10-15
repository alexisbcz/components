# components

The shadcn/ui components for Go. Each component renders HTML on the server, and plain JavaScript drives its client behavior.

## Install

```sh
go get github.com/alexisbcz/components
go mod tidy
```

## Use

```go
package main

import (
	"net/http"

	"github.com/alexisbcz/components/ui/button"
	"github.com/alexisbcz/components/ui/runtime"
	h "github.com/alexisbcz/libhtml"
)

func main() {
	http.Handle(runtime.Prefix, runtime.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := h.Html(
			h.Head(runtime.Scripts(r)),
			h.Body(button.Button(button.VariantOutline, "Click me")),
		)
		h.WriteHTTP(w, http.StatusOK, page)
	})
	http.ListenAndServe(":8080", nil)
}
```

## License

MIT. See [LICENSE](LICENSE) and [NOTICE](NOTICE).
