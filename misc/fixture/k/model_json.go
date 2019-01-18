// generated by jwg -output misc/fixture/k/model_json.go misc/fixture/k; DO NOT EDIT

package k

import (
	"encoding/json"
	o1 "github.com/favclip/jwg/misc/other/v1"
	o2 "github.com/favclip/jwg/misc/other/v2"
)

// FooJSON is jsonized struct for Foo.
type FooJSON struct {
	Test *o1.Test `json:"test,omitempty"`
}

// FooJSONList is synonym about []*FooJSON.
type FooJSONList []*FooJSON

// FooPropertyEncoder is property encoder for [1]sJSON.
type FooPropertyEncoder func(src *Foo, dest *FooJSON) error

// FooPropertyDecoder is property decoder for [1]sJSON.
type FooPropertyDecoder func(src *FooJSON, dest *Foo) error

// FooPropertyInfo stores property information.
type FooPropertyInfo struct {
	fieldName string
	jsonName  string
	Encoder   FooPropertyEncoder
	Decoder   FooPropertyDecoder
}

// FieldName returns struct field name of property.
func (info *FooPropertyInfo) FieldName() string {
	return info.fieldName
}

// JSONName returns json field name of property.
func (info *FooPropertyInfo) JSONName() string {
	return info.jsonName
}

// FooJSONBuilder convert between Foo to FooJSON mutually.
type FooJSONBuilder struct {
	_properties        map[string]*FooPropertyInfo
	_jsonPropertyMap   map[string]*FooPropertyInfo
	_structPropertyMap map[string]*FooPropertyInfo
	Test               *FooPropertyInfo
}

// NewFooJSONBuilder make new FooJSONBuilder.
func NewFooJSONBuilder() *FooJSONBuilder {
	jb := &FooJSONBuilder{
		_properties:        map[string]*FooPropertyInfo{},
		_jsonPropertyMap:   map[string]*FooPropertyInfo{},
		_structPropertyMap: map[string]*FooPropertyInfo{},
		Test: &FooPropertyInfo{
			fieldName: "Test",
			jsonName:  "test",
			Encoder: func(src *Foo, dest *FooJSON) error {
				if src == nil {
					return nil
				}
				dest.Test = src.Test
				return nil
			},
			Decoder: func(src *FooJSON, dest *Foo) error {
				if src == nil {
					return nil
				}
				dest.Test = src.Test
				return nil
			},
		},
	}
	jb._structPropertyMap["Test"] = jb.Test
	jb._jsonPropertyMap["test"] = jb.Test
	return jb
}

// Properties returns all properties on FooJSONBuilder.
func (b *FooJSONBuilder) Properties() []*FooPropertyInfo {
	return []*FooPropertyInfo{
		b.Test,
	}
}

// AddAll adds all property to FooJSONBuilder.
func (b *FooJSONBuilder) AddAll() *FooJSONBuilder {
	b._properties["Test"] = b.Test
	return b
}

// Add specified property to FooJSONBuilder.
func (b *FooJSONBuilder) Add(info *FooPropertyInfo) *FooJSONBuilder {
	b._properties[info.fieldName] = info
	return b
}

// AddByJSONNames add properties to FooJSONBuilder by JSON property name. if name is not in the builder, it will ignore.
func (b *FooJSONBuilder) AddByJSONNames(names ...string) *FooJSONBuilder {
	for _, name := range names {
		info := b._jsonPropertyMap[name]
		if info == nil {
			continue
		}
		b._properties[info.fieldName] = info
	}
	return b
}

// AddByNames add properties to FooJSONBuilder by struct property name. if name is not in the builder, it will ignore.
func (b *FooJSONBuilder) AddByNames(names ...string) *FooJSONBuilder {
	for _, name := range names {
		info := b._structPropertyMap[name]
		if info == nil {
			continue
		}
		b._properties[info.fieldName] = info
	}
	return b
}

// Remove specified property to FooJSONBuilder.
func (b *FooJSONBuilder) Remove(info *FooPropertyInfo) *FooJSONBuilder {
	delete(b._properties, info.fieldName)
	return b
}

// RemoveByJSONNames remove properties to FooJSONBuilder by JSON property name. if name is not in the builder, it will ignore.
func (b *FooJSONBuilder) RemoveByJSONNames(names ...string) *FooJSONBuilder {

	for _, name := range names {
		info := b._jsonPropertyMap[name]
		if info == nil {
			continue
		}
		delete(b._properties, info.fieldName)
	}
	return b
}

// RemoveByNames remove properties to FooJSONBuilder by struct property name. if name is not in the builder, it will ignore.
func (b *FooJSONBuilder) RemoveByNames(names ...string) *FooJSONBuilder {
	for _, name := range names {
		info := b._structPropertyMap[name]
		if info == nil {
			continue
		}
		delete(b._properties, info.fieldName)
	}
	return b
}

// Convert specified non-JSON object to JSON object.
func (b *FooJSONBuilder) Convert(orig *Foo) (*FooJSON, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &FooJSON{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// ConvertList specified non-JSON slice to JSONList.
func (b *FooJSONBuilder) ConvertList(orig []*Foo) (FooJSONList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(FooJSONList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

// Convert specified JSON object to non-JSON object.
func (orig *FooJSON) Convert() (*Foo, error) {
	ret := &Foo{}

	b := NewFooJSONBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// Convert specified JSONList to non-JSON slice.
func (jsonList FooJSONList) Convert() ([]*Foo, error) {
	orig := ([]*FooJSON)(jsonList)

	list := make([]*Foo, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

// Marshal non-JSON object to JSON string.
func (b *FooJSONBuilder) Marshal(orig *Foo) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}

// BarJSON is jsonized struct for Bar.
type BarJSON struct {
	Tests []*o2.Test `json:"tests,omitempty"`
}

// BarJSONList is synonym about []*BarJSON.
type BarJSONList []*BarJSON

// BarPropertyEncoder is property encoder for [1]sJSON.
type BarPropertyEncoder func(src *Bar, dest *BarJSON) error

// BarPropertyDecoder is property decoder for [1]sJSON.
type BarPropertyDecoder func(src *BarJSON, dest *Bar) error

// BarPropertyInfo stores property information.
type BarPropertyInfo struct {
	fieldName string
	jsonName  string
	Encoder   BarPropertyEncoder
	Decoder   BarPropertyDecoder
}

// FieldName returns struct field name of property.
func (info *BarPropertyInfo) FieldName() string {
	return info.fieldName
}

// JSONName returns json field name of property.
func (info *BarPropertyInfo) JSONName() string {
	return info.jsonName
}

// BarJSONBuilder convert between Bar to BarJSON mutually.
type BarJSONBuilder struct {
	_properties        map[string]*BarPropertyInfo
	_jsonPropertyMap   map[string]*BarPropertyInfo
	_structPropertyMap map[string]*BarPropertyInfo
	Tests              *BarPropertyInfo
}

// NewBarJSONBuilder make new BarJSONBuilder.
func NewBarJSONBuilder() *BarJSONBuilder {
	jb := &BarJSONBuilder{
		_properties:        map[string]*BarPropertyInfo{},
		_jsonPropertyMap:   map[string]*BarPropertyInfo{},
		_structPropertyMap: map[string]*BarPropertyInfo{},
		Tests: &BarPropertyInfo{
			fieldName: "Tests",
			jsonName:  "tests",
			Encoder: func(src *Bar, dest *BarJSON) error {
				if src == nil {
					return nil
				}
				dest.Tests = src.Tests
				return nil
			},
			Decoder: func(src *BarJSON, dest *Bar) error {
				if src == nil {
					return nil
				}
				dest.Tests = src.Tests
				return nil
			},
		},
	}
	jb._structPropertyMap["Tests"] = jb.Tests
	jb._jsonPropertyMap["tests"] = jb.Tests
	return jb
}

// Properties returns all properties on BarJSONBuilder.
func (b *BarJSONBuilder) Properties() []*BarPropertyInfo {
	return []*BarPropertyInfo{
		b.Tests,
	}
}

// AddAll adds all property to BarJSONBuilder.
func (b *BarJSONBuilder) AddAll() *BarJSONBuilder {
	b._properties["Tests"] = b.Tests
	return b
}

// Add specified property to BarJSONBuilder.
func (b *BarJSONBuilder) Add(info *BarPropertyInfo) *BarJSONBuilder {
	b._properties[info.fieldName] = info
	return b
}

// AddByJSONNames add properties to BarJSONBuilder by JSON property name. if name is not in the builder, it will ignore.
func (b *BarJSONBuilder) AddByJSONNames(names ...string) *BarJSONBuilder {
	for _, name := range names {
		info := b._jsonPropertyMap[name]
		if info == nil {
			continue
		}
		b._properties[info.fieldName] = info
	}
	return b
}

// AddByNames add properties to BarJSONBuilder by struct property name. if name is not in the builder, it will ignore.
func (b *BarJSONBuilder) AddByNames(names ...string) *BarJSONBuilder {
	for _, name := range names {
		info := b._structPropertyMap[name]
		if info == nil {
			continue
		}
		b._properties[info.fieldName] = info
	}
	return b
}

// Remove specified property to BarJSONBuilder.
func (b *BarJSONBuilder) Remove(info *BarPropertyInfo) *BarJSONBuilder {
	delete(b._properties, info.fieldName)
	return b
}

// RemoveByJSONNames remove properties to BarJSONBuilder by JSON property name. if name is not in the builder, it will ignore.
func (b *BarJSONBuilder) RemoveByJSONNames(names ...string) *BarJSONBuilder {

	for _, name := range names {
		info := b._jsonPropertyMap[name]
		if info == nil {
			continue
		}
		delete(b._properties, info.fieldName)
	}
	return b
}

// RemoveByNames remove properties to BarJSONBuilder by struct property name. if name is not in the builder, it will ignore.
func (b *BarJSONBuilder) RemoveByNames(names ...string) *BarJSONBuilder {
	for _, name := range names {
		info := b._structPropertyMap[name]
		if info == nil {
			continue
		}
		delete(b._properties, info.fieldName)
	}
	return b
}

// Convert specified non-JSON object to JSON object.
func (b *BarJSONBuilder) Convert(orig *Bar) (*BarJSON, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &BarJSON{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// ConvertList specified non-JSON slice to JSONList.
func (b *BarJSONBuilder) ConvertList(orig []*Bar) (BarJSONList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(BarJSONList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

// Convert specified JSON object to non-JSON object.
func (orig *BarJSON) Convert() (*Bar, error) {
	ret := &Bar{}

	b := NewBarJSONBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// Convert specified JSONList to non-JSON slice.
func (jsonList BarJSONList) Convert() ([]*Bar, error) {
	orig := ([]*BarJSON)(jsonList)

	list := make([]*Bar, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

// Marshal non-JSON object to JSON string.
func (b *BarJSONBuilder) Marshal(orig *Bar) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}
