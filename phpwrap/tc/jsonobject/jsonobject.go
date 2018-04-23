package fprotostd_phpwrap_jsonobject

import (
	"github.com/RangelReale/fdep"
	"github.com/RangelReale/fproto-wrap/phpwrap"
)

const (
	TCID_JSONOBJECT fproto_phpwrap.TCID = "ad6a0d05-2ff0-42eb-b90e-dde487fd5141"
)

// Converts between google.protobuf.Timestamp and time.Time
type TypeConverterPlugin_JSONObject struct {
}

func (t *TypeConverterPlugin_JSONObject) GetTypeConverter(tp *fdep.DepType) fproto_phpwrap.TypeConverter {
	if tp.DepFile.FilePath == "github.com/RangelReale/fproto-wrap-std/jsonobject.proto" &&
		tp.DepFile.ProtoFile.PackageName == "fproto_wrap" &&
		tp.Name == "JSONObject" {
		return &TypeConverter_JSONObject{}
	}
	return nil
}

//
// Time
// Converts between google.protobuf.Timestamp and \DateTime
//

type TypeConverter_JSONObject struct {
}

func (t *TypeConverter_JSONObject) TCID() fproto_phpwrap.TCID {
	return TCID_JSONOBJECT
}

func (t *TypeConverter_JSONObject) TypeName(g *fproto_phpwrap.GeneratorFile, tntype fproto_phpwrap.TypeNameType, options uint32) string {
	return "array"
}

func (t *TypeConverter_JSONObject) IsScalar() bool {
	return false
}

func (t *TypeConverter_JSONObject) GenerateImport(g *fproto_phpwrap.GeneratorFile, varSrc string, varDest string, varError string) (generated bool, err error) {
	g.P(varDest, " = json_decode(", varSrc, "->getValue(), TRUE);")
	return true, nil
}

func (t *TypeConverter_JSONObject) GenerateExport(g *fproto_phpwrap.GeneratorFile, varSrc string, varDest string, varError string) (generated bool, err error) {
	g.P(varDest, " = new \\Fproto_wrap\\JSONObject();")
	g.P(varDest, "->setValue(json_encode(", varSrc, "));")

	return true, nil
}
