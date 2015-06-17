package j

import (
	"encoding/json"
	"testing"
)

func TestModel(t *testing.T) {
	obj := &Foo{
		Tmp:  &Temp{"a"},
		Bar:  Bar{"b"},
		Buzz: &Buzz{"c"},
		Hoge: Hoge{"d"},
		Fuga: &Fuga{"d"},
	}
	jsonObj, err := NewFooJsonBuilder().AddAll().Convert(obj)
	if err != nil {
		t.Fatal(err.Error())
	}
	b, err := json.Marshal(jsonObj)
	if err != nil {
		t.Fatal(err.Error())
	}
	if string(b) != `{"tmp":{"Temp1":"a"},"bar":{"Bar1":"b"},"buzz":{"Buzz1":"c"},"hoge":{"hoge1":"d"},"fuga":{"fuga1":"d"}}` {
		t.Errorf("not expected: %s", string(b))
	}
}
