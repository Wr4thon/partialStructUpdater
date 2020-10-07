package templates

// TemplateMapping is a mapping from the templateKey to a template.
var TemplateMapping map[string]string = map[string]string{
	FileHeaderTemplateName:   FileHeader,
	UpdateMaskTemplateName:   UpdateMask,
	UpdateMethodTemplateName: UpdateMethod,
	UpdateMaskSliceName:      UpdateMaskSlice,
	UpdateMaskConstsName:     UpdateMaskConsts,
	UpdateStructTemplateName: UpdateStruct,
}
