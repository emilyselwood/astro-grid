package main

/*
TODO: output directory
TODO: log entries per bucket
TODO: build Folder structure for data.
TODO: html display of data.
TODO: Setup scripts
TODO: README.md
TODO: LICENCE.md
TODO: Document.
TODO: Blog post
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/wselwood/gompcreader"
)

const maxDistance = 6.0
const gridSize int = int(maxDistance * 10)

/*
GridEntry is a cell in the table. This will contain the cords and the values for the cell.
*/
type GridEntry struct {
	X       int
	Y       int
	StartX  float64
	EndX    float64
	StartY  float64
	EndY    float64
	Count   int32
	Special string
}

var inputfile = flag.String("in", "", "the minor planet center file to read")
var outputDir = flag.String("out", "", "the output path to write the structure")

func newGrid() *[gridSize][gridSize]GridEntry {
	var resultTable [gridSize][gridSize]GridEntry
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			resultTable[x][y].X = x
			resultTable[x][y].Y = y

			resultTable[x][y].StartX = float64(x) / 10.0
			resultTable[x][y].EndX = float64(x+1) / 10.0

			resultTable[x][y].StartY = float64(y) / 10.0
			resultTable[x][y].EndY = float64(y+1) / 10.0

			resultTable[x][y].Count = -1
			resultTable[x][y].Special = ""

		}
	}

	return &resultTable
}

func scaleAxis(in float64) int32 {
	if in <= maxDistance {
		return int32(in * 10.0)
	}
	return int32(in) + int32(gridSize)
}

func outputGrid(resultTable *[gridSize][gridSize]GridEntry) {
	f, err := os.Create(*outputDir + "/data.json")
	if err != nil {
		log.Fatal("Error opening datafile", err)
	}
	defer f.Close()
	f.WriteString("[")
	first := true
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {

			if resultTable[x][y].Count > -1 || resultTable[x][y].Special != "" {
				if first {
					first = false
				} else {
					f.WriteString(",\n")
				}

				entry := resultTable[x][y]

				js, e := json.Marshal(entry)
				if e == nil {
					f.WriteString(fmt.Sprintf("%s", js))
				} else {
					log.Fatal("error json marshal", e)
				}
			}
		}
	}
	f.WriteString("]")
}

func openOrCreateFile(path string) (*os.File, error) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		// create the file.
		f, err := os.Create(path)
		f.WriteString("id\n")
		return f, err
	} else if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
}

func pathIsDir(path string) (bool, error) {
	pathStat, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return pathStat.IsDir(), err
}

func addToFile(x int32, y int32, id string) error {
	path := fmt.Sprintf("%s/%d", *outputDir, x)
	exists, err := pathIsDir(path)
	if err != nil {
		return err
	} else if !exists {
		os.Mkdir(path, 0777)
	}
	// horribly inefficent should keep a cache of open files but its not a problem
	// given the size of the data we are working with.
	f, err := openOrCreateFile(fmt.Sprintf("%s/%d.txt", path, y))
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(id + "\n")

	return nil
}

func main() {

	flag.Parse()

	if *inputfile == "" {
		log.Fatal("No input file provided. Use the -in /path/to/file")
	}

	if *outputDir == "" {
		log.Fatal("No output path provided Use -out /output/path")
	}

	exists, err := pathIsDir(*outputDir)
	if err != nil {
		log.Fatal("Could not check output path existance")
	} else if !exists {
		log.Fatal("Output path does not exist or is not a directory")
	}

	mpcReader, err := gompcreader.NewMpcReader(*inputfile)
	if err != nil {
		log.Fatal("error creating mpcReader ", err)
	}
	defer mpcReader.Close()

	var count int64

	var resultTable = newGrid()

	result, err := mpcReader.ReadEntry()

	for err == nil {

		CS := result.SemimajorAxis * result.OrbitalEccentricity

		perihelion := result.SemimajorAxis - CS
		apohelion := result.SemimajorAxis + CS

		x := scaleAxis(perihelion)
		y := scaleAxis(apohelion)

		if x < int32(gridSize) && y < int32(gridSize) {
			if resultTable[x][y].Count == -1 {
				resultTable[x][y].Count = 0
			}
			resultTable[x][y].Count = resultTable[x][y].Count + 1

			e := addToFile(x, y, result.ID)
			if e != nil {
				log.Fatal(fmt.Sprintf("error updating for %d:%d with id: %s", x, y, result.ID), e)
			}
		}

		result, err = mpcReader.ReadEntry()
		count = count + 1
	}

	if err != nil && err != io.EOF {
		log.Fatal(fmt.Sprintf("error reading line %d\n", count), err)
	}

	// Now fill in the specal entries for the major planets.
	resultTable[10][10].Special = "Earth"
	resultTable[15][15].Special = "Mars"
	resultTable[52][52].Special = "Jupiter"

	outputGrid(resultTable)

}
