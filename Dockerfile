FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/ .
EXPOSE 8081
ENTRYPOINT ["./main"]
