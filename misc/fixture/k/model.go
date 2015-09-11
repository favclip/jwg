//go:generate jwg -output model_json.go .

package k
import (
	o1 "github.com/favclip/jwg/misc/other/v1"
	o2 "github.com/favclip/jwg/misc/other/v2"
)

// +jwg
type Foo struct {
	Test *o1.Test
}
// +jwg
type Bar struct {
	Tests []*o2.Test
}
