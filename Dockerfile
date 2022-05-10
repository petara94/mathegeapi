FROM golang:1.18.1

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app

EXPOSE 8080

CMD ["/app/app"]
