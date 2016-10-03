package web

import (
    "../imagegen"
    "../imagetocsv"
    "../learning"
    "bytes"
    "encoding/base64"
    "fmt"
    "github.com/satori/go.uuid"
    "image/png"
    "log"
    "net/http"
    "os"
    "html/template"
    "path"
    "strings"
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
	Headers []string
    Data [][]string
	Accuracy string
}

func parseResultsTable(inTable string) ([]string, [][]string) {
    tableArr := strings.Split(inTable, "\n")
    tableHeaders := strings.Split(tableArr[0], "\t")
    var tableData [][]string
    for _, tableRow := range tableArr[2:len(tableArr)-2] {
        semiSplit := strings.Split(tableRow, "\t\t")
        var tableRowSplit []string
        for _, splitString := range semiSplit {
            tableRowSplit = append(tableRowSplit, strings.Split(splitString, "\t")...)
        }
        tableData = append(tableData, tableRowSplit)
    }
    return tableHeaders, tableData
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form)
	shape := r.Form["shape"]
	imageShape := shape[0]
  baseFileName := uuid.NewV4().String()
	if imageShape == "square" {
		baseFileName = "square-" + baseFileName
	}
	imageFile := baseFileName + ".png"
	imagegen.GenerateImage(imageShape, imageFile)
  // Convert to csv
  imagecsv := baseFileName + ".csv"
  err := imagetocsv.ConvertImageSet(imagecsv, imageFile)
  if err != nil {
      log.Fatalln("Could not convert image to csv")
  }

  results := learning.PerformAnalysis("training.csv", imagecsv);
	log.Println(results.Summary)
  //results := learning.AnalysisResult{"hey", 1.3}

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
        log.Println(err.Error())
		log.Println("unable to parse image template.")
	} else {
        headers, tableData := parseResultsTable(results.Summary)
		data := EndResult{str, headers, tableData, fmt.Sprintf("%.2f", results.Accuracy * 100)}

		log.Println("\n", results.Summary)
		if err = tmpl.ExecuteTemplate(w, "layout", data); err != nil {
			log.Println(err.Error())
			log.Println("unable to execute body template.")
		}
	}
}

func Serve() {
  http.HandleFunc("/", viewHandler)
  http.HandleFunc("/showresults/", resultHandler)
	log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}
