FROM golang:latest


WORKDIR /app

COPY . .

RUN go build -o app .

CMD ["./app"]

# RUN rm config.yml


