package templates

// FileHeaderTemplateName represents the name of the FileHeader.
const FileHeaderTemplateName = "fileHeader"

// FileHeaderData contains the fields required for filling out the
// FileHeader template.
type FileHeaderData struct {
	PackageName string
}

// FileHeader is the template for the file outline.
const FileHeader = `package {{.PackageName}}
// generated code
`
