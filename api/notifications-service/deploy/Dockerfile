FROM golang:alpine AS builder

WORKDIR /notifications-service

ADD ../go.mod .

COPY . .

RUN go build -o ./build/notifications.exe ./cmd/main.go

FROM alpine

WORKDIR /notifications-service

COPY ./configs/.env /notifications-service/configs/.env
COPY --from=builder /notifications-service/build /notifications-service

CMD ["/notifications-service/notifications.exe"]