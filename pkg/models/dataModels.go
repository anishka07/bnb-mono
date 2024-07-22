package models

type DataModel struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      interface{}
	CSRFToken string
	Warning   string
	Error     string
}
