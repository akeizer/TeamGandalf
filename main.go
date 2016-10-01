package main

import (
    "os"
    "fmt"
    "flag"
    "./imagetocsv"
    "./learning"
    "github.com/sjwhitworth/golearn/evaluation"
)

func helpText() {
    fmt.Printf("Usage: %s [flags] out_file in_file(s)\n", os.Args[0])
}

func main() {
    totalpixels := 400

    if (len(os.Args) < 2) {
        helpText()
        os.Exit(1)
    }

    var train = flag.Bool("train", false, "If we should use the inputs as training")
    var help = flag.Bool("h", false, "Should we display the help information")

    flag.Parse()

    if *help {
        helpText()
        os.Exit(1)
    }
    args := flag.Args()
    outfilename := args[0]

    // open output file
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
        train, test := learning.ReadTrainingTestData("trainingFile", "testFile")

        c := learning.TrainAndClassifyData(train, test)

        fmt.Println(evaluation.GetSummary(c))
        fmt.Println(evaluation.GetAccuracy(c))
    }

}
