package main

import (
	"fmt"
	"log"

	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/request"
)

func GetMethodsPayments() error {
	h := map[string]string{

		"accept":        "application/json",
		"content-type":  "application/json",
		"Authorization": `Bearer TEST-6812762136376103-020807-45d3ade4692fa87ac9b5b987554b77bf-811772071`,
	}

	opt := request.NewOptions("GET", "https://api.mercadopago.com/v1/payment_methods", "", 3000, h)

	r := opt.Request()
	if r.Err != nil {
		return r.Err
	}

	//fmt.Println(string(r.Body))

	return nil

}

func main() {

	// p := pix.NewPix(1, "Pagamento dos Serviços", time.Minute*10, "http://google.com.br", "wpsolucoes@gmail.com")

	// output, err := p.CreatePix()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Printf("%#v\n", output)

	o := pix.OutputPix{
		IDExternalTransaction: 54514338755,
	}

	// ok, err := o.CancelPix()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Printf("%v\n", ok)

	ok, err := o.RefundPix()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v\n", ok)

	// // e, _ := time.Parse("08/02/2023T14:57", "2023-02-08T14:57:22.931-04:00")
	// fmt.Println(e)
}
