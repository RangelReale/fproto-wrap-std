package fprotostd_phpwrap_uuid

import (
	"github.com/RangelReale/fdep"
	"github.com/RangelReale/fproto-wrap/phpwrap"
)

const (
	TCID_UUID     fproto_phpwrap.TCID = "5f381deb-11a8-4ab7-ae80-7501c1dabd95"
	TCID_NULLUUID fproto_phpwrap.TCID = "d0b80892-1684-45a2-a30a-0e794c51a42a"
)

//
// UUID
// Converts between fproto_wrap.UUID and \Ramsey\Uuid\Uuid
//

type TypeConverterPlugin_UUID struct {
}

func (t *TypeConverterPlugin_UUID) GetTypeConverter(tp *fdep.DepType) fproto_phpwrap.TypeConverter {
	if tp.DepFile.FilePath == "github.com/RangelReale/fproto-wrap/uuid.proto" &&
		tp.DepFile.ProtoFile.PackageName == "fproto_wrap" &&
		tp.Name == "UUID" {
		return &TypeConverter_UUID{}
	}
	if tp.DepFile.FilePath == "github.com/RangelReale/fproto-wrap/uuid.proto" &&
		tp.DepFile.ProtoFile.PackageName == "fproto_wrap" &&
		tp.Name == "NullUUID" {
		return &TypeConverter_NullUUID{}
	}
	return nil
}

//
// UUID
// Converts between fproto_wrap.UUID and \Ramsey\Uuid\Uuid
//

type TypeConverter_UUID struct {
}

func (t *TypeConverter_UUID) TCID() fproto_phpwrap.TCID {
	return TCID_UUID
}

func (t *TypeConverter_UUID) TypeName(g *fproto_phpwrap.GeneratorFile, tntype fproto_phpwrap.TypeNameType) string {

	switch tntype {
	case fproto_phpwrap.TNT_NS_TYPENAME:
		return "\\Ramsey\\Uuid\\Uuid"
	}

	return "Uuid"
}

func (t *TypeConverter_UUID) IsScalar() bool {
	return false
}

func (t *TypeConverter_UUID) GenerateImport(g *fproto_phpwrap.GeneratorFile, varSrc string, varDest string, varError string) (generated bool, err error) {
	g.P(varDest, " = \\Ramsey\\Uuid\\Uuid::fromString(", varSrc, "->getValue());")

	return true, nil
}

func (t *TypeConverter_UUID) GenerateExport(g *fproto_phpwrap.GeneratorFile, varSrc string, varDest string, varError string) (generated bool, err error) {
	g.P(varDest, " = new \\Fproto_wrap\\UUID();")
	g.P(varDest, "->setValue(", varSrc, "->toString());")

	return true, nil
}

//
// NullUUID
// Converts between fproto_wrap.NullUUID and \Ramsey\Uuid\Uuid
//

type TypeConverter_NullUUID struct {
}

func (t *TypeConverter_NullUUID) TCID() fproto_phpwrap.TCID {
	return TCID_UUID
}

func (t *TypeConverter_NullUUID) TypeName(g *fproto_phpwrap.GeneratorFile, tntype fproto_phpwrap.TypeNameType) string {

	switch tntype {
	case fproto_phpwrap.TNT_NS_TYPENAME:
		return "\\Ramsey\\Uuid\\Uuid"
	}

	return "Uuid"
}

func (t *TypeConverter_NullUUID) IsScalar() bool {
	return false
}

func (t *TypeConverter_NullUUID) GenerateImport(g *fproto_phpwrap.GeneratorFile, varSrc string, varDest string, varError string) (generated bool, err error) {
	g.P(varDest, " = null;")
	g.P("if (", varSrc, "->getValue() != '') {")
	g.In()
	g.P(varDest, " = \\Ramsey\\Uuid\\Uuid::fromString(", varSrc, "->getValue());")
	g.Out()
	g.P("}")

	return true, nil
}

func (t *TypeConverter_NullUUID) GenerateExport(g *fproto_phpwrap.GeneratorFile, varSrc string, varDest string, varError string) (generated bool, err error) {
	g.P(varDest, " = new \\Fproto_wrap\\UUID();")
	g.P(varDest, "->setValue(", varSrc, "->toString());")

	return true, nil
}
