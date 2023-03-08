# API para geraçao de Pix 

- Esta é uma api que permite a criaçao, cancelamento e reembolso de transaçoes Pix. Ela tambem permite uso de sistema de mensageria RabbitMQ para proccessamento, alem dos enpoints via http.

## Instalaçao

- go get github.com/gothello/gopix

## Necessario:
- Conhecimentos basicos docker.
- Docker instalado na maquina.

## Configuração:
- Configure suas credencias de accesso a api mercado pago, mysql, rabbitmq, e uma chave privada que sera analisada a  cada request na api via endpoints http.
- arquvo de configuracao usado é o yml
- caso o arquivo de configuracao esteja fora do diretorio main defina a variavel config.Path recendo o diretorio onde se encontra o mesmo.

