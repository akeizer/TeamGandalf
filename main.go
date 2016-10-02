package main

import (
    "os"
    "fmt"
    "flag"
    "./imagetocsv"
    "./learning"
    "./web"
)

func helpText() {
    fmt.Printf("Usage: %s [flags] out_file in_file(s)\n", os.Args[0])
}

func main() {
    totalpixels := 400

    var train = flag.Bool("train", false, "If we should use the inputs as training")
    var help = flag.Bool("h", false, "Should we display the help information")
    var isWeb = flag.Bool("web", false, "Setup webserver")
    flag.Parse()
    if len(os.Args) < 2 && !*isWeb {
        helpText()
        os.Exit(1)
    }

    if *help {
        helpText()
        os.Exit(1)
    }
    args := flag.Args()
    if *isWeb {
      web.Serve()
    } else {
      // open output file
      outfilename := args[0]
      if !*train {
          outfile, err := os.Create(outfilename)
          if err != nil {
              panic(fmt.Sprintf("Failed to create output file: %s ", err))
          }
          outfile.WriteString(imagetocsv.CreateHeaderRow(totalpixels) + "\n")
          for _, arg := range args[1:] {
              outfile.WriteString(imagetocsv.ConvertToCSV(arg) + "\n")
          }
          defer outfile.Close()
      }

      // want this code here so that it doesnt tell us that learning is unused
      // but also don't want it to fail while we change how the CSV is formatted
      if (false) {
          results := learning.PerformAnalysis("trainingFile", "testFile");
          fmt.Printf("Summary: %s", results.Summary)
          fmt.Printf("Accuracy: %e", results.Accuracy)
      }
    }
}
