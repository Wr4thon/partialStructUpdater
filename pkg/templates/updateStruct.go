package templates

/*
type PersonUpdate struct {
	Person     Person           `json:"person"`
	UpdateMask PersonUpdateMask `json:"updateMask"`
}
*/

// UpdateStructTemplateName represents the name of the FileHeader.
const UpdateStructTemplateName = "updateStruct"

// UpdateStructData contains the fields required for filling out the
// UpdateStruct template.
type UpdateStructData struct {
	Typename            string
	UpdateMaskFieldName string
	UpdateTypeName      string
}

// UpdateStruct is the template for the file outline.
const UpdateStruct = `
type {{.UpdateTypeName}} struct {
	{{.Typename}} {{.Typename}}
	UpdateMask {{.UpdateMaskFieldName}}
}
`
