FROM golang:alpine AS builder

WORKDIR /payments-service

ADD ../go.mod .

COPY . .

RUN go build -o ./build/payment.exe ./cmd/main.go

FROM alpine

WORKDIR /payments-service

COPY ./configs/.env /payments-service/configs/.env
COPY --from=builder /payments-service/build /payments-service

CMD ["/payments-service/payment.exe"]