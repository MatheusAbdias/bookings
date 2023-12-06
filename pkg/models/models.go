package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMpa  map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
