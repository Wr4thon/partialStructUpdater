package templates

const (
	// UpdateMaskTemplateName represents the name of the UpdateMask
	UpdateMaskTemplateName = "updateMask"
	// UpdateMaskSliceName represents the name of the UpdateMaskSlice
	UpdateMaskSliceName = "updateMaskSlice"
	// UpdateMaskConstsName represents the name of the UpdateMaskConsts
	UpdateMaskConstsName = "updateMaskConsts"
)

// UpdateMaskData contains the fields required for filling out the
// UpdateMask template.
type UpdateMaskData struct {
	TypeName string
}

// UpdateMask is the template for the UpdateMask.
const UpdateMask = `
type {{.TypeName}}UpdateMask int
`

// UpdateMaskConstsData contains the fields required for filling out the
// UpdateMaskConsts template.
type UpdateMaskConstsData struct {
	UpdateMaskName string
	Fields         []FieldData
}

// UpdateMaskConsts is the template for the UpdateMaskConsts.
const UpdateMaskConsts = `
const ({{range .Fields}}
		{{.UpdateMaskPropertyFieldName}} {{$.UpdateMaskName}} = 1 << iota{{end}}
)
`

// UpdateMaskSliceData contains the fields required for filling out the
// UpdateMaskSlice template.
type UpdateMaskSliceData struct {
	UpdateMaskFieldName string
	UpdateMaskName      string
	Fields              []FieldData
}

// UpdateMaskSlice is the template for the UpdateMaskSlice.
const UpdateMaskSlice = `
var {{.UpdateMaskFieldName}} []{{.UpdateMaskName}} = []{{.UpdateMaskName}}{ {{range .Fields}}
		{{.UpdateMaskPropertyFieldName}},{{end}}
}
`
