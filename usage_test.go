package jwg

import (
	"encoding/json"
	"testing"

	"github.com/favclip/jwg/misc/fixture/a"
	"github.com/favclip/jwg/misc/fixture/b"
	"github.com/favclip/jwg/misc/fixture/e"
	"github.com/favclip/jwg/misc/fixture/f"
	i "github.com/favclip/jwg/misc/fixture/i"
)

func TestBasicUsage1(t *testing.T) {
	src := &a.Sample{"Foo!"}

	builder := a.NewSampleJSONBuilder()
	builder.AddAll()
	jsonStruct, err := builder.Convert(src)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	json, err := json.Marshal(jsonStruct)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(json) != `{"foo":"Foo!"}` {
		t.Log("json is not expected, actual", string(json))
		t.Fail()
	}
}

func TestBasicUsage2(t *testing.T) {
	src := &a.Sample{"Foo!"}

	builder := a.NewSampleJSONBuilder()
	builder.AddAll()
	json, err := builder.Marshal(src)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(json) != `{"foo":"Foo!"}` {
		t.Log("json is not expected, actual", string(json))
		t.Fail()
	}
}

func TestWithRemove(t *testing.T) {
	src := &b.Sample{"A", "B", 0, 1, "E", 0, nil}

	builder := b.NewSampleJSONBuilder()
	builder.AddAll()
	builder.Remove(builder.A)
	json, err := builder.Marshal(src)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	// A was removed. C was omitted. E was ignored.
	if string(json) != `{"b":"B","d":"1"}` {
		t.Log("json is not expected, actual", string(json))
		t.Fail()
	}
}

func TestWithAdd(t *testing.T) {
	src := &b.Sample{"A", "B", 0, 1, "E", 0, nil}

	builder := b.NewSampleJSONBuilder()
	builder.Add(builder.D)
	json, err := builder.Marshal(src)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(json) != `{"d":"1"}` {
		t.Log("json is not expected, actual", string(json))
		t.Fail()
	}
}

func TestReplacePropertyBuilder1(t *testing.T) {
	src := &a.Sample{"Foo"}

	builder := a.NewSampleJSONBuilder()
	builder.Foo.Encoder = func(src *a.Sample, dest *a.SampleJSON) error {
		dest.Foo = src.Foo + "!!!"
		return nil
	}
	builder.AddAll()
	json, err := builder.Marshal(src)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(json) != `{"foo":"Foo!!!"}` {
		t.Log("json is not expected, actual", string(json))
		t.Fail()
	}
}

func TestReplacePropertyBuilder2(t *testing.T) {
	builder := i.NewPeopleJSONBuilder()
	builder.AddAll()
	builder.Remove(builder.ShowPrivateInfo)
	builder.List.Encoder = func(src *i.People, dest *i.PeopleJSON) error {
		if src == nil {
			return nil
		}
		b := i.NewPersonJSONBuilder().AddAll()
		if !src.ShowPrivateInfo {
			b.Remove(b.Password)
		}
		list, err := b.ConvertList(src.List)
		if err != nil {
			return err
		}
		dest.List = ([]*i.PersonJSON)(list)
		return nil
	}

	people := &i.People{ShowPrivateInfo: false, List: []*i.Person{&i.Person{
		Name:     "vvakame",
		Age:      30,
		Password: "pw",
	}}}
	json, err := builder.Convert(people)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(json.List) != 1 {
		t.Log("json.List is not expected, actual", len(json.List))
		t.Fail()
	}
	if json.List[0].Password != "" {
		t.Log("password is not expected, actual", json.List[0].Password)
		t.Fail()
	}
}

func TestWithPointerField(t *testing.T) {
	str := "Hi!"
	src := &e.Sample{
		&str,
		&e.Foo{},
	}

	builder := e.NewSampleJSONBuilder()
	builder.AddAll()
	json, err := builder.Marshal(src)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(json) != `{"str":"Hi!","foo":{}}` {
		t.Log("json is not expected, actual", string(json))
		t.Fail()
	}
}

func TestWithImportStatement(t *testing.T) {
	src := &f.SampleF{nil, nil}

	builder := f.NewSampleFJSONBuilder()
	builder.AddAll()
	json, err := builder.Marshal(src)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(json) != `{}` {
		t.Log("json is not expected, actual", string(json))
		t.Fail()
	}
}

func TestConvertJsonStructToVanillaStruct(t *testing.T) {
	json := &a.SampleJSON{"foo!"}
	src, err := json.Convert()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if src.Foo != "foo!" {
		t.Log("src.Foo is not expected, actual", src.Foo)
		t.Fail()
	}
}
