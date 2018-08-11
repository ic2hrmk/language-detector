package naive_detector

import (
	"github.com/abadojack/whatlanggo"
)

type NaiveDetector struct {}

//
// Detector constructor
//
func NewNaiveDetector() *NaiveDetector{
	return &NaiveDetector{}
}

//
// GetLanguage method
//
func (rcv *NaiveDetector) GetLanguage(phrase string) (string, error) {
	// TODO: tune processing
	return whatlanggo.LangToString(whatlanggo.Detect(phrase).Lang), nil
}