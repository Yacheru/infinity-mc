FROM golang:alpine AS builder

WORKDIR /deliver-service

ADD ../go.mod .

COPY . .

RUN go build -o ./build/deliver.exe ./cmd/main.go

FROM alpine

WORKDIR /deliver-service

COPY ./configs/.env /deliver-service/configs/.env
COPY --from=builder /deliver-service/build /deliver-service

CMD ["/deliver-service/deliver.exe"]