package main

import (
    "flag"
    "fmt"
    "github.com/joshkergan/TeamGandalf/imagetocsv"
    "github.com/joshkergan/TeamGandalf/learning"
    "github.com/sjwhitworth/golearn/evaluation"
    "os"
    "os/exec"
    "strings"
)

func helpText() {
    fmt.Printf("Usage: %s [flags] out_file in_file(s)\n", os.Args[0])
}

func buildImageConvert(output string, inputs []string) (*exec.Cmd, error) {
    path, err := exec.LookPath("imagecsv")
    if err != nil {
        return nil, err
    }
    return exec.Command(path, output, strings.Join(inputs, " ")), nil
}

func runImageConvert(imageConverter *exec.Cmd) error {
    return imageConverter.Run()
}

func buildMLExec(training bool, inputFile *os.File) (*exec.Cmd, error) {
    path, err := exec.LookPath("learning")
    if err != nil {
        return nil, err
    }
    return exec.Command(path, inputFile.Name()), nil
}

func runML(machine *exec.Cmd) error {
    return machine.Run()
}

func main() {
    // Set up flags
    train := flag.Bool("train", false, "Use the input files as training")
    raw := flag.Bool("data", false, "If the input needs to be converted")
    help := flag.Bool("h", false, "Print help text")

    totalpixels := 400

    if len(os.Args) < 2 {
        *help = true
    }

    if *help {
        helpText()
        os.Exit(1)
    }
    args := flag.Args()
    outfilename := args[0]

    // open output file
    if !*raw {
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
    if !*train {
        train, test := learning.ReadTrainingTestData("trainingFile", "testFile")

        c := learning.TrainAndClassifyData(train, test)

        fmt.Println(evaluation.GetSummary(c))
        fmt.Println(evaluation.GetAccuracy(c))
    }
}
