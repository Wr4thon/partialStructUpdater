package templates

// UpdateMethodTemplateName represents the name of the UpdateMethod.
const UpdateMethodTemplateName = "updateTemplate"

// UpdateMethodData contains the fields required for filling out the
// UpdateMethod template.
type UpdateMethodData struct {
	TypeName            string
	MethodReceiverName  string
	UpdateMaskFieldName string
	UpdateTypeName      string
	Fields              []FieldData
}

// FieldData contains the data for a Field.
type FieldData struct {
	Name                        string
	Type                        string
	UpdateMaskPropertyFieldName string
}

// UpdateMethod is the template for the Update method.
const UpdateMethod = `
func ({{.MethodReceiverName}} *{{.TypeName}}) Update(update {{.UpdateTypeName}}) {
	for _, value := range {{.UpdateMaskFieldName}} {
		switch value { {{range .Fields}} 
		case {{.UpdateMaskPropertyFieldName}} & update.UpdateMask:
			{{$.MethodReceiverName}}.{{.Name}} = update.{{$.TypeName}}.{{.Name}}{{end}}
		}
	}
}
`

// UpdateMethodCaseTemplate is the template for the case in the update method
// switch.
const UpdateMethodCaseTemplate = `
case PersonMaskFirstName & updateMask:
	{{.MethodReceiverName}}.FirstName = newValues.FirstName
	break
}
`
