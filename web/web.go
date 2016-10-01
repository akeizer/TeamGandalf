package web

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>TeamGandalf</h1><div><a>https://github.com/AKeizer/TeamGandalf</a> Testomg</div>")
}

func Serve() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}
