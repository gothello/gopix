package usecase

import (
	"encoding/json"
	"log"

	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/rabbit"
)

type RabbitInputCancel struct {
	IDPAY int64 `json:"id_pay"`
}

type RabbitOutputCancel struct {
	ID     string  `json:"id"`
	IDPAY  int64   `json:"id_pay"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
	Email  string  `json:"email"`
	Error  error   `json:"error"`
}

func (r *RabbitConnectionUseCase) RabbitCancelPixUseCase(input chan rabbit.RabbitInputChan, repository entity.PixRepository, queues map[string]string) {
	for i := range input {
		if i.Error != nil {
			log.Println(i.Error)
			err := rabbit.Publish(
				r.Conn,
				queues["CANCELLED"],
				&RabbitOutputCancel{Error: i.Error},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err.Error())
			}

			continue
		}

		var icancel RabbitInputCancel

		if err := json.Unmarshal(i.Delivery.Body, &icancel); err != nil {
			log.Println(err)
			err := rabbit.Publish(
				r.Conn,
				queues["CANCELLED"],
				&RabbitOutputCancel{Error: err},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err)
			}

			continue
		}

		o, err := r.PixRepositoryUseCase.GetByIdPayment(icancel.IDPAY)
		if err != nil {
			log.Println(err)
			err := rabbit.Publish(
				r.Conn,
				queues["CANCELLED"],
				&RabbitOutputCancel{Error: err},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err)
			}

			continue
		}

		if err := o.CancelPix(); err != nil {
			log.Println(err)
			err := rabbit.Publish(
				r.Conn,
				queues["CANCELLED"],
				&RabbitOutputCancel{Error: err},
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
				queues["CANCELLED"],
				&RabbitOutputCancel{Error: err},
			)

			if err != nil {
				log.Printf("erro to publish message error: %s\n", err)
			}

			continue
		}

		err = rabbit.Publish(
			r.Conn,
			queues["CANCELLED"],
			&RabbitOutputCancel{
				ID:     o.ID,
				IDPAY:  o.IDExternalTransaction,
				Status: o.Status,
				Amount: o.Amount,
				Email:  o.Email,
				Error:  nil,
			},
		)

		if err != nil {
			if err != nil {
				log.Printf("erro to publish message error: %s\n", err)
			}
		}

	}
}
