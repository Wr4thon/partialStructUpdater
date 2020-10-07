package loader

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

// TypeLoader is the interface used to extract types from a file.
type TypeLoader interface {
	LoadByTypename(typeName string) (TypeMeta, error)
}

type binaryTypeloader struct {
}

// NewBinaryTypeloader creates a new Instance of the binaryTypeLoader
// which loads the typeinformation directly from the current binary.
func NewBinaryTypeloader() TypeLoader {
	return &binaryTypeloader{}
}

func (btl *binaryTypeloader) LoadByTypename(typeName string) (TypeMeta, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedExportsFile |
			packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo |
			packages.NeedTypesSizes |
			packages.NeedModule,

		Dir: "/home/johannes/go/src/github.com/Wr4thon/partialStructUpdater/testData",
	}

	pkgs, err := packages.Load(cfg)

	if err != nil {
		panic(err)
	}

	if len(pkgs) != 1 {
		panic(fmt.Errorf("got unexpected number of packages %v", len(pkgs)))
	}

	pkg := pkgs[0]

	targetType := pkg.Types.Scope().Lookup(typeName)
	if targetType == nil {
		return nil, fmt.Errorf("failed to find typedeclaration for %v", typeName)
	}

	targetMeta := &typeMeta{
		targetType:         targetType,
		typeName:           typeName,
		methodReceiverName: strings.ToLower(string(typeName[0])),
		packageName:        btl.getPackageName(targetType),
	}

	fields, err := btl.getFields(targetMeta)
	if err != nil {
		return nil, errors.Wrap(err, "error while trying to load fields")
	}

	targetMeta.fields = fields

	return targetMeta, nil
}

func (btl *binaryTypeloader) getPackageName(targetMeta types.Object) string {
	return targetMeta.Pkg().Name()
}

func (btl *binaryTypeloader) getFields(targetMeta *typeMeta) ([]FieldMeta, error) {
	str, ok := targetMeta.targetType.Type().Underlying().(*types.Struct)
	if !ok {
		return nil, errors.New("underlying type is not a struct")
	}

	fields := make([]FieldMeta, str.NumFields())

	for i := 0; i < len(fields); i++ {
		field := str.Field(i)

		fieldMeta := &fieldMeta{
			name:   field.Name(),
			parent: targetMeta,
		}

		switch v := field.Type().Underlying().(type) {
		case (*types.Basic):
			fieldMeta.fieldType = v.Name()
		case (*types.Struct):
			named := field.Type().(*types.Named)
			fieldMeta.fieldType = strings.Join([]string{named.Obj().Pkg().Path(), named.Obj().Name()}, ".")
		default:
			return nil, errors.Errorf("unknown type: %v", v)
		}

		fields[i] = fieldMeta
	}

	return fields, nil
}
