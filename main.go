package main

import (
    "os"
    "os/exec"
    "fmt"
    "flag"
)

func helpText() {
    fmt.Printf("Usage: %s [flags] out_file in_file(s)\n", os.Args[0])
}

func buildImageConvert(output string, inputs []string) (*Cmd, error) {
    path, err := exec.LookPath("imagecsv")
    if err != nil {
        return nil, err
    }
    return exec.Command(path, output, inputs)
}

func runImageConvert(imageConverter *Cmd) (error) {
    return imageConverter.Run()
}

func buildMLExec(training bool, inputFile *File) (*Cmd, error) {
    path, err := exec.LookPath("learning")
    if err != nil {
        return nil, err
    }
    return exec.Command(path, inputFile.Name())
}

func runML(machine *Cmd) (error) {
    return machine.Run()
}

func main() {
    // Set up flags
    train := flag.Bool("train", false, "Use the input files as training")
    raw := flag.Bool("data", false, "If the input needs to be converted")



    convertPath, err := exec.LookPath("imagetocsv")

}
