package main

import (
    "flag"
    "fmt"
    "./imagetocsv"
    "./learning"
    "./web"
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
    help := flag.Bool("h", false, "Print help text")
    isWeb := flag.Bool("web", false, "Setup webserver")

    if len(os.Args) < 2 && !*isWeb {
        helpText()
        os.Exit(1)
    }

    flag.Parse()
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
          err := imagetocsv.ConvertImageSet(outfilename, args[1:]...)
        if err != nil {
            panic(fmt.Sprintf("Failed to create output file: %s ", err))
        }
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
