// generated by jwg -type Sample -output misc/fixture/l/model_json.go -transcripttag swagger,includes misc/fixture/l; DO NOT EDIT

package l

import (
	"encoding/json"
)

// SampleJSON is jsonized struct for Sample.
type SampleJSON struct {
	B string `json:"b,omitempty" swagger:",enum=ok|ng"`
	C string `json:"c,omitempty" swagger:",enum=ok|ng"`
	D string `json:"d,omitempty" swagger:",enum=ok|ng" includes:"test"`
	E string `json:"e,omitempty" swagger:",enum=ok|ng"`
}

// SampleJSONList is synonym about []*SampleJSON.
type SampleJSONList []*SampleJSON

// SamplePropertyEncoder is property encoder for [1]sJSON.
type SamplePropertyEncoder func(src *Sample, dest *SampleJSON) error

// SamplePropertyDecoder is property decoder for [1]sJSON.
type SamplePropertyDecoder func(src *SampleJSON, dest *Sample) error

// SamplePropertyInfo stores property information.
type SamplePropertyInfo struct {
	name    string
	Encoder SamplePropertyEncoder
	Decoder SamplePropertyDecoder
}

// SampleJSONBuilder convert between Sample to SampleJSON mutually.
type SampleJSONBuilder struct {
	_properties        map[string]*SamplePropertyInfo
	_jsonPropertyMap   map[string]*SamplePropertyInfo
	_structPropertyMap map[string]*SamplePropertyInfo
	B                  *SamplePropertyInfo
	C                  *SamplePropertyInfo
	D                  *SamplePropertyInfo
	E                  *SamplePropertyInfo
}

// NewSampleJSONBuilder make new SampleJSONBuilder.
func NewSampleJSONBuilder() *SampleJSONBuilder {
	jb := &SampleJSONBuilder{
		_properties:        map[string]*SamplePropertyInfo{},
		_jsonPropertyMap:   map[string]*SamplePropertyInfo{},
		_structPropertyMap: map[string]*SamplePropertyInfo{},
		B: &SamplePropertyInfo{
			name: "B",
			Encoder: func(src *Sample, dest *SampleJSON) error {
				if src == nil {
					return nil
				}
				dest.B = src.B
				return nil
			},
			Decoder: func(src *SampleJSON, dest *Sample) error {
				if src == nil {
					return nil
				}
				dest.B = src.B
				return nil
			},
		},
		C: &SamplePropertyInfo{
			name: "C",
			Encoder: func(src *Sample, dest *SampleJSON) error {
				if src == nil {
					return nil
				}
				dest.C = src.C
				return nil
			},
			Decoder: func(src *SampleJSON, dest *Sample) error {
				if src == nil {
					return nil
				}
				dest.C = src.C
				return nil
			},
		},
		D: &SamplePropertyInfo{
			name: "D",
			Encoder: func(src *Sample, dest *SampleJSON) error {
				if src == nil {
					return nil
				}
				dest.D = src.D
				return nil
			},
			Decoder: func(src *SampleJSON, dest *Sample) error {
				if src == nil {
					return nil
				}
				dest.D = src.D
				return nil
			},
		},
		E: &SamplePropertyInfo{
			name: "E",
			Encoder: func(src *Sample, dest *SampleJSON) error {
				if src == nil {
					return nil
				}
				dest.E = src.E
				return nil
			},
			Decoder: func(src *SampleJSON, dest *Sample) error {
				if src == nil {
					return nil
				}
				dest.E = src.E
				return nil
			},
		},
	}
	jb._structPropertyMap["B"] = jb.B
	jb._jsonPropertyMap["b"] = jb.B
	jb._structPropertyMap["C"] = jb.C
	jb._jsonPropertyMap["c"] = jb.C
	jb._structPropertyMap["D"] = jb.D
	jb._jsonPropertyMap["d"] = jb.D
	jb._structPropertyMap["E"] = jb.E
	jb._jsonPropertyMap["e"] = jb.E
	return jb
}

// AddAll adds all property to SampleJSONBuilder.
func (b *SampleJSONBuilder) AddAll() *SampleJSONBuilder {
	b._properties["B"] = b.B
	b._properties["C"] = b.C
	b._properties["D"] = b.D
	b._properties["E"] = b.E
	return b
}

// Add specified property to SampleJSONBuilder.
func (b *SampleJSONBuilder) Add(infos ...*SamplePropertyInfo) *SampleJSONBuilder {
	for _, info := range infos {
		b._properties[info.name] = info
	}
	return b
}

// AddByJSONNames add properties to SampleJSONBuilder by JSON property name. if name is not in the builder, it will ignore.
func (b *SampleJSONBuilder) AddByJSONNames(names ...string) *SampleJSONBuilder {
	for _, name := range names {
		info := b._jsonPropertyMap[name]
		if info == nil {
			continue
		}
		b._properties[info.name] = info
	}
	return b
}

// AddByNames add properties to SampleJSONBuilder by struct property name. if name is not in the builder, it will ignore.
func (b *SampleJSONBuilder) AddByNames(names ...string) *SampleJSONBuilder {
	for _, name := range names {
		info := b._structPropertyMap[name]
		if info == nil {
			continue
		}
		b._properties[info.name] = info
	}
	return b
}

// Remove specified property to SampleJSONBuilder.
func (b *SampleJSONBuilder) Remove(infos ...*SamplePropertyInfo) *SampleJSONBuilder {
	for _, info := range infos {
		delete(b._properties, info.name)
	}
	return b
}

// RemoveByJSONNames remove properties to SampleJSONBuilder by JSON property name. if name is not in the builder, it will ignore.
func (b *SampleJSONBuilder) RemoveByJSONNames(names ...string) *SampleJSONBuilder {

	for _, name := range names {
		info := b._jsonPropertyMap[name]
		if info == nil {
			continue
		}
		delete(b._properties, info.name)
	}
	return b
}

// RemoveByNames remove properties to SampleJSONBuilder by struct property name. if name is not in the builder, it will ignore.
func (b *SampleJSONBuilder) RemoveByNames(names ...string) *SampleJSONBuilder {
	for _, name := range names {
		info := b._structPropertyMap[name]
		if info == nil {
			continue
		}
		delete(b._properties, info.name)
	}
	return b
}

// Convert specified non-JSON object to JSON object.
func (b *SampleJSONBuilder) Convert(orig *Sample) (*SampleJSON, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &SampleJSON{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// ConvertList specified non-JSON slice to JSONList.
func (b *SampleJSONBuilder) ConvertList(orig []*Sample) (SampleJSONList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(SampleJSONList, len(orig))
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
func (orig *SampleJSON) Convert() (*Sample, error) {
	ret := &Sample{}

	b := NewSampleJSONBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// Convert specified JSONList to non-JSON slice.
func (jsonList SampleJSONList) Convert() ([]*Sample, error) {
	orig := ([]*SampleJSON)(jsonList)

	list := make([]*Sample, len(orig))
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
func (b *SampleJSONBuilder) Marshal(orig *Sample) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}
