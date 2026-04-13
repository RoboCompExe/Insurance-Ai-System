
FROM golang:1.21
WORKDIR /app
COPY . .
RUN go mod init backend && go mod tidy
RUN go build -o main .
CMD ["./main"]
