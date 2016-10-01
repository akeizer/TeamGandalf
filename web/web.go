package web

import (
	"html/template"
	"net/http"
	"fmt"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/main.html")
	if err != nil {
		fmt.Print(err)
		return
	}
	t.Execute(w, nil)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Recieved input")
}

func Serve() {
  http.HandleFunc("/", viewHandler)
  http.HandleFunc("/showresults/", resultHandler)
  http.ListenAndServe(":8080", nil)
}
