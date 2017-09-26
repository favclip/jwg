//go:generate jwg -type Sample -output model_json.go -noOmitemptyFieldType=bool,int .

package t

type Sample struct {
	A string
	B string `json:",omitempty"`
	C bool
	D bool `json:",omitempty"`
	E int
	F int `json:",omitempty"`
	G int64
}
