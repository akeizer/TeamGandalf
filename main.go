package main

import (
    "flag"
    "fmt"
    "./imagetocsv"
    "./imagegen"
    "./learning"
    "./web"
    "os"
    "os/exec"
    "strings"
	  "github.com/satori/go.uuid"
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
    train := flag.Bool("train", true, "Use the input files as training")
    help := flag.Bool("h", false, "Print help text")
    isWeb := flag.Bool("web", false, "Setup webserver")
    shouldRun := flag.Bool("run", false, "Run command line")

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
      if !*train {
        outfilename := args[0]
        err := imagetocsv.ConvertImageSet(outfilename, args[1:]...)
        if err != nil {
            panic(fmt.Sprintf("Failed to create output file: %s ", err))
        }
      } else if *shouldRun {
        imageShape := args[0]
        baseFileName := uuid.NewV4().String()
      	imageFile := baseFileName + ".png"
      	imagegen.GenerateImage(imageShape, imageFile)
        // Convert to csv
        imagecsv := baseFileName + ".csv"
        err := imagetocsv.ConvertImageSet(imagecsv, imageFile)
        if err != nil {
            fmt.Printf("Could not convert image to csv")
            return
        }
        results := learning.PerformAnalysis("training.csv", imagecsv);
        fmt.Printf("Summary: %s", results.Summary)
        fmt.Printf("Accuracy: %e", results.Accuracy)
      }
    }
}
