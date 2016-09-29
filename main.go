package main

import (
    "os"
    "fmt"
    "flag"
)

func main() {
    if (len(os.Args) < 2) {
        fmt.Printf("Usage: imagetocsv out_file in_file(s)")
        return
    }

    args := os.Args[1:]
    outfilename := args[0]

    // open output file
    outfile, err := os.Create(outfilename)
    if err != nil {
        panic(err)
    }
    defer outfile.Close()

    // write csvs
    for _, arg := range args[1:] {
        outfile.WriteString(imageToCSV(arg))
    }
}
