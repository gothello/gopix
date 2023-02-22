package pix

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	//	"os"
	"time"

	"github.com/gothello/go-pix-mercado-pago/pix/request"
)

var (
	BASE_URL      = "https://api.mercadopago.com"
	SECRET_KEY_MP = os.Getenv("SECRET_KEY_MP")
)

func init() {
	if SECRET_KEY_MP == "" {
		log.Fatalln("error key to access api Mercado Pago is empty")
	}
}

func (p *InputPix) CreatePix() (*OutputPix, error) {

	fin := "2006-01-02T15:04:05.000-07:00"
	fout := "15:04 02/01/2006"

	expiration := time.Now().Add(p.TimeOfExpiration).Format(fin)

	headers := map[string]string{
		"accept":        "application/json",
		"content-type":  "application/json",
		"Authorization": `Bearer ` + SECRET_KEY_MP,
	}

	body := fmt.Sprintf(`{
		"transaction_amount": %v,
		"description": "%s",
		"payment_method_id": "pix",
		"payer": {
		  "email": "%s",
		},
		"date_of_expiration": "%s",
		"notification_url": "%s"
	}`, p.Amount, p.Description, p.Email, expiration, p.UrlNotify)

	opt := request.NewOptions("POST", BASE_URL+"/v1/payments", body, 0, headers)

	r := opt.Request()
	if r.Err != nil {
		return nil, r.Err
	}

	//exportfmt.Printf("%#v", string(r.Body))

	var dt ResponseMP

	if err := json.Unmarshal(r.Body, &dt); err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, errors.New("error parse location")
	}

	//	fmt.Println(dt)

	out := &OutputPix{
		ID:                    p.ID,
		IDExternalTransaction: dt.ID,
		CreateAt:              time.Now().In(loc).Format(fout),
		ExpiresAt:             time.Now().In(loc).Add(p.TimeOfExpiration).Format(fout),
		Status:                dt.Status,
		Type:                  dt.PaymentMethod.ID,
		Amount:                p.Amount,
		Ticket:                dt.PointOfInteraction.TransactionData.TicketURL,
		Email:                 p.Email,
		QrCode:                dt.PointOfInteraction.TransactionData.QrCode,
		QrCodeBase64:          dt.PointOfInteraction.TransactionData.QrCodeBase64,
	}

	return out, nil
}

func (p *OutputPix) CancelPix() error {

	h := map[string]string{
		"Authorization": "Bearer " + SECRET_KEY_MP,
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
		"Authorization": "Bearer " + SECRET_KEY_MP,
		"Content-Type":  "application/json",
	}

	//	body := `{"amount": 1}`

	end := fmt.Sprintf(BASE_URL+"/v1/payments/%v/refunds", p.IDExternalTransaction)

	opt := request.NewOptions("POST", end, "", 0, h)

	resp := opt.Request()
	if resp.Err != nil {
		return resp.Err
	}

	//	fmt.Println(string(resp.Body))

	if resp.Response.StatusCode == 400 {
		return errors.New("The action requested is not valid for the current payment state")
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

func (o *OutputPix) GetStatusPayment(timeoutRequest int) error {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return errors.New("error parse location")
	}

	fout := "15:04 02/01/2006"

	opt := request.NewOptions("GET", o.Ticket, "", 0, map[string]string{})

	for {
		time.Sleep(time.Second * time.Duration(timeoutRequest))
		if time.Now().In(loc).Format(fout) == o.ExpiresAt {
			if o.Status == "pending" {
				return errors.New("client not pay")
			}
			break
		}

		resp := opt.Request()
		if resp.Err != nil {
			return resp.Err
		}

		re := regexp.MustCompile(`<h1 class="ticket__large-text">Este pagamento já foi realizado</h1>`)

		if re.Match(resp.Body) {
			return errors.New("approved")
		}
	}

	return errors.New("client not pay")
}
