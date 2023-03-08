package usecase

import (
	"log"
	"strconv"

	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/rabbit"
	"github.com/rabbitmq/amqp091-go"
)

var (
	NETIFY_CONNECTION *amqp091.Connection
	NETIFY_QUEUE      = "approved"
)

type RabbitApprovedOutput struct {
	ID        string  `json:"id"`
	IDPAY     int64   `json:"id_pay"`
	CreateAt  string  `json:"created_at"`
	ExpiresAt string  `json:"expires_at"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
	Email     string  `json:"email"`
}

type NotifyPixUseCase struct {
	PixRepositoryUseCase entity.PixRepository
}

func NewNotifyPixUseCase(entity entity.PixRepository) *NotifyPixUseCase {
	return &NotifyPixUseCase{
		PixRepositoryUseCase: entity,
	}
}

func (n *NotifyPixUseCase) Execute(idpay string) error {
	parse, err := strconv.ParseInt(idpay, 10, 64)
	if err != nil {
		return err
	}

	out, err := n.PixRepositoryUseCase.GetByIdPayment(parse)
	if err != nil {
		return err
	}

	if err := out.UpdatePaymentStatus(); err != nil {

		out.Status = err.Error()
		if err := n.PixRepositoryUseCase.Update(out); err != nil {
			return err
		}

		err = rabbit.Publish(
			NETIFY_CONNECTION,
			NETIFY_QUEUE,
			&RabbitApprovedOutput{
				ID:        out.ID,
				IDPAY:     out.IDExternalTransaction,
				CreateAt:  out.CreateAt,
				ExpiresAt: out.ExpiresAt,
				Status:    out.Status,
				Amount:    out.Amount,
				Email:     out.Email,
			},
		)

		if err != nil {
			log.Println(err)
		}

		return nil
	}

	out.Status = "approved"

	err = rabbit.Publish(
		NETIFY_CONNECTION,
		NETIFY_QUEUE,
		&RabbitApprovedOutput{
			ID:        out.ID,
			IDPAY:     out.IDExternalTransaction,
			CreateAt:  out.CreateAt,
			ExpiresAt: out.ExpiresAt,
			Status:    out.Status,
			Amount:    out.Amount,
			Email:     out.Email,
		},
	)

	if err != nil {
		log.Println(err)
	}

	if err := n.PixRepositoryUseCase.Update(out); err != nil {
		return err
	}

	return nil
}
