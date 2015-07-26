package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
Dimension defines an axis on the result
*/
type Dimension struct {
	Name        string         `json:"n"`
	MinValue    int            `json:"min"`
	MaxValue    int            `json:"max"`
	GridSize    int            `json:"grid"`
	Description string         `json:"desc"`
	Extractor   ValueExtractor `json:"-"`
}

/*
NumDimentions contains the number of dimensions we are processing.
*/
const NumDimentions = 3

/*
BuildDimensions will create the standard set of dimensions
*/
func BuildDimensions() [NumDimentions]Dimension {
	var result [NumDimentions]Dimension

	result[0] = buildApohelion()
	result[1] = buildPerihelion()
	result[2] = buildYearOfFirstObs()
	return result
}

func buildApohelion() Dimension {
	var result Dimension

	result.Name = "Apohelion"
	result.MinValue = 0
	result.MaxValue = 6
	result.GridSize = 60

	result.Extractor = &ApohelionExtractor{6, 10.0}

	return result
}

func buildPerihelion() Dimension {
	var result Dimension

	result.Name = "Perihelion"
	result.MinValue = 0
	result.MaxValue = 6
	result.GridSize = 60

	result.Extractor = &PerihelionExtractor{6, 10.0}

	return result
}

func buildYearOfFirstObs() Dimension {
	var result Dimension

	result.Name = "Year-Of-First-Obs"
	result.MinValue = 1915
	result.MaxValue = 2015
	result.GridSize = 100
	result.Extractor = &YearOfFirstObsExtractor{1915}

	return result
}

/*
RenderDimensions output the dimension listing to the outputDir given
*/
func RenderDimensions(outputDir string, dimensions []Dimension) {

	out := fmt.Sprintf("%s/dimensions.json", outputDir)
	f, err := os.Create(out)
	if err != nil {
		log.Fatal("Error opening datafile", err)
	}
	defer f.Close()

	js, e := json.Marshal(dimensions)
	if e == nil {
		f.WriteString(fmt.Sprintf("%s\n", js))
	}
}
