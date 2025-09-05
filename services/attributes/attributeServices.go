package attributes

import (
	"Store-Dio/repo"
)

type AttributeService struct {
	AttributeRepo *repo.AttributeRepo
}

func NewAttributeService(attributeRepo *repo.AttributeRepo) *AttributeService {
	return &AttributeService{
		AttributeRepo: attributeRepo,
	}
}
