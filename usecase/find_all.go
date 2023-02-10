package usecase

import (
	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/service"
)

type FindAllPixUseCase struct {
	ServiceUseCase service.PixRepositoryUseCase
}

func NewFindAllPixUseCase(s service.PixRepositoryUseCase) *FindAllPixUseCase {
	return &FindAllPixUseCase{
		ServiceUseCase: s,
	}
}

func (s FindAllPixUseCase) Execute() ([]*pix.OutputPix, error) {
	outputs, err := s.ServiceUseCase.GetAll()
	if err != nil {
		return nil, err
	}

	return outputs, nil
}
