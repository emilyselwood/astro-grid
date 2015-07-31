package main

import (
	"fmt"
	"math"

	"github.com/wselwood/gompcreader"
)

/*
ValueExtractor is for extracting the cell this should live in.
*/
type ValueExtractor interface {
	ExtractCell(*gompcreader.MinorPlanet) int32
	Extract(*gompcreader.MinorPlanet) string
}

/*
ApohelionExtractor extracts the values for Apohelion
*/
type ApohelionExtractor struct {
	maxValue   float64
	multiplier float64
}

/*
ExtractCell extracts the cell value for Apohelion
*/
func (extractor *ApohelionExtractor) ExtractCell(in *gompcreader.MinorPlanet) int32 {
	CS := in.SemimajorAxis * in.OrbitalEccentricity
	apohelion := in.SemimajorAxis + CS
	return scaleAxis(apohelion, extractor.maxValue, extractor.multiplier)
}

/*
Extract extracts the start value for the bucket this MinorPlanet is in.
*/
func (extractor *ApohelionExtractor) Extract(in *gompcreader.MinorPlanet) string {
	CS := in.SemimajorAxis * in.OrbitalEccentricity
	apohelion := in.SemimajorAxis + CS

	return fmt.Sprintf("%3.1f", float64(int64(apohelion*extractor.multiplier))/extractor.multiplier)
}

/*
PerihelionExtractor extracts values for Perihelion
*/
type PerihelionExtractor struct {
	maxValue   float64
	multiplier float64
}

/*
ExtractCell extracts the cell value for Apohelion
*/
func (extractor *PerihelionExtractor) ExtractCell(in *gompcreader.MinorPlanet) int32 {
	CS := in.SemimajorAxis * in.OrbitalEccentricity
	apohelion := in.SemimajorAxis - CS
	return scaleAxis(apohelion, extractor.maxValue, extractor.multiplier)
}

/*
Extract extracts the start value for the bucket this MinorPlanet is in.
*/
func (extractor *PerihelionExtractor) Extract(in *gompcreader.MinorPlanet) string {
	CS := in.SemimajorAxis * in.OrbitalEccentricity
	apohelion := in.SemimajorAxis - CS

	return fmt.Sprintf("%3.1f", float64(int64(apohelion*extractor.multiplier))/extractor.multiplier)
}

/*
YearOfFirstObsExtractor is for pulling out the year of the first obervation
*/
type YearOfFirstObsExtractor struct {
	minValue int64
}

/*
ExtractCell extracts the cell number for the year of first observation.
*/
func (extractor *YearOfFirstObsExtractor) ExtractCell(in *gompcreader.MinorPlanet) int32 {
	if in.YearOfFirstObservation >= extractor.minValue {
		return int32(in.YearOfFirstObservation - extractor.minValue)
	}
	return -1
}

/*
Extract the value for the year of first observation.
*/
func (extractor *YearOfFirstObsExtractor) Extract(in *gompcreader.MinorPlanet) string {
	return fmt.Sprintf("%d", in.YearOfFirstObservation)
}

/*
YearOfLastObsExtractor extractor for getting at the year of last observation.
*/
type YearOfLastObsExtractor struct {
	minValue int64
}

/*
ExtractCell extracts the cell number for the year of first observation.
*/
func (extractor *YearOfLastObsExtractor) ExtractCell(in *gompcreader.MinorPlanet) int32 {
	if in.YearOfLastObservation >= extractor.minValue {
		return int32(in.YearOfLastObservation - extractor.minValue)
	}
	return -1
}

/*
Extract the value for the year of first observation.
*/
func (extractor *YearOfLastObsExtractor) Extract(in *gompcreader.MinorPlanet) string {
	return fmt.Sprintf("%d", in.YearOfLastObservation)
}

/*
OrbitalEccentricityExtractor does what it says on the tin.
*/
type OrbitalEccentricityExtractor struct {
}

/*
ExtractCell for the orbital eccentricity
*/
func (extractor *OrbitalEccentricityExtractor) ExtractCell(in *gompcreader.MinorPlanet) int32 {
	return int32(in.OrbitalEccentricity * 100)
}

/*
Extract the orbital eccentricity
*/
func (extractor *OrbitalEccentricityExtractor) Extract(in *gompcreader.MinorPlanet) string {
	return fmt.Sprintf("%2.2f", in.OrbitalEccentricity)
}

/*
InclinationToTheEclipticExtractor does what it says on the tin.
*/
type InclinationToTheEclipticExtractor struct {
}

/*
ExtractCell for the InclinationToTheEcliptic
*/
func (extractor *InclinationToTheEclipticExtractor) ExtractCell(in *gompcreader.MinorPlanet) int32 {
	return int32(in.InclinationToTheEcliptic / 2)
}

/*
Extract the InclinationToTheEcliptic
*/
func (extractor *InclinationToTheEclipticExtractor) Extract(in *gompcreader.MinorPlanet) string {
	return fmt.Sprintf("%f", in.InclinationToTheEcliptic)
}

/*
SemimajorAxisExtractor also does what it says on the tin.
*/
type SemimajorAxisExtractor struct {
	maxValue   float64
	multiplier float64
}

/*
ExtractCell for the SemimajorAxisExtractor
*/
func (extractor *SemimajorAxisExtractor) ExtractCell(in *gompcreader.MinorPlanet) int32 {
	return scaleAxis(in.SemimajorAxis, extractor.maxValue, extractor.multiplier)
}

/*
Extract the SemimajorAxisExtractor
*/
func (extractor *SemimajorAxisExtractor) Extract(in *gompcreader.MinorPlanet) string {
	return fmt.Sprintf("%3.1f", float64(int64(in.SemimajorAxis*extractor.multiplier))/extractor.multiplier)
}

/*
AbsoluteMagnitudeExtractor also does what it says on the tin.
*/
type AbsoluteMagnitudeExtractor struct {
	maxValue   float64
	multiplier float64
}

/*
ExtractCell for the InclinationToTheEcliptic
*/
func (extractor *AbsoluteMagnitudeExtractor) ExtractCell(in *gompcreader.MinorPlanet) int32 {
	return int32(float64(int64(in.AbsoluteMagnitude*extractor.multiplier)) / extractor.multiplier)
}

/*
Extract the InclinationToTheEcliptic
*/
func (extractor *AbsoluteMagnitudeExtractor) Extract(in *gompcreader.MinorPlanet) string {
	return fmt.Sprintf("%3.1f", in.AbsoluteMagnitude)
}

func scaleAxis(in float64, maxValue float64, multiplier float64) int32 {
	if in <= maxValue {
		return int32(in * multiplier)
	}
	return -1
}

func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor((f*shift)+.5) / shift
}
