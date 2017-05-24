//go:generate jwg -type Sample -output model_json.go -noOmitempty .

package n

type Sample struct {
	A string `json:",omitempty"`
	B string
	C string `json:",shouldemit"`
}
