package fprotostd_phpwrap_pbwrappers

import (
	"github.com/RangelReale/fdep"
	"github.com/RangelReale/fproto-wrap/phpwrap"
)

const (
	TCID_PBWRAPPERS fproto_phpwrap.TCID = "20411a81-907d-4d23-b9eb-fd1707e86cb7"
)

// Converts between google.protobuf.Timestamp and time.Time
type TypeConverterPlugin_PBWrappers struct {
}

func (t *TypeConverterPlugin_PBWrappers) GetTypeConverter(tp *fdep.DepType) fproto_phpwrap.TypeConverter {
	if tp.DepFile.FilePath == "google/protobuf/wrappers.proto" &&
		tp.DepFile.ProtoFile.PackageName == "google.protobuf" {
		return &TypeConverter_PBWrappers{
			SourceTypeName: tp.Name,
		}
	}
	return nil
}

//
// Time
// Converts between google.protobuf.Timestamp and \DateTime
//

type TypeConverter_PBWrappers struct {
	SourceTypeName string
}

func (t *TypeConverter_PBWrappers) TCID() fproto_phpwrap.TCID {
	return TCID_PBWRAPPERS
}

func (t *TypeConverter_PBWrappers) TypeName(g *fproto_phpwrap.GeneratorFile, tntype fproto_phpwrap.TypeNameType, options uint32) string {
	switch t.SourceTypeName {
	case "Int64Value", "Int32Value", "UInt64Value", "UInt32Value":
		return "int"
	case "DoubleValue", "FloatValue":
		return "float"
	case "StringValue":
		return "string"
	case "ByteValue":
		return "string"
	case "BoolValue":
		return "boolean"
	default:
		return "mixed"
	}
}

func (t *TypeConverter_PBWrappers) IsScalar() bool {
	return true
}

func (t *TypeConverter_PBWrappers) GenerateImport(g *fproto_phpwrap.GeneratorFile, varSrc string, varDest string, varError string) (generated bool, err error) {
	g.P(varDest, " = null;")
	g.P("if (", varSrc, " !== null) {")
	g.In()

	g.P(varDest, " = ", varSrc, "->getValue();")

	g.Out()
	g.P("}")

	return true, nil
}

func (t *TypeConverter_PBWrappers) GenerateExport(g *fproto_phpwrap.GeneratorFile, varSrc string, varDest string, varError string) (generated bool, err error) {
	g.P(varDest, " = null;")
	g.P("if (", varSrc, " !== '') {")
	g.In()

	g.P(varDest, ` = new \Google\Protobuf\`, t.SourceTypeName, "();")
	g.P(varDest, "->setValue(", varSrc, ");")

	g.Out()
	g.P("}")

	return true, nil
}
