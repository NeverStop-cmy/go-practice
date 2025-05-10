package stdlib

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, 你访问的路径是: %s", r.URL.Path[1:])
}

func RunHttpServerDemo() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
