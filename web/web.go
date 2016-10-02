package web

import (
	"html/template"
	"net/http"
	"fmt"
	"log"
	"os"
  "path"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("web", "static", "templates", "layout.html")
  fp := path.Join("web", "static", "templates", r.URL.Path)
 		log.Println(lp)
		log.Println(fp)
	// check for template existence
	info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

	// ensure request is not for a directory
	if info.IsDir() {
    http.NotFound(w, r)
    return
  }
	tmpl, err := template.ParseFiles(lp, fp)
 	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
 	}
	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Recieved input")
}

func Serve() {
  http.HandleFunc("/", viewHandler)
  http.HandleFunc("/showresults/", resultHandler)
	log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}
