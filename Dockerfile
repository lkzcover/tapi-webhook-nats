FROM golang:1.22 as builder
LABEL authors="lkzcover"

WORKDIR /app

RUN git clone https://github.com/lkzcover/webhook-nats .
RUN mkdir -p ./bin
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/webhook-nats ./cmd/main.go

FROM scratch

COPY --from=builder /app/bin/webhook-nats /webhook-nats

EXPOSE 8080

CMD ["/webhook-nats"]