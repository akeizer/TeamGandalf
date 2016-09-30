package main

import (
    "os"
    "fmt"
    "flag"
    "./imagetocsv"
    "./learning"
)

func helpText() {
    fmt.Printf("Usage: %s [flags] out_file in_file(s)\n", os.Args[0])
}

func main() {
    if (len(os.Args) < 2) {
        helpText()
        os.Exit(1)
    }

    var train = flag.Bool("train", false, "Use the input files as training")
    var help = flag.Bool("h", false, "Display the help information")

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
        for _, arg := range args[1:] {
            outfile.WriteString(imagetocsv.ConvertToCSV(arg))
        }
        defer outfile.Close()
    }

    data := learning.RetrieveData(outfilename)
    fmt.Println(data)
}
