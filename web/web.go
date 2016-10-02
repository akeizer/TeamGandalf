package web

import (
	"html/template"
	"net/http"
	"log"
	"os"
  "path"
	"github.com/AKeizer/TeamGandalf/imagegen"
	"github.com/AKeizer/TeamGandalf/learning"
	"github.com/AKeizer/TeamGandalf/imagetocsv"
	"github.com/satori/go.uuid"
	"image/png"
	"bytes"
	"encoding/base64"
	"strconv"
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

type EndResult struct {
	Image string
	Summary string
	Accuracy string
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form)
	shape := r.Form["shape"]
	imageShape := shape[0]
	imageFile := uuid.NewV4().String() + ".png"
	imagegen.GenerateImage(imageShape, imageFile)
	// results := learning.PerformAnalysis("training.csv", imageFile);
	results := learning.AnalysisResult{"hey", 1.3}

	lp := path.Join("web", "static", "templates", "layout.html")
  fp := path.Join("web", "static", "templates", "results.html")

	img, err := imagetocsv.ReadImage(imageFile)
    if err != nil {
        log.Fatalln("unable to read image.")
    }
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Fatalln("unable to encode image.")
	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.ParseFiles(lp, fp); err != nil {
		log.Println("unable to parse image template.")
	} else {
		// data := map[string]interface{}{"Image": str, "Summary" results.Summary, "Accuracy", results.Accuracy}
		data := EndResult{str, results.Summary, strconv.FormatFloat(results.Accuracy, 'e', -1, 64)}
		log.Println(data.Summary)
		log.Println(data.Image)
		if err = tmpl.ExecuteTemplate(w, "layout", data); err != nil {
			log.Println(err.Error())
			log.Println("unable to execute template.")
		}
	}
}

func Serve() {
  http.HandleFunc("/", viewHandler)
  http.HandleFunc("/showresults/", resultHandler)
	log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}
