FROM golang:1.22

WORKDIR /message_tui

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go test -v -c tests/blockchain_test.go

CMD ["./tests.test"]