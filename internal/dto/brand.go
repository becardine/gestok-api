package dto

import (
	"encoding/json"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type CreateBrandInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateBrandInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (input *CreateBrandInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *CreateBrandInput) ToEntity() *entity.Brand {
	return &entity.Brand{
		ID:          common.NewID(),
		Name:        input.Name,
		Description: input.Description,
	}
}
