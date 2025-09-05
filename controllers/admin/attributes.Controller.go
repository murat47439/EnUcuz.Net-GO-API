package admin

import (
	"Store-Dio/services/attributes"
)

type AttributeController struct {
	AttributeService *attributes.AttributeService
}

func NewAttributeController(attributeService *attributes.AttributeService) *AttributeController {
	return &AttributeController{
		AttributeService: attributeService,
	}
}
