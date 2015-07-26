package main

/*
GridEntry is a cell in the table. This will contain the cords and the values for the cell.
*/
type GridEntry struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	StartX  string `json:"sx"`
	StartY  string `json:"sy"`
	Count   int32  `json:"c"`
	Special string `json:"s,omitempty"`
}

/*
Grid holds a result table
*/
type Grid struct {
	SizeX int
	SizeY int
	G     [][]GridEntry
}

/*
BuildGrid constructs a grid of the defined size
*/
func BuildGrid(sizeX int, sizeY int) Grid {
	var result Grid
	result.SizeX = sizeX
	result.SizeY = sizeY
	result.G = make([][]GridEntry, sizeX)
	for i := 0; i < sizeX; i++ {
		result.G[i] = make([]GridEntry, sizeY)
	}
	return result
}

/*
BuildResultsGrid builds a results grid
*/
func BuildResultsGrid(dimentions []Dimension) [][]Grid {
	resultTable := make([][]Grid, NumDimentions)
	for i := 0; i < NumDimentions; i++ {
		resultTable[i] = make([]Grid, NumDimentions)
		for j := 0; j < NumDimentions; j++ {
			resultTable[i][j] = BuildGrid(dimentions[i].GridSize, dimentions[j].GridSize)
		}
	}
	return resultTable
}
