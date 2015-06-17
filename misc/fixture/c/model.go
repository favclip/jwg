package c

// test for struct with json tag and another tags definition

type Sample struct {
	A string `json:"foo!" datastore:"-"`
	B string `json:",omitempty" datastore:",noindex" endpoints:"req"`
}
