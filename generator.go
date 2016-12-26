package jwg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/favclip/genbase"
)

// ErrInvalidTranscriptTags means invalid transcriptTag included.
var ErrInvalidTranscriptTags = errors.New("do not contains json tag in transcriptTagNames")

// BuildSource represents source code of assembling..
type BuildSource struct {
	g                  *genbase.Generator
	pkg                *genbase.PackageInfo
	typeInfos          genbase.TypeInfos
	transcriptTagNames []string // e.g. swagger etc... copy struct tag to *JSON struct

	Structs []*BuildStruct
}

// BuildStruct represents struct of assembling..
type BuildStruct struct {
	parent   *BuildSource
	typeInfo *genbase.TypeInfo

	Fields []*BuildField
}

// BuildField represents field of BuildStruct.
type BuildField struct {
	parent    *BuildStruct
	fieldInfo *genbase.FieldInfo

	Name           string
	Embed          bool
	Tag            *BuildTag
	TranscriptTags []string
}

// BuildTag represents tag of BuildField.
type BuildTag struct {
	field *BuildField

	Name      string
	Ignore    bool // e.g. Secret string `json:"-"`
	DoNotEmit bool // e.g. Field int `json:",omitempty"`
	String    bool // e.g. Int64String int64 `json:",string"`
}

// Parse construct *BuildSource from package & type information.
func Parse(pkg *genbase.PackageInfo, typeInfos genbase.TypeInfos, transcriptTagNames []string) (*BuildSource, error) {
	for _, tagName := range transcriptTagNames {
		if tagName == "json" {
			return nil, ErrInvalidTranscriptTags
		}
	}

	bu := &BuildSource{
		g:                  genbase.NewGenerator(pkg),
		pkg:                pkg,
		typeInfos:          typeInfos,
		transcriptTagNames: transcriptTagNames,
	}

	bu.g.AddImport("encoding/json", "")

	for _, typeInfo := range typeInfos {
		err := bu.parseStruct(typeInfo)
		if err != nil {
			return nil, err
		}
	}

	return bu, nil
}

func (b *BuildSource) parseStruct(typeInfo *genbase.TypeInfo) error {
	structType, err := typeInfo.StructType()
	if err != nil {
		return err
	}

	st := &BuildStruct{
		parent:   b,
		typeInfo: typeInfo,
	}

	for _, fieldInfo := range structType.FieldInfos() {
		if len := len(fieldInfo.Names); len == 0 {
			// embedded struct in outer struct or multiply field declarations
			// https://play.golang.org/p/bcxbdiMyP4
			typeName, err := genbase.ExprToBaseTypeName(fieldInfo.Type)
			if err != nil {
				return err
			}
			err = b.parseField(st, typeInfo, fieldInfo, typeName, true)
			if err != nil {
				return err
			}
			continue
		} else {
			for _, nameIdent := range fieldInfo.Names {
				err := b.parseField(st, typeInfo, fieldInfo, nameIdent.Name, false)
				if err != nil {
					return err
				}
			}
		}
	}

	b.Structs = append(b.Structs, st)

	return nil
}

func (b *BuildSource) parseField(st *BuildStruct, typeInfo *genbase.TypeInfo, fieldInfo *genbase.FieldInfo, name string, embed bool) error {
	field := &BuildField{
		parent:    st,
		fieldInfo: fieldInfo,
		Name:      name,
		Embed:     embed,
	}
	st.Fields = append(st.Fields, field)

	tag := &BuildTag{
		field: field,
	}
	{
		typeName, err := genbase.ExprToTypeName(fieldInfo.Type)
		if err != nil {
			return err
		}
		if typeName == "int64" {
			tag.String = true
		}
	}
	field.Tag = tag
	if field.Embed {
		// do not emit tag name for embed struct in default behavior
	} else if strings.IndexFunc(field.Name, func(r rune) bool { return unicode.IsLower(r) }) == -1 {
		// lower char is not contains.
		// convert to lower case.
		// e.g. ID -> id
		tag.Name = strings.ToLower(field.Name)
	} else {
		tag.Name = strings.ToLower(field.Name[:1]) + field.Name[1:]
	}

	if fieldInfo.Tag != nil {
		tagText := fieldInfo.Tag.Value[1 : len(fieldInfo.Tag.Value)-1]
		tagKeys := genbase.GetKeys(tagText)
		structTag := reflect.StructTag(tagText)

		for _, key := range tagKeys {
			for _, target := range b.transcriptTagNames {
				if target == key {
					tag := structTag.Get(key)
					field.TranscriptTags = append(field.TranscriptTags, fmt.Sprintf(`%s:"%s"`, key, tag))
				}
			}
			if key != "json" {
				continue
			}

			jsonTag := structTag.Get("json")
			if jsonTag == "-" {
				tag.Ignore = true
				continue
			}
			if idx := strings.Index(jsonTag, ","); idx == -1 {
				tag.Name = jsonTag
			} else {
				for idx != -1 || jsonTag != "" {
					value := jsonTag
					if idx != -1 {
						value = jsonTag[:idx]
						jsonTag = jsonTag[idx+1:]
					} else {
						jsonTag = jsonTag[len(value):]
					}
					idx = strings.Index(jsonTag, ",")

					if value == "omitempty" {
						tag.DoNotEmit = true
					} else if value == "string" {
						tag.String = true
					} else if value != "" {
						if strings.IndexFunc(value, func(r rune) bool { return unicode.IsLower(r) }) == -1 {
							// lower char is not contains.
							// convert to lower case.
							// e.g. ID -> id
							tag.Name = strings.ToLower(value)
						} else {
							tag.Name = strings.ToLower(value[:1]) + value[1:]
						}
					}
				}
			}
		}
	}

	needImport, packageIdent := genbase.IsReferenceToOtherPackage(fieldInfo.Type)
	if needImport && tag.Ignore == false {
		importSpec := typeInfo.FileInfo.FindImportSpecByIdent(packageIdent)
		if importSpec != nil && importSpec.Name != nil {
			b.g.AddImport(importSpec.Path.Value, importSpec.Name.Name)
		} else if importSpec != nil {
			b.g.AddImport(importSpec.Path.Value, "")
		}
	}

	return nil
}

// Emit generate wrapper code.
func (b *BuildSource) Emit(args *[]string) ([]byte, error) {
	b.g.PrintHeader("jwg", args)

	for _, st := range b.Structs {
		err := st.emit(b.g)
		if err != nil {
			return nil, err
		}
	}

	return b.g.Format()
}

func (st *BuildStruct) emit(g *genbase.Generator) error {
	g.Printf("// %[1]sJSON is jsonized struct for %[1]s.\n", st.Name())

	// generate FooJSON struct from Foo struct
	g.Printf("type %sJSON struct {\n", st.Name())
	for _, field := range st.Fields {
		if field.Tag.Ignore {
			continue
		}
		postfix := ""
		if field.WithJWG() {
			postfix = "JSON"
		}
		tagString := field.Tag.TagString()
		if len(field.TranscriptTags) != 0 {
			tagString += " " + strings.Join(field.TranscriptTags, " ")
		}
		if tagString != "" {
			tagString = fmt.Sprintf("`%s`", tagString)
		}
		if field.Embed {
			g.Printf("%s%s %s\n", field.fieldInfo.TypeName(), postfix, tagString)
		} else {
			g.Printf("%s %s%s %s\n", field.Name, field.fieldInfo.TypeName(), postfix, tagString)
		}
	}
	g.Printf("}\n\n")

	g.Printf("// %[1]sJSONList is synonym about []*%[1]sJSON.\n", st.Name())
	g.Printf("type %[1]sJSONList []*%[1]sJSON\n\n", st.Name())

	// generate property builder
	g.Printf("// %[1]sPropertyEncoder is property encoder for [1]sJSON.\n", st.Name())
	g.Printf("type %[1]sPropertyEncoder func(src *%[1]s, dest *%[1]sJSON) error\n\n", st.Name())
	g.Printf("// %[1]sPropertyDecoder is property decoder for [1]sJSON.\n", st.Name())
	g.Printf("type %[1]sPropertyDecoder func(src *%[1]sJSON, dest *%[1]s) error\n\n", st.Name())

	// generate property info
	g.Printf(`
			// %[1]sPropertyInfo stores property information.
			type %[1]sPropertyInfo struct {
				name	string
				Encoder %[1]sPropertyEncoder
				Decoder %[1]sPropertyDecoder
			}

			`, st.Name())

	// generate json builder
	g.Printf("// %[1]sJSONBuilder convert between %[1]s to %[1]sJSON mutually.\n", st.Name())
	g.Printf("type %sJSONBuilder struct {\n", st.Name())
	g.Printf("_properties map[string]*%sPropertyInfo\n", st.Name())
	for _, field := range st.Fields {
		if field.Tag.Ignore {
			continue
		}
		g.Printf("%s *%sPropertyInfo\n", field.Name, st.Name())
	}
	g.Printf("}\n")

	// generate new json builder factory function
	g.Printf("// New%[1]sJSONBuilder make new %[1]sJSONBuilder.\n", st.Name())
	g.Printf("func New%[1]sJSONBuilder() *%[1]sJSONBuilder {\n", st.Name())
	g.Printf("return &%sJSONBuilder{\n", st.Name())
	g.Printf("_properties: map[string]*%sPropertyInfo{},\n", st.Name())
	for _, field := range st.Fields {
		if field.Tag.Ignore {
			continue
		}
		if field.Embed {
			if field.WithJWG() && field.IsPtr() {
				typeName, err := genbase.ExprToBaseTypeName(field.fieldInfo.Type)
				if err != nil {
					return err
				}
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							} else if src.%[2]s == nil {
								return nil
							}
							d, err := New%[3]sJSONBuilder().AddAll().Convert(src.%[2]s)
							if err != nil {
								return err
							}
							dest.%[2]sJSON = d
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							} else if src.%[2]sJSON == nil {
								return nil
							}
							d, err := src.%[2]sJSON.Convert()
							if err != nil {
								return err
							}
							dest.%[2]s = d
							return nil
						},
					},
					`, st.Name(), field.Name, typeName)
			} else if field.WithJWG() {
				typeName, err := genbase.ExprToBaseTypeName(field.fieldInfo.Type)
				if err != nil {
					return err
				}
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							}
							d, err := New%[3]sJSONBuilder().AddAll().Convert(&src.%[2]s)
							if err != nil {
								return err
							}
							dest.%[2]sJSON = *d
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							}
							d, err := src.%[2]sJSON.Convert()
							if err != nil {
								return err
							}
							dest.%[2]s = *d
							return nil
						},
					},
					`, st.Name(), field.Name, typeName)
			} else {
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							}
							dest.%[2]s = src.%[2]s
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							}
							dest.%[2]s = src.%[2]s
							return nil
						},
					},
					`, st.Name(), field.Name)
			}
		} else {
			if field.WithJWG() && field.IsPtrArrayPtr() {
				typeName, err := genbase.ExprToBaseTypeName(field.fieldInfo.Type)
				if err != nil {
					return err
				}
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							} else if src.%[2]s == nil {
								return nil
							}
							list, err := New%[3]sJSONBuilder().AddAll().ConvertList(*src.%[2]s)
							if err != nil {
								return err
							}
							dest.%[2]s = (*[]*%[3]sJSON)(&list)
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							} else if src.%[2]s == nil {
								return nil
							}
							list := make([]*%[3]s, len(*src.%[2]s))
							for idx, obj := range *src.%[2]s {
								if obj == nil {
									continue
								}
								d, err := obj.Convert()
								if err != nil {
									return err
								}
								list[idx] = d
							}
							dest.%[2]s = &list
							return nil
						},
					},
					`, st.Name(), field.Name, typeName)
			} else if field.WithJWG() && field.IsPtrArray() {
				typeName, err := genbase.ExprToBaseTypeName(field.fieldInfo.Type)
				if err != nil {
					return err
				}
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							} else if src.%[2]s == nil {
								return nil
							}
							b := New%[3]sJSONBuilder().AddAll()
							list := make([]%[3]sJSON, len(*src.%[2]s))
							for idx, obj := range *src.%[2]s {
								d, err := b.Convert(&obj)
								if err != nil {
									return err
								}
								list[idx] = *d
							}
							dest.%[2]s = &list
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							} else if src.%[2]s == nil {
								return nil
							}
							list := make([]%[3]s, len(*src.%[2]s))
							for idx, obj := range *src.%[2]s {
								d, err := obj.Convert()
								if err != nil {
									return err
								}
								list[idx] = *d
							}
							dest.%[2]s = &list
							return nil
						},
					},
					`, st.Name(), field.Name, typeName)
			} else if field.WithJWG() && field.IsArrayPtr() {
				typeName, err := genbase.ExprToBaseTypeName(field.fieldInfo.Type)
				if err != nil {
					return err
				}
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							}
							list, err := New%[3]sJSONBuilder().AddAll().ConvertList(src.%[2]s)
							if err != nil {
								return err
							}
							dest.%[2]s = ([]*%[3]sJSON)(list)
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							}
							list := make([]*%[3]s, len(src.%[2]s))
							for idx, obj := range src.%[2]s {
								if obj == nil {
									continue
								}
								d, err := obj.Convert()
								if err != nil {
									return err
								}
								list[idx] = d
							}
							dest.%[2]s = list
							return nil
						},
					},
					`, st.Name(), field.Name, typeName)
			} else if field.WithJWG() && field.IsArray() {
				typeName, err := genbase.ExprToBaseTypeName(field.fieldInfo.Type)
				if err != nil {
					return err
				}
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							}
							b := New%[3]sJSONBuilder().AddAll()
							list := make([]%[3]sJSON, len(src.%[2]s))
							for idx, obj := range src.%[2]s {
								d, err := b.Convert(&obj)
								if err != nil {
									return err
								}
								list[idx] = *d
							}
							dest.%[2]s = list
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							}
							list := make([]%[3]s, len(src.%[2]s))
							for idx, obj := range src.%[2]s {
								d, err := obj.Convert()
								if err != nil {
									return err
								}
								list[idx] = *d
							}
							dest.%[2]s = list
							return nil
						},
					},
					`, st.Name(), field.Name, typeName)
			} else if field.WithJWG() && field.IsPtr() {
				typeName, err := genbase.ExprToBaseTypeName(field.fieldInfo.Type)
				if err != nil {
					return err
				}
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							} else if src.%[2]s == nil {
								return nil
							}
							d, err := New%[3]sJSONBuilder().AddAll().Convert(src.%[2]s)
							if err != nil {
								return err
							}
							dest.%[2]s = d
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							} else if src.%[2]s == nil {
								return nil
							}
							d, err := src.%[2]s.Convert()
							if err != nil {
								return err
							}
							dest.%[2]s = d
							return nil
						},
					},
					`, st.Name(), field.Name, typeName)
			} else if field.WithJWG() {
				typeName, err := genbase.ExprToBaseTypeName(field.fieldInfo.Type)
				if err != nil {
					return err
				}
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							}
							d, err := New%[3]sJSONBuilder().AddAll().Convert(&src.%[2]s)
							if err != nil {
								return err
							}
							dest.%[2]s = *d
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							}
							d, err := src.%[2]s.Convert()
							if err != nil {
								return err
							}
							dest.%[2]s = *d
							return nil
						},
					},
					`, st.Name(), field.Name, typeName)
			} else {
				g.Printf(`%[2]s: &%[1]sPropertyInfo{
						name: "%[2]s",
						Encoder: func(src *%[1]s, dest *%[1]sJSON) error {
							if src == nil {
								return nil
							}
							dest.%[2]s = src.%[2]s
							return nil
						},
						Decoder: func(src *%[1]sJSON, dest *%[1]s) error {
							if src == nil {
								return nil
							}
							dest.%[2]s = src.%[2]s
							return nil
						},
					},
					`, st.Name(), field.Name)
			}
		}
	}
	g.Printf("}\n")
	g.Printf("}\n\n")

	// generate AddAll method
	g.Printf("// AddAll adds all property to %[1]sJSONBuilder.\n", st.Name())
	g.Printf("func (b *%[1]sJSONBuilder) AddAll() *%[1]sJSONBuilder {\n", st.Name())
	for _, field := range st.Fields {
		if field.Tag.Ignore {
			continue
		}
		if field.Name != "" {
			g.Printf("b._properties[\"%[1]s\"] = b.%[1]s\n", field.Name)
		} else {
			// TODO add support for embed other struct
		}
	}
	g.Printf("return b\n")
	g.Printf("}\n\n")

	// generate method for modifing and marshaling
	g.Printf(`
			// Add specified property to %[1]sJSONBuilder.
			func (b *%[1]sJSONBuilder) Add(info *%[1]sPropertyInfo) *%[1]sJSONBuilder {
				b._properties[info.name] = info
				return b
			}

			// Remove specified property to %[1]sJSONBuilder.
			func (b *%[1]sJSONBuilder) Remove(info *%[1]sPropertyInfo) *%[1]sJSONBuilder {
				delete(b._properties, info.name)
				return b
			}

			// Convert specified non-JSON object to JSON object.
			func (b *%[1]sJSONBuilder) Convert(orig *%[1]s) (*%[1]sJSON, error) {
				if orig == nil {
				  return nil, nil
				}
				ret := &%[1]sJSON{}

				for _, info := range b._properties {
					if err := info.Encoder(orig, ret); err != nil {
						return nil, err
					}
				}

				return ret, nil
			}

			// ConvertList specified non-JSON slice to JSONList.
			func (b *%[1]sJSONBuilder) ConvertList(orig []*%[1]s) (%[1]sJSONList, error) {
				if orig == nil {
					return nil, nil
				}

				list := make(%[1]sJSONList, len(orig))
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
			func (orig *%[1]sJSON) Convert() (*%[1]s, error) {
				ret := &%[1]s{}

				b := New%[1]sJSONBuilder().AddAll()
				for _, info := range b._properties {
					if err := info.Decoder(orig, ret); err != nil {
						return nil, err
					}
				}

				return ret, nil
			}

			// Convert specified JSONList to non-JSON slice.
			func (jsonList %[1]sJSONList) Convert() ([]*%[1]s, error) {
				orig := ([]*%[1]sJSON)(jsonList)

				list := make([]*%[1]s, len(orig))
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
			func (b *%[1]sJSONBuilder) Marshal(orig *%[1]s) ([]byte, error) {
				ret, err :=  b.Convert(orig)
				if err != nil {
					return nil, err
				}
				return json.Marshal(ret)
			}
		`, st.Name())

	g.Printf("\n\n")

	return nil
}

// Name returns struct type name.
func (st *BuildStruct) Name() string {
	return st.typeInfo.Name()
}

// WithJWG returns is field jwg annotated type.
func (f *BuildField) WithJWG() bool {
	fieldType, err := genbase.ExprToBaseTypeName(f.fieldInfo.Type)
	if err != nil {
		panic(err)
	}
	for _, st := range f.parent.parent.Structs {
		if fieldType == st.Name() {
			return true
		}
	}
	return false
}

// IsPtr returns field type is pointer.
func (f *BuildField) IsPtr() bool {
	return f.fieldInfo.IsPtr()
}

// IsArray returns field type is array.
func (f *BuildField) IsArray() bool {
	return f.fieldInfo.IsArray()
}

// IsPtrArray returns field type is pointer array.
func (f *BuildField) IsPtrArray() bool {
	return f.fieldInfo.IsPtrArray()
}

// IsArrayPtr returns field type is array of pointer.
func (f *BuildField) IsArrayPtr() bool {
	return f.fieldInfo.IsArrayPtr()
}

// IsPtrArrayPtr returns field type is pointer of pointer array.
func (f *BuildField) IsPtrArrayPtr() bool {
	return f.fieldInfo.IsPtrArrayPtr()
}

// TagString build tag string.
func (tag *BuildTag) TagString() string {
	result := tag.Name
	result += ",omitempty"
	if tag.String { // TODO add special support for int64
		result += ",string"
	}
	return "json:\"" + result + "\""
}
