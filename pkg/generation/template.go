package generation

import (
	"io"
	"text/template"

	"github.com/Wr4thon/partialStructUpdater/pkg/loader"
	"github.com/Wr4thon/partialStructUpdater/pkg/templates"
	"github.com/pkg/errors"
)

// Generator is the interface to generate an partial struct updater.
type Generator interface {
	GenerateFile(typeMeta loader.TypeMeta, writer io.Writer) error
}

type templateGenerator struct {
	templates map[string]*template.Template
}

func appendTemplate(append map[string]*template.Template, templateName string, content string) (err error) {
	var t *template.Template

	if t, err = template.New(templateName).Parse(content); err != nil {
		err = errors.Wrapf(err, "error while parsing template %v", templateName)
	} else {
		append[templateName] = t
	}

	return
}

// NewTemplateGenerator returns a new instance of the templateNameGenerator struct.
func NewTemplateGenerator() (Generator, error) {
	templateMap := make(map[string]*template.Template)

	for key, template := range templates.TemplateMapping {
		if err := appendTemplate(templateMap, key, template); err != nil {
			return nil, errors.Wrap(err, "error wile appending")
		}
	}

	return &templateGenerator{
		templates: templateMap,
	}, nil
}

func (tg *templateGenerator) GenerateFile(typeMeta loader.TypeMeta, writer io.Writer) error {
	if err := tg.generate(writer, tg.generateFileHeader(typeMeta)); err != nil {
		return errors.Wrap(err, "error while generating the fileHeader")
	}

	if err := tg.generate(writer, tg.generateUpdateMask(typeMeta)); err != nil {
		return errors.Wrap(err, "error while generating the updateMask")
	}

	if err := tg.generate(writer, tg.generateUpdateMaskConstsData(typeMeta)); err != nil {
		return errors.Wrap(err, "error while generating the updateMaskConstsData")
	}

	if err := tg.generate(writer, tg.generateUpdateStruct(typeMeta)); err != nil {
		return errors.Wrap(err, "error while generating the updateStruct")
	}

	if err := tg.generate(writer, tg.generateUpdateMaskSliceData(typeMeta)); err != nil {
		return errors.Wrap(err, "error while generating the updateMaskSliceData")
	}

	if err := tg.generate(writer, tg.generateUpdateMethod(typeMeta)); err != nil {
		return errors.Wrap(err, "error while generating the update method")
	}

	return nil
}

func (tg *templateGenerator) generate(writer io.Writer, dataGenerator func() (string, interface{})) (err error) {
	var t *template.Template

	var ok bool

	templateName, data := dataGenerator()
	if t, ok = tg.templates[templateName]; !ok {
		err = errors.Errorf("template with name %v not found", templateName)
		return
	}

	// TODO: template.Lookup(templateName) -> no map nonsense

	if err = t.Execute(writer, data); err != nil {
		err = errors.Wrap(err, "error while executing the template")
	}

	return
}

func (tg *templateGenerator) generateUpdateMask(typeMeta loader.TypeMeta) func() (string, interface{}) {
	return func() (string, interface{}) {
		return templates.UpdateMaskTemplateName,
			templates.UpdateMaskData{
				TypeName: typeMeta.TypeName(),
			}
	}
}

func (tg *templateGenerator) generateUpdateMethod(typeMeta loader.TypeMeta) func() (string, interface{}) {
	return func() (string, interface{}) {
		return templates.UpdateMethodTemplateName,
			templates.UpdateMethodData{
				MethodReceiverName:  typeMeta.MethodReceiverName(),
				TypeName:            typeMeta.TypeName(),
				Fields:              tg.getFields(typeMeta),
				UpdateMaskFieldName: typeMeta.UpdateMaskFieldName(),
				UpdateTypeName:      typeMeta.UpdateTypeName(),
			}
	}
}

func (tg *templateGenerator) getFields(typeMeta loader.TypeMeta) []templates.FieldData {
	fields := make([]templates.FieldData, len(typeMeta.Fields()))
	for i, fieldMeta := range typeMeta.Fields() {
		fields[i] = templates.FieldData{
			Name:                        fieldMeta.Name(),
			Type:                        fieldMeta.Type(),
			UpdateMaskPropertyFieldName: fieldMeta.UpdateMaskFieldName(),
		}
	}

	return fields
}

func (tg *templateGenerator) generateUpdateMaskSliceData(typeMeta loader.TypeMeta) func() (string, interface{}) {
	return func() (string, interface{}) {
		return templates.UpdateMaskSliceName,
			templates.UpdateMaskSliceData{
				Fields:              tg.getFields(typeMeta),
				UpdateMaskFieldName: typeMeta.UpdateMaskFieldName(),
				UpdateMaskName:      typeMeta.UpdateMaskName(),
			}
	}
}

func (tg *templateGenerator) generateUpdateMaskConstsData(typeMeta loader.TypeMeta) func() (string, interface{}) {
	return func() (string, interface{}) {
		return templates.UpdateMaskConstsName,
			templates.UpdateMaskConstsData{
				Fields:         tg.getFields(typeMeta),
				UpdateMaskName: typeMeta.UpdateMaskName(),
			}
	}
}

func (tg *templateGenerator) generateFileHeader(typeMeta loader.TypeMeta) func() (string, interface{}) {
	return func() (string, interface{}) {
		return templates.FileHeaderTemplateName,
			templates.FileHeaderData{
				PackageName: typeMeta.PackageName(),
			}
	}
}

func (tg *templateGenerator) generateUpdateStruct(typeMeta loader.TypeMeta) func() (string, interface{}) {
	return func() (string, interface{}) {
		return templates.UpdateStructTemplateName,
			templates.UpdateStructData{
				Typename:            typeMeta.TypeName(),
				UpdateMaskFieldName: typeMeta.UpdateMaskName(),
				UpdateTypeName:      typeMeta.UpdateTypeName(),
			}
	}
}
