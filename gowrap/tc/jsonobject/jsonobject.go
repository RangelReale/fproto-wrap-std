package fprotostd_gowrap_jsonobject

import (
	"github.com/RangelReale/fdep"
	"github.com/RangelReale/fproto-wrap/gowrap"
)

const (
	TCID_JSONOBJECT fproto_gowrap.TCID = "ad6a0d05-2ff0-42eb-b90e-dde487fd5141"
)

//
// JSONObject
// Converts between fproto_gowrap.tc.json.JSON and map[string]interface{}
//

type TypeConverterPlugin_JSONObject struct {
}

func (t *TypeConverterPlugin_JSONObject) GetTypeConverter(tp *fdep.DepType) fproto_gowrap.TypeConverter {
	if tp.DepFile.FilePath == "github.com/RangelReale/fproto-wrap-std/jsonobject.proto" &&
		tp.DepFile.ProtoFile.PackageName == "fproto_wrap" &&
		tp.Name == "JSONObject" {
		return &TypeConverter_JSONObject{}
	}
	return nil
}

//
// JSONObject
// Converts between fproto_gowrap.tc.jsonobject.JSONObject and interface{}
//

type TypeConverter_JSONObject struct {
}

func (t *TypeConverter_JSONObject) TCID() fproto_gowrap.TCID {
	return TCID_JSONOBJECT
}

func (t *TypeConverter_JSONObject) TypeName(g *fproto_gowrap.GeneratorFile, tntype fproto_gowrap.TypeNameType) string {
	return "interface{}"
}

func (t *TypeConverter_JSONObject) IsPointer() bool {
	return true
}

func (t *TypeConverter_JSONObject) GenerateImport(g *fproto_gowrap.GeneratorFile, varSrc string, varDest string, varError string) (checkError bool, err error) {
	alias := g.DeclDep("encoding/json", "json")

	g.P("if ", varSrc, " != nil && ", varSrc, ".Value != \"\" {")
	g.In()

	g.P("jtemp := make(map[string]interface{})")
	g.P("err = ", alias, ".Unmarshal([]byte(", varSrc, ".Value), jtemp)")
	g.P("if err != nil {")
	g.In()
	g.P(varDest, " = jtemp")
	g.Out()
	g.P("}")

	g.Out()
	g.P("}")

	return true, nil
}

func (t *TypeConverter_JSONObject) GenerateExport(g *fproto_gowrap.GeneratorFile, varSrc string, varDest string, varError string) (checkError bool, err error) {
	alias := g.DeclDep("encoding/json", "json")

	tp, err := g.G().GetDep().FindType("fproto_wrap.JSONObject")
	if err != nil {
		return false, err
	}
	tsource := g.G().GetTypeSource(tp)

	/*
		tc_go, err := g.G().GetGoType("", "fproto_wrap.JSONObject")
		if err != nil {
			return false, err
		}
	*/

	g.P("if ", varSrc, " != nil {")
	g.In()

	g.P("var jtemp []byte")
	g.P(varDest, " = ", tsource.TypeName(g, fproto_gowrap.TNT_EMPTYVALUE))

	g.P("jtemp, err = ", alias, ".Marshal(", varSrc, ")")
	g.P("if err != nil {")
	g.In()
	g.P(varDest, ".Value = string(jtemp)")
	g.Out()
	g.P("}")
	g.Out()

	g.Out()
	g.P("}")

	return true, nil
}
