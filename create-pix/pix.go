package pix

import (
	"encoding/json"
	"errors"
	"fmt"

	//	"os"
	"time"

	//	"github.com/gothello/go-pix-mercado-pago/create-pix/service"

	"github.com/gothello/go-pix-mercado-pago/request"
)

var (
	BASE_URL   = "https://api.mercadopago.com"
	SECRET_KEY = "APP_USR-6812762136376103-020807-d1f289344c1a03ccb01f6c75801acd7a-811772071"
)

func (p *InputPix) CreatePix() (*OutputPix, error) {

	expiration := time.Now().Add(p.TimeOfExpiration).Format("2006-01-02T15:04:05.000-07:00")

	headers := map[string]string{
		"accept":        "application/json",
		"content-type":  "application/json",
		"Authorization": `Bearer ` + SECRET_KEY,
	}

	body := fmt.Sprintf(`{
		"transaction_amount": %v,
		"description": "%s",
		"payment_method_id": "pix",
		"payer": {
		  "email": "%s",
		  "first_name": "Adam Dev",
		  "last_name": "GOPIX API",
		  "identification": {
			"type": "CPF",
			"number": "01234567890"
		  }
		},
		"date_of_expiration": "%s",
		"notification_url": "%s"
	}`, p.Amount, p.Description, p.Email, expiration, p.UrlNotify)

	opt := request.NewOptions("POST", BASE_URL+"/v1/payments", body, 0, headers)

	r := opt.Request()
	if r.Err != nil {
		return nil, r.Err
	}

	//	fmt.Printf("%#v", string(r.Body))

	var dt ResponseMP

	if err := json.Unmarshal(r.Body, &dt); err != nil {
		return nil, err
	}

	return &OutputPix{
		ID:                    p.ID,
		IDExternalTransaction: dt.ID,
		CreateAt:              dt.DateCreated,
		ExpiresAt:             dt.DateOfExpiration,
		Status:                dt.Status,
		Type:                  dt.PaymentMethod.ID,
		Amount:                p.Amount,
		Ticket:                dt.PointOfInteraction.TransactionData.TicketURL,
		QrCode:                dt.PointOfInteraction.TransactionData.QrCode,
		QrCodeBase64:          dt.PointOfInteraction.TransactionData.QrCodeBase64,
	}, nil
}

func (p *OutputPix) CancelPix() error {

	h := map[string]string{
		"Authorization": "Bearer " + SECRET_KEY,
		"Content-Type":  "application/json",
	}

	body := `{"status": "cancelled"}`

	end := fmt.Sprintf(BASE_URL+"/v1/payments/%v", p.IDExternalTransaction)

	opt := request.NewOptions("PUT", end, body, 0, h)

	resp := opt.Request()
	if resp.Err != nil {
		return resp.Err
	}

	var rm ResponseMP

	if err := json.Unmarshal(resp.Body, &rm); err != nil {
		return err
	}

	if rm.Status == "cancelled" {
		p.Status = rm.Status
		return nil
	}

	return errors.New("error in cancel transaction")
}

func (p *OutputPix) RefundPix() error {

	h := map[string]string{
		"Authorization": "Bearer " + SECRET_KEY,
		"Content-Type":  "application/json",
	}

	//	body := `{"amount": 1}`

	end := fmt.Sprintf(BASE_URL+"/v1/payments/%v/refunds", p.IDExternalTransaction)

	opt := request.NewOptions("POST", end, "", 0, h)

	resp := opt.Request()
	if resp.Err != nil {
		return resp.Err
	}

	var rr RefundData

	if err := json.Unmarshal(resp.Body, &rr); err != nil {
		return err
	}

	if rr.Status == "approved" {
		p.Status = "payment refund"
		return nil
	}

	return errors.New("error in refund transaction")
}
