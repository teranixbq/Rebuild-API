FROM golang:1.20-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

######## Start a new stage from scratch #######
FROM alpine:latest  

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["./main"]
