//go:generate jwg -output model_json.go .

package d

// test for struct with tagged comment

// +jwg
type Sample struct {
	A string `json:"foo!" datastore:"-"`
	B string `json:",omitempty" datastore:",noindex" endpoints:"req"`
}

type IgnoredSample struct {
	A string `json:"foo!" datastore:"-"`
	B string `json:",omitempty" datastore:",noindex" endpoints:"req"`
}
