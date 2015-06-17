//go:generate jwg -output model_json.go .

package i

// +jwg
type Person struct {
	Name     string
	Age      int
	Password string
}

// +jwg
type People struct {
	ShowPrivateInfo bool
	List            []*Person
}
