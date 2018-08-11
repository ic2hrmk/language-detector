package resources

import "lang-detector/language"

type LanguageDetectorService struct {
	detector language.Detector
}

func NewLanguageDetectorService(detector language.Detector) *LanguageDetectorService {
	return &LanguageDetectorService{
		detector: detector,
	}
}
