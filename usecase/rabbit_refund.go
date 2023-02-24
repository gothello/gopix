package usecase

import (
	"encoding/json"
	"log"

	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/rabbit"
)

type RabbitInputRefund struct {
	IDPAY int64 `json:"id_pay"`
}

type RabbitOutputRefund struct {
	ID     string
	IDPAY  int64
	Status string
	Amount float64
	Email  string
	Error  error
}

func (r *RabbitConnectionUseCase) RabbitRefundPixUseCase(input chan rabbit.RabbitInputChan, repository entity.PixRepository, queues map[string]string) {
	for i := range input {
		if i.Error != nil {
			log.Println(i.Error)
			err := rabbit.Publish(
				r.Conn,
				queues["REFUNDED"],
				&RabbitOutputRefund{Error: i.Error},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err.Error())
			}

			continue
		}

		var irefund RabbitInputRefund

		if err := json.Unmarshal(i.Delivery.Body, &irefund); err != nil {
			log.Println(err)
			err := rabbit.Publish(
				r.Conn,
				queues["REFUNDED"],
				&RabbitOutputRefund{Error: err},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err)
			}

			continue
		}

		o, err := r.PixRepositoryUseCase.GetByIdPayment(irefund.IDPAY)
		if err != nil {
			log.Println(err)
			err := rabbit.Publish(
				r.Conn,
				queues["REFUNDED"],
				&RabbitOutputRefund{Error: err},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err)
			}

			continue
		}

		if err := o.RefundPix(); err != nil {
			log.Println(err)
			err := rabbit.Publish(
				r.Conn,
				queues["REFUNDED"],
				&RabbitOutputRefund{Error: err},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err)
			}

			continue
		}

		if err := r.PixRepositoryUseCase.Update(o); err != nil {
			log.Println(err)
			err := rabbit.Publish(
				r.Conn,
				queues["REFUNDED"],
				&RabbitOutputRefund{Error: err},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err)
			}

			continue
		}

		err = rabbit.Publish(
			r.Conn,
			queues["REFUNDED"],
			&RabbitOutputRefund{
				ID:     o.ID,
				IDPAY:  o.IDExternalTransaction,
				Status: o.Status,
				Amount: o.Amount,
				Email:  o.Email,
				Error:  nil,
			},
		)

		if err != nil {
			log.Printf("erro to publish output transcation in queue: %s\n", err.Error())

		}
	}
}
