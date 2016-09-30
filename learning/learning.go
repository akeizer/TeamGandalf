package learning

import (
    "fmt"
    "github.com/sjwhitworth/golearn/base"
)

func RetrieveData(filename string) *base.DenseInstances {
  XORData, err := base.ParseCSVToInstances(filename, false)
  if err != nil  {
    panic(fmt.Sprintf("Couldn't load CSV file (error %s)", err))
  }
  return XORData
}
