package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/rabbit"
	"github.com/gothello/go-pix-mercado-pago/usecase"
	"github.com/gothello/go-pix-mercado-pago/utils"
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
	flag.StringVar(&API_PORT, "port", "3000", "api port to running api")
	flag.StringVar(&SECRET_AUTH_KEY, "secret", "", "key to authorization access on api")
	flag.Parse()

	if SECRET_AUTH_KEY == "" {
		var err error

		SECRET_AUTH_KEY, err = utils.GenerateAuthAcess()
		if err != nil {
			log.Fatalln(err)
		}

		utils.SECRET_AUTH_KEY = SECRET_AUTH_KEY

		color.Red("YOU NOT PREDEFINED KEY")
		color.Green("SYSTEM GENERATE KEY FOR YOU")
		color.Green(SECRET_AUTH_KEY)
	}
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
	if err := http.ListenAndServe(fmt.Sprintf(":%s", API_PORT), nil); err != nil {
		log.Fatalln(err)
	}
}
