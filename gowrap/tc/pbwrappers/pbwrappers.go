package fprotostd_gowrap_pbwrappers

import (
	"fmt"

	"github.com/RangelReale/fdep"
	"github.com/RangelReale/fproto-wrap/gowrap"
)

const (
	TCID_PBWRAPPERS fproto_gowrap.TCID = "20411a81-907d-4d23-b9eb-fd1707e86cb7"
)

// Converts between google.protobuf.XXXValue and a struct from the ptypes subpackage
type TypeConverterPlugin_PBWrappers struct {
}

func (t *TypeConverterPlugin_PBWrappers) GetTypeConverter(tp *fdep.DepType) fproto_gowrap.TypeConverter {
	if tp.DepFile.FilePath == "google/protobuf/wrappers.proto" &&
		tp.DepFile.ProtoFile.PackageName == "google.protobuf" {
		return &TypeConverter_PBWrappers{
			SourceTypeName: tp.Name,
		}
	}
	return nil
}

// Converter
type TypeConverter_PBWrappers struct {
	SourceTypeName string
}

func (t *TypeConverter_PBWrappers) TCID() fproto_gowrap.TCID {
	return TCID_PBWRAPPERS
}

func (t *TypeConverter_PBWrappers) TypeName(g *fproto_gowrap.GeneratorFile, tntype fproto_gowrap.TypeNameType, options uint32) string {
	alias := g.DeclDep("github.com/RangelReale/fproto-wrap-std/gowrap/tc/pbwrappers/ptypes", "pbwrappers_ptypes")
	return fmt.Sprintf("%s.%s", alias, t.SourceTypeName)
}

func (t *TypeConverter_PBWrappers) IsPointer() bool {
	return true
}

func (t *TypeConverter_PBWrappers) GenerateImport(g *fproto_gowrap.GeneratorFile, varSrc string, varDest string, varError string) (checkError bool, err error) {
	alias := g.DeclDep("github.com/RangelReale/fproto-wrap-std/gowrap/tc/pbwrappers/ptypes", "pbwrappers_ptypes")

	g.P(varDest, " = ", alias, ".", t.SourceTypeName, "{}")
	g.P("if ", varSrc, " != nil {")
	g.In()

	g.P(varDest, ".Valid = true")
	g.P(varDest, ".WValue = ", varSrc, ".Value")

	g.Out()
	g.P("}")

	return true, nil
}

func (t *TypeConverter_PBWrappers) GenerateExport(g *fproto_gowrap.GeneratorFile, varSrc string, varDest string, varError string) (checkError bool, err error) {
	pb_alias := g.DeclDep("github.com/golang/protobuf/ptypes/wrappers", "pb_types")

	g.P("if ", varSrc, ".Valid {")
	g.In()

	g.P(varDest, " = &", pb_alias, ".", t.SourceTypeName, "{}")
	g.P(varDest, ".Value = ", varSrc, ".WValue")

	g.Out()
	g.P("}")

	return false, nil
}
