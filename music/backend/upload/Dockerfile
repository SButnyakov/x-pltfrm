FROM golang:1.23-alpine

WORKDIR /upload

COPY . .

RUN go mod download && go mod verify
RUN go mod tidy
RUN go build -C ./cmd -o ./../main  

CMD ["./main"]
