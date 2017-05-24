package m

// test for struct with shouldemit json tag

type Sample struct {
	A string `json:",omitempty"`
	B string
	C string `json:",shouldemit"`
}
