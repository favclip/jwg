# jwg

`JSON wrapper generator` or `json web go`.

## Description

`jwg` can generate Web API friendly struct.

We need a `lowerCamelCase` property name, not a `UpperCamelCase`.
jwg generate json tag with lower camel case automatically.

Another case, We need a access control of struct property in web app.

```
type User struct {
	Name              string
	EncryptedPassword string // I want to hide this field!
}
```

```
user := &User {
	"go-chan",
	"5ch",
}
builder := NewUserJsonBuilder()
builder.AddAll()
builder.Remove(builder.EncryptedPassword)
json, _ := builder.Marshal(user)
// emit `{"name":"vvakame"}`
fmt.Println(string(json))
```

### Example

[from](https://github.com/favclip/jwg/blob/master/misc/fixture/a/model.go):

```
type Sample struct {
	Foo string
}
```

[to](https://github.com/favclip/jwg/blob/master/misc/fixture/a/model_json.go):

```
// generated!
type SampleJson struct {
	Foo string `json:"foo,omitempty"`
}
```

usage:

```
src := &Sample{"Foo!"}

builder := NewSampleJsonBuilder() // generated!
builder.AddAll()
var jsonStruct *SampleJson = builder.Convert(src)

json, err := json.Marshal(jsonStruct)
```

[other example](https://github.com/favclip/jwg/blob/master/usage_test.go).

### With `go generate`

```
$ ls -la .
total 16
drwxr-xr-x@ 4 vvakame  staff   136  1 27 17:27 .
drwxr-xr-x@ 6 vvakame  staff   204  1 27 13:14 ..
-rw-r--r--@ 1 vvakame  staff   366  1 27 17:27 model.go
$ cat model.go
//go:generate jwg -output model_json.go .

package d

// test for struct with tagged comment

// +jwg
type Sample struct {
	A string `json:"foo!" datastore:"-"`
	B string `json:",omitempty" datastore:",noindex" endpoints:"req"`
}
$ go generate
$ ls -la .
total 16
drwxr-xr-x@ 4 vvakame  staff   136  1 27 17:27 .
drwxr-xr-x@ 6 vvakame  staff   204  1 27 13:14 ..
-rw-r--r--@ 1 vvakame  staff   366  1 27 17:27 model.go
-rw-r--r--@ 1 vvakame  staff  1549  1 27 17:27 model_json.go
```

## Installation

```
$ go get github.com/favclip/jwg/cmd/jwg
$ jwg
Usage of jwg:
	jwg [flags] [directory]
	jwg [flags] files... # Must be a single package
Flags:
  -output="": output file name; default srcdir/<type>_string.go
  -type="": comma-separated list of type names; must be set
```

## Command sample

Model with type specific option.

```
$ cat misc/fixture/a/model.go
package a

// test for basic struct definition

type Sample struct {
	Foo string
}
$ jwg -type Sample -output misc/fixture/a/model_json.go misc/fixture/a
```

Model with tagged comment.

```
$ cat misc/fixture/d/model.go
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
$ jwg -output misc/fixture/d/model_json.go misc/fixture/d
```
