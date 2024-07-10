FROM golang:alpine AS builder

WORKDIR /punishments-service

ADD ../go.mod .

COPY . .

RUN go build -o ./build/punishments.exe ./cmd/main.go

FROM alpine

WORKDIR /punishments-service

COPY ./configs/.env /punishments-service/configs/.env
COPY --from=builder /punishments-service/build /punishments-service

CMD ["/punishments-service/punishments.exe"]