package main

import (
    "os"
    "fmt"
    "flag"
    "github.com/joshkergan/TeamGandalf/imagetocsv"
)

func helpText() {
    fmt.Printf("Usage: %s [flags] out_file in_file(s)\n", os.Args[0])
}

func main() {
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
    var outfile os.File
    if !*train {
        outfile, err := os.Create(outfilename)
        if err != nil {
            fmt.Println("Failed to create output file: ", err)
        }
        defer outfile.Close()
    }

    // write csvs
    for _, arg := range args[1:] {
        outfile.WriteString(imagetocsv.ConvertToCSV(arg))
    }
}
