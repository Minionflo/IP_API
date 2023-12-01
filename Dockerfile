FROM golang:1.21.4-alpine3.18
RUN mkdir -p /app
WORKDIR /app
COPY main.go /app/
RUN go build main.go
CMD ["./main"]
