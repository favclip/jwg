package l

// test for struct with json tag and others

type Sample struct {
	A string  `json:"-"`
	B string  `json:",omitempty" swagger:",enum=ok|ng"`
	C string  `swagger:",enum=ok|ng"`
	D string  `swagger:",enum=ok|ng" includes:"test"`
	E string  `swagger:",enum=ok|ng" excludes:"test"`
}
