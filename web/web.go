package web

import (
	"html/template"
	"net/http"
	"fmt"
)

type Data struct {
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/main.html")
	// d := Data{}
	if err != nil {
		fmt.Print(err)
		return
	}
	t.Execute(w, nil)
}

func Serve() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}
