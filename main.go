package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/rabbit"
	"github.com/gothello/go-pix-mercado-pago/usecase"
	"github.com/gothello/go-pix-mercado-pago/web"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	QUEUES = map[string]string{
		"CREATE":   "newpix",
		"CREATED":  "createdpix",
		"APPROVED": "approved",
	}

	API_PORT        = ""
	SECRET_AUTH_KEY = ""
)

func init() {
}

func LoadAllUseCases(service *entity.RespositoryMySql) *web.PixHandlers {
	create := usecase.NewCreatePixUseCase(service)
	cancel := usecase.NewCancelUseCase(service)
	refund := usecase.NewRefundUseCase(service)
	find := usecase.NewFindPixUseCase(service)
	findall := usecase.NewFindAllPixUseCase(service)

	handlers := web.NewPixHandlers(create, cancel, refund, find, findall)

	return handlers
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Println(err)
	}

	rep := entity.NewRespositoryMySql(db)
	handlers := LoadAllUseCases(rep)
	handlers.LoadRoutes()

	rabbitUseCase := usecase.NewRabbitConnectionUseCase(conn, rep)
	inChan := make(chan rabbit.RabbitInputChan)
	go rabbit.Consumer(conn, QUEUES["CREATE"], inChan)
	go rabbitUseCase.RabbitCreatePixUseCase(inChan, rep, QUEUES)

	log.Println("api running port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln(err)
	}
}
