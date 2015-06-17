//go:generate jwg -output model_json.go .

package j

// +jwg
type Foo struct {
	Tmp *Temp
	Bar
	*Buzz
	Hoge
	*Fuga
}

type Temp struct {
	Temp1 string
}

type Bar struct {
	Bar1 string
}

type Buzz struct {
	Buzz1 string
}

// +jwg
type Hoge struct {
	Hoge1 string
}

// +jwg
type Fuga struct {
	Fuga1 string
}
