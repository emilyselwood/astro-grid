package main

/*
TODO: html display of data.
TODO: Setup scripts
TODO: README.md
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

var inputfile = flag.String("in", "", "the minor planet center file to read")
var outputDir = flag.String("out", "", "the output path to write the structure")

func outputGrid(dimentions []Dimension, resultTable [][]Grid) {
	for i := 0; i < NumDimentions; i++ {
		for j := 0; j < NumDimentions; j++ {
			path := fmt.Sprintf("%s/%s/%s/", *outputDir, dimentions[i].Name, dimentions[j].Name)
			os.MkdirAll(path, 0777)

			f, err := os.Create(fmt.Sprintf("%s/data.json", path))
			if err != nil {
				log.Fatal("Error opening datafile", err)
			}
			defer f.Close()
			f.WriteString("[")
			first := true

			var table = resultTable[i][j].G

			for x := 0; x < resultTable[i][j].SizeX; x++ {
				for y := 0; y < resultTable[i][j].SizeY; y++ {
					entry := table[x][y]
					if entry.Count > 0 || entry.Special != "" {
						if first {
							first = false
						} else {
							f.WriteString(",\n")
						}

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
	}
}

func createPathIfNeeded(path string) {
	os.MkdirAll(path, 0777)
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

	var dimentions = BuildDimensions()
	var resultTable = BuildResultsGrid(dimentions[:])

	var count int64

	//var resultTable = newGrid()

	result, err := mpcReader.ReadEntry()

	for err == nil {

		for i := 0; i < NumDimentions; i++ {
			for j := 0; j < NumDimentions; j++ {

				x := dimentions[i].Extractor.ExtractCell(result)
				y := dimentions[j].Extractor.ExtractCell(result)
				if x > 0 && y > 0 {
					grid := resultTable[i][j].G

					grid[x][y].Count = grid[x][y].Count + 1

					if grid[x][y].X == 0 {
						grid[x][y].X = int(x)
						grid[x][y].Y = int(y)
						grid[x][y].StartX = dimentions[i].Extractor.Extract(result)
						grid[x][y].StartY = dimentions[j].Extractor.Extract(result)
					}
				}
			}
		}

		result, err = mpcReader.ReadEntry()
		count = count + 1
	}

	if err != nil && err != io.EOF {
		log.Fatal(fmt.Sprintf("error reading line %d\n", count), err)
	}

	// Now fill in the specal entries for the major planets.
	//resultTable[10][10].Special = "Earth"
	//resultTable[15][15].Special = "Mars"
	//resultTable[52][52].Special = "Jupiter"

	outputGrid(dimentions[:], resultTable)
	RenderDimensions(*outputDir, dimentions[:])
}
