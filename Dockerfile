
FROM golang as buildando



WORKDIR /app

ADD . /app

RUN go build -o gopix



# FROM alpine

# WORKDIR /app
# COPY --from=buildando /app/gopix /app/



# ENTRYPOINT ./gopix