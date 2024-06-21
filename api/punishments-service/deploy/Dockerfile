FROM golang:alpine3.19

WORKDIR /punishments-service

COPY ../ /punishments-service

RUN go mod download
RUN go build -o ./build/go_build_main_go.exe ./cmd/main.go

CMD ["./build/go_build_main_go.exe"]
