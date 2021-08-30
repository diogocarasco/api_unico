
FROM golang:1.17-alpine3.14 as builder

LABEL maintainer="Diogo Carasco <carasco@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod download 

RUN go build -o unico_api main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/unico_api .

EXPOSE 8080

CMD ["./unico_api"]
