FROM golang:latest

WORKDIR /go/remote-schema
COPY . .
RUN ls -la

RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air","-c", ".air.toml"]