package usecase

import (
	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/service"
)

type FindAllPixUseCase struct {
	PixRepositoryUseCase service.PixRepositoryUseCase
}

func NewFindAllPixUseCase(s service.PixRepositoryUseCase) *FindAllPixUseCase {
	return &FindAllPixUseCase{
		PixRepositoryUseCase: s,
	}
}

func (s FindAllPixUseCase) Execute() ([]*pix.OutputPix, error) {
	outputs, err := s.PixRepositoryUseCase.GetAll()
	if err != nil {
		return nil, err
	}

	return outputs, nil
}
