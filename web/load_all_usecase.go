package web

import (
	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/usecase"
)

func LoadAllUseCases(service *entity.RespositoryMySql) *PixHandlers {
	create := usecase.NewCreatePixUseCase(service)
	cancel := usecase.NewCancelUseCase(service)
	refund := usecase.NewRefundUseCase(service)
	find := usecase.NewFindPixUseCase(service)
	findall := usecase.NewFindAllPixUseCase(service)

	handlers := NewPixHandlers(create, cancel, refund, find, findall)

	return handlers
}
