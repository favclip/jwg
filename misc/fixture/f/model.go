package f

// test for import statement support

import (
	"github.com/favclip/jwg/misc/fixture/a"
	bravo "github.com/favclip/jwg/misc/fixture/b"
	// . "github.com/favclip/jwg/misc/fixture/c" // unsupported!
)

type SampleF struct {
	A *a.Sample
	B *bravo.Sample
	// C *Sample
}
