package usecase

import (
	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/service"
)

type UpdatePixUseCase struct {
	ServiceUseCase service.PixRepositoryUseCase
}

func NewUpdatePixUseCase(s service.PixRepositoryUseCase) *CreatePixUseCase {
	return &CreatePixUseCase{
		ServiceUseCase: s,
	}
}

func (u UpdatePixUseCase) Execute(p *pix.OutputPix) error {
	if err := u.ServiceUseCase.Update(p); err != nil {
		return err
	}

	return nil
}
