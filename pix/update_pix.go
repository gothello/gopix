package pix

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gothello/go-pix-mercado-pago/pix/request"
)

var (
	base = "https://api.mercadopago.com/v1/payments/%d"
)

func (o *OutputPix) UpdatePaymentStatus() error {

	h := map[string]string{
		"Authorization": "Bearer " + SECRET_KEY_MP,
		"Content-Type":  "application/json",
	}

	opt := request.NewOptions("GET", fmt.Sprintf(base, o.IDExternalTransaction), "", 0, h)

	resp := opt.Request()
	if resp.Err != nil {
		return resp.Err
	}

	var response map[string]interface{}

	if err := json.Unmarshal(resp.Body, &response); err != nil {
		return err
	}

	if response["status"].(string) == "approved" {
		return nil
	}

	return errors.New(response["status"].(string))
}
