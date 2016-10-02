package learning

import (
    "fmt"
    "github.com/sjwhitworth/golearn/base"
    "github.com/sjwhitworth/golearn/evaluation"
    "github.com/sjwhitworth/golearn/knn"
)

type AnalysisResult struct {
  Summary string
  Accuracy float64
}

func RetrieveData(filename string) *base.DenseInstances {
  XORData, err := base.ParseCSVToInstances(filename, false)
  if err != nil  {
    panic(fmt.Sprintf("Couldn't load CSV file (error %s)", err))
  }
  return XORData
}


func ReadTrainingTestData(trainName string, testName string) (*base.DenseInstances, *base.DenseInstances) {
	classAttrs := make(map[int]base.Attribute)
	classAttrs[0] = base.NewCategoricalAttribute()
	classAttrs[0].SetName("label")
	// Setup the class Attribute to be in its own group
	classAttrGroups := make(map[string]string)
	classAttrGroups["label"] = "ClassGroup"
	// The rest can go in a default group
	attrGroups := make(map[string]string)

	inst1, err := base.ParseCSVToInstancesWithAttributeGroups(
		trainName,
		attrGroups,
		classAttrGroups,
		classAttrs,
		true,
	)
	if err != nil {
		panic(err)
	}
	inst2, err := base.ParseCSVToTemplatedInstances(
		testName,
		true,
		inst1,
	)
	if err != nil {
		panic(err)
	}
	return inst1, inst2
}

func TrainAndClassifyData(train *base.DenseInstances, test *base.DenseInstances) map[string]map[string]int {
    classifier := knn.NewKnnClassifier("euclidean", 1)
// note we can only optimize if the size of the data from test/ train have the
    // same dimensions
    classifier.AllowOptimisations = true
    classifier.Fit(train)
    predictions := classifier.Predict(test)
    c, err := evaluation.GetConfusionMatrix(test, predictions)
    if err != nil {
        panic(err)
    }
    return c
}

func PerformAnalysis(trainFile string, testFile string) AnalysisResult{
  train, test := ReadTrainingTestData(trainFile, testFile)
  c := TrainAndClassifyData(train, test)
  result := AnalysisResult{evaluation.GetSummary(c), evaluation.GetAccuracy(c)}
  return result;
}
