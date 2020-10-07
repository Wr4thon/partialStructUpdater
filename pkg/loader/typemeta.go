package loader

import (
	"go/types"
	"strings"
)

const updateMask = "UpdateMask"

// TypeMeta contains the metadata of the type.
type TypeMeta interface {
	PackageName() string
	TypeName() string
	MethodReceiverName() string
	Fields() []FieldMeta
	UpdateMaskFieldName() string
	UpdateMaskName() string
	UpdateTypeName() string
}

type typeMeta struct {
	targetType         types.Object
	packageName        string
	typeName           string
	methodReceiverName string
	fields             []FieldMeta
}

func (tm *typeMeta) PackageName() string {
	return tm.packageName
}

func (tm *typeMeta) TypeName() string {
	return tm.typeName
}

func (tm *typeMeta) MethodReceiverName() string {
	return tm.methodReceiverName
}

func (tm *typeMeta) Fields() []FieldMeta {
	return tm.fields
}

func (tm *typeMeta) UpdateMaskFieldName() string {
	return strings.ToLower(string(tm.typeName[0])) + tm.typeName[1:] + updateMask
}

func (tm *typeMeta) UpdateMaskName() string {
	return tm.typeName + updateMask
}

func (tm *typeMeta) UpdateTypeName() string {
	return tm.typeName + "Update"
}

// FieldMeta contains the metadata of fields.
type FieldMeta interface {
	Name() string
	Type() string
	UpdateMaskFieldName() string
}

type fieldMeta struct {
	parent    *typeMeta
	name      string
	fieldType string
}

func (fm *fieldMeta) Name() string {
	return fm.name
}

func (fm *fieldMeta) Type() string {
	return fm.fieldType
}

func (fm *fieldMeta) UpdateMaskFieldName() string {
	return fm.parent.UpdateMaskName() + fm.name
}
