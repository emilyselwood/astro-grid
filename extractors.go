package main

import (
	"fmt"

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
	if in.YearOfFirstObservation > extractor.minValue {
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

func scaleAxis(in float64, maxValue float64, multiplier float64) int32 {
	if in <= maxValue {
		return int32(in * multiplier)
	}
	return -1
}
