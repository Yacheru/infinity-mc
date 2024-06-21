FROM golang:alpine3.19

WORKDIR /payments-service

COPY .. /payments-service

RUN go mod download
RUN go build -o ./build/go_build_main_go.exe ./cmd/main.go

CMD ["./build/go_build_main_go.exe"]
