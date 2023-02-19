package usecase

import (
	"encoding/json"
	"log"
	"time"

	// "github.com/gothello/go-pix-mercado-pago/rabbit"~
	//	"github.com/gothello/go-pix-mercado-pago/rabbit"

	//"github.com/gothello/go-pix-mercado-pago/rabbit"
	"github.com/google/uuid"
	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/pix"
	"github.com/gothello/go-pix-mercado-pago/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitOutputPix struct {
	ID           string  `json:"id"`
	IDPAY        int64   `json:"id_pay"`
	CreateAt     string  `json:"created_at"`
	ExpiresAt    string  `json:"expires_at"`
	Status       string  `json:"status"`
	Type         string  `json:"type"`
	Amount       float64 `json:"amount"`
	Email        string  `json:"email"`
	QrCode       string  `json:"qrcode"`
	QrCodeBase64 string  `json:"qrcodebase"`
}

type RabbitInputChan struct {
	Delivery amqp.Delivery
	Error    error
}

type RabbitConnectionUseCase struct {
	PixRepositoryUseCase entity.PixRepository
	Conn                 *amqp.Connection
}

func NewRabbitConnectionUseCase(c *amqp.Connection, s entity.PixRepository) *RabbitConnectionUseCase {
	return &RabbitConnectionUseCase{
		PixRepositoryUseCase: s,
		Conn:                 c,
	}
}

func (r *RabbitConnectionUseCase) RabbitCreatePixUseCase(in chan rabbit.RabbitInputChan, entity entity.PixRepository, queues map[string]string) {

	for i := range in {
		if i.Error != nil {
			log.Println(i.Error)
			continue
		}

		var input InputPix

		err := json.Unmarshal(i.Delivery.Body, &input)
		if err != nil {
			log.Println(err)
			continue
		}

		if input.Email == "" {
			log.Println("email not informed")
			continue
		}

		if input.ExpirationMinutes < 10 {
			input.ExpirationMinutes = 10
		}

		i := &pix.InputPix{
			ID:               uuid.New().String(),
			Amount:           input.Amount,
			Description:      input.Description,
			TimeOfExpiration: time.Duration(input.ExpirationMinutes) * time.Minute,
			UrlNotify:        input.UrlNotify,
			Email:            input.Email,
		}

		output, err := i.CreatePix()
		if err != nil {
			log.Println(err)
			continue
		}

		//save in database sql
		if err := r.PixRepositoryUseCase.Insert(output); err != nil {
			log.Println(err)
		}

		go func() error {
			if err := output.GetStatusPayment(20); err != nil {
				if err.Error() == "client not pay" {
					output.Status = "cancelled"
					if err := r.PixRepositoryUseCase.Update(output); err != nil {
						return err
					}

					if err := output.CancelPix(); err != nil {
						return err
					}

					return nil
				}

				if err.Error() == "approved" {
					output.Status = "approved"
					if err := r.PixRepositoryUseCase.Update(output); err != nil {
						return err
					}

					go func() {
						if err := rabbit.Publish(r.Conn, queues["APPROVED"], output); err != nil {
							log.Println(err)
						}
					}()

					return nil
				}

				return err
			}

			return nil
		}()

		err = rabbit.Publish(r.Conn, queues["CREATED"], &RabbitOutputPix{
			ID:           output.ID,
			IDPAY:        output.IDExternalTransaction,
			CreateAt:     output.CreateAt,
			ExpiresAt:    output.ExpiresAt,
			Status:       output.Status,
			Type:         output.Type,
			Amount:       output.Amount,
			Email:        input.Email,
			QrCode:       output.QrCode,
			QrCodeBase64: output.QrCodeBase64,
		})

		if err != nil {
			log.Fatalln(err)
		}
	}
}
