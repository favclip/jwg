package g

// test for import statement support

type Inner struct {
	A string
}

type Sample struct {
	A1  Inner
	A2  *Inner
	As1 []Inner
	As2 []*Inner
	As3 *[]Inner
	As4 *[]*Inner
}
