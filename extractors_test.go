package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wselwood/gompcreader"
)

type parameterTestCase struct {
	inSemimajorAxis       float64
	inOrbitalEccentricity float64
	out                   string
	outCell               int
}

var apohelionTestCases = []parameterTestCase{
	{0.5, 0.1, "0.5", 5},
	{1.0, 0.0, "1.0", 10},
	{1.0, 0.01, "1.0", 10},
	{1.0, 0.1, "1.1", 11},
	{5.0, 0.2, "6.0", 60},
	{5.0, 0.1, "5.5", 55},
	{10.0, 0.1, "11.0", -1},
	{10.0, 0, "10.0", 100},
	{9.0, 0.1, "9.9", 99},
	{9.0, 0.2, "10.8", -1},
}

/*
TestApohelionExtractor tests basic apohelion results
*/
func TestApohelionExtractor(t *testing.T) {
	extractor := ApohelionExtractor{10, 10.0}
	for _, tt := range apohelionTestCases {
		var input gompcreader.MinorPlanet

		input.SemimajorAxis = tt.inSemimajorAxis
		input.OrbitalEccentricity = tt.inOrbitalEccentricity
		assert.Equal(t, tt.outCell, extractor.ExtractCell(&input), "incorrect cell")
		assert.Equal(t, tt.out, extractor.Extract(&input), "incorrect message")
	}
}

var perihelionTestCases = []parameterTestCase{
	{0.5, 0.1, "0.4", 4},
	{1.0, 0.0, "1.0", 10},
	{1.0, 0.01, "0.9", 9},
	{1.0, 0.1, "0.9", 9},
	{5.0, 0.2, "4.0", 40},
	{5.0, 0.1, "4.5", 45},
	{10.0, 0.1, "9.0", 90},
	{10.0, 0, "10.0", 100},
	{10.1, 0.1, "9.0", 90},
	{10.2, 0.01, "10.0", -1},
}

/*
TestApohelionExtractor tests basic apohelion results
*/
func TestPerihelionExtractor(t *testing.T) {
	extractor := PerihelionExtractor{10, 10.0}
	for _, tt := range perihelionTestCases {
		var input gompcreader.MinorPlanet

		input.SemimajorAxis = tt.inSemimajorAxis
		input.OrbitalEccentricity = tt.inOrbitalEccentricity
		assert.Equal(t, tt.outCell, extractor.ExtractCell(&input), "incorrect cell")
		assert.Equal(t, tt.out, extractor.Extract(&input), "incorrect message")
	}
}

type yearTestCase struct {
	in      int64
	out     string
	outCell int
}

var yearOfFirstObsTestCases = []yearTestCase{
	{2015, "2015", 100},
	{2000, "2000", 85},
	{1916, "1916", 1},
	{1915, "1915", 0},
	{1914, "1914", -1},
}

func TestYearOfFirstObsExtractor(t *testing.T) {
	extractor := YearOfFirstObsExtractor{1915}
	for _, tt := range yearOfFirstObsTestCases {
		var input gompcreader.MinorPlanet
		input.YearOfFirstObservation = tt.in
		assert.Equal(t, tt.outCell, extractor.ExtractCell(&input), "incorrect cell %d %d", tt.in, tt.outCell)
		assert.Equal(t, tt.out, extractor.Extract(&input), "incorrect message %d %s", tt.in, tt.out)
	}
}

var yearOfLastObsTestCases = []yearTestCase{
	{2015, "2015", 100},
	{2000, "2000", 85},
	{1916, "1916", 1},
	{1915, "1915", 0},
	{1914, "1914", -1},
}

func TestYearOfLastObsExtractor(t *testing.T) {
	extractor := YearOfLastObsExtractor{1915}
	for _, tt := range yearOfLastObsTestCases {
		var input gompcreader.MinorPlanet
		input.YearOfLastObservation = tt.in
		assert.Equal(t, tt.outCell, extractor.ExtractCell(&input), "incorrect cell %d %d", tt.in, tt.outCell)
		assert.Equal(t, tt.out, extractor.Extract(&input), "incorrect message %d %s", tt.in, tt.out)
	}
}

type floatTestCase struct {
	in      float64
	out     string
	outCell int
}

var orbitalEccentricityTestCases = []floatTestCase{
	{0.99, "0.99", 99},
	{0.554, "0.55", 55},
	{0.555, "0.55", 55},
	{0.0, "0.00", 0},
}

func TestOrbitalEccentricityExtractor(t *testing.T) {
	extractor := OrbitalEccentricityExtractor{}
	for _, tt := range orbitalEccentricityTestCases {
		var input gompcreader.MinorPlanet
		input.OrbitalEccentricity = tt.in

		assert.Equal(t, tt.outCell, extractor.ExtractCell(&input), "incorrect cell %f %d", tt.in, tt.outCell)
		assert.Equal(t, tt.out, extractor.Extract(&input), "incorrect message %f %s", tt.in, tt.out)
	}
}

//InclinationToTheEclipticExtractor
var inclinationTestCases = []floatTestCase{
	{12.0, "12.0", 6},
	{45.4, "44.0", 22},
	{45.6, "44.0", 22},
	{44.4, "44.0", 22},
	{43.6, "42.0", 21},
}

func TestInclinationToTheEclipticExtractor(t *testing.T) {
	extractor := InclinationToTheEclipticExtractor{}
	for _, tt := range inclinationTestCases {
		var input gompcreader.MinorPlanet
		input.InclinationToTheEcliptic = tt.in

		assert.Equal(t, tt.outCell, extractor.ExtractCell(&input), "incorrect cell %f %d", tt.in, tt.outCell)
		assert.Equal(t, tt.out, extractor.Extract(&input), "incorrect message %f %s", tt.in, tt.out)
	}
}

//SemimajorAxisExtractor
var semimajorAxisTestCases = []floatTestCase{
	{12.34, "12.3", -1},
	{5.34, "5.3", 53},
}

func TestSemimajorAxisExtractor(t *testing.T) {
	extractor := SemimajorAxisExtractor{10.0, 10}
	for _, tt := range semimajorAxisTestCases {
		var input gompcreader.MinorPlanet
		input.SemimajorAxis = tt.in

		assert.Equal(t, tt.outCell, extractor.ExtractCell(&input), "incorrect cell %f %d", tt.in, tt.outCell)
		assert.Equal(t, tt.out, extractor.Extract(&input), "incorrect message %f %s", tt.in, tt.out)
	}
}

//AbsoluteMagnitudeExtractor
var absoluteMagnitudeTestCases = []floatTestCase{
	{-1.123, "-1.0", 1},
	{1.43, "1.0", 9},
	{14.5, "14.5", 53},
}

func TestAbsoluteMagnitudeExtractor(t *testing.T) {
	extractor := AbsoluteMagnitudeExtractor{30, 10.0, 1.5, 3}
	for _, tt := range absoluteMagnitudeTestCases {
		var input gompcreader.MinorPlanet
		input.AbsoluteMagnitude = tt.in

		assert.Equal(t, tt.outCell, extractor.ExtractCell(&input), "incorrect cell %f %d", tt.in, tt.outCell)
		assert.Equal(t, tt.out, extractor.Extract(&input), "incorrect message %f %s", tt.in, tt.out)
	}
}
