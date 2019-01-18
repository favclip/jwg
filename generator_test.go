package jwg

import (
	"io/ioutil"
	"testing"

	"github.com/favclip/genbase"
	"github.com/pmezard/go-difflib/difflib"
)

func TestGeneratorParsePackageDir(t *testing.T) {

	p := &genbase.Parser{}
	pInfo, err := p.ParsePackageDir("./misc/fixture/a")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if pInfo.Name() != "a" {
		t.Log("package name is not a, actual", pInfo.Name())
		t.Fail()
	}
	if len(pInfo.Files) != 2 {
		t.Log("package files length is not 2, actual", len(pInfo.Files))
		t.Fail()
	}
	if pInfo.Dir != "./misc/fixture/a" {
		t.Log("package dir is not ./misc/fixture/a, actual", pInfo.Dir)
		t.Fail()
	}
}

func TestGeneratorParsePackageFiles(t *testing.T) {
	p := &genbase.Parser{}
	pInfo, err := p.ParsePackageFiles([]string{"./misc/fixture/a/model.go"})
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if pInfo.Name() != "a" {
		t.Log("package name is not a, actual", pInfo.Name())
		t.Fail()
	}
	if len(pInfo.Files) != 1 {
		t.Log("package files length is not 1, actual", len(pInfo.Files))
		t.Fail()
	}
	if pInfo.Dir != "." {
		t.Log("package dir is not ., actual", pInfo.Dir)
		t.Fail()
	}
}

func TestGeneratorGenerate(t *testing.T) {

	testCase := []string{"a", "b", "c", "g", "h", "l", "n", "t"}
	for _, postFix := range testCase {
		p := &genbase.Parser{}
		pInfo, err := p.ParsePackageDir("./misc/fixture/" + postFix)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		var args []string
		var typeNames []string
		var transcriptTags []string
		var noOmitempty bool
		var noOmitemptyFieldTypes []string
		switch postFix {
		case "g":
			args = []string{"-type", "Sample,Inner", "-output", "misc/fixture/" + postFix + "/model_json.go", "misc/fixture/" + postFix}
			typeNames = []string{"Sample", "Inner"}
		case "l":
			args = []string{"-type", "Sample", "-output", "misc/fixture/" + postFix + "/model_json.go", "-transcripttag", "swagger,includes", "misc/fixture/" + postFix}
			typeNames = []string{"Sample"}
			transcriptTags = []string{"swagger", "includes"}
		case "n":
			args = []string{"-type", "Sample", "-output", "misc/fixture/" + postFix + "/model_json.go", "-noOmitempty", "misc/fixture/" + postFix}
			typeNames = []string{"Sample"}
			noOmitempty = true
		case "t":
			args = []string{"-type", "Sample", "-output", "misc/fixture/" + postFix + "/model_json.go", "-noOmitemptyFieldType=bool,int", "misc/fixture/" + postFix}
			typeNames = []string{"Sample"}
			noOmitemptyFieldTypes = []string{"bool", "int"}
		default:
			args = []string{"-type", "Sample", "-output", "misc/fixture/" + postFix + "/model_json.go", "misc/fixture/" + postFix}
			typeNames = []string{"Sample"}
		}

		bu, err := Parse(pInfo, pInfo.CollectTypeInfos(typeNames), &ParseOptions{
			TranscriptTagNames:    transcriptTags,
			NoOmitempty:           noOmitempty,
			NoOmitemptyFieldTypes: noOmitemptyFieldTypes,
		})
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		src, err := bu.Emit(&args)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		expected, err := ioutil.ReadFile("./misc/fixture/" + postFix + "/model_json.go")
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if string(src) != string(expected) {
			diff := difflib.UnifiedDiff{
				A:       difflib.SplitLines(string(expected)),
				B:       difflib.SplitLines(string(src)),
				Context: 5,
			}
			d, err := difflib.GetUnifiedDiffString(diff)
			if err != nil {
				t.Fatal(err)
			}
			t.Fatal(d)
		}
	}
}

func TestCollectTaggedTypes(t *testing.T) {
	p := &genbase.Parser{}
	pInfo, err := p.ParsePackageFiles([]string{"./misc/fixture/d/model.go"})
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	names := pInfo.CollectTaggedTypeInfos("+jwg")
	if len(names) != 1 {
		t.Log("names length is not 1, actual", len(names))
		t.Fail()
	}
	if names[0].Name() != "Sample" {
		t.Log("name[0] is not \"Sample\", actual", names[0].Name())
		t.Fail()
	}
}
