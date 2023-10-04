FROM golang:1.21.0

WORKDIR /app

ADD . /app

COPY .env /app/cmd/.env

RUN apt-get update && apt-get install -y postgresql-client

COPY ./postgres.sh /app/wait-for-postgres.sh

RUN chmod +x /app/postgres.sh

WORKDIR /app/cmd

RUN go mod download

RUN go test -v ./...

RUN go build -o main .

CMD ["/app/cmd/main"]
