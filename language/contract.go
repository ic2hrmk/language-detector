package language

type Detector interface {
	GetLanguage(phrase string) (language string, err error)
}