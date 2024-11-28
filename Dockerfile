FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o websocket-app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/websocket-app .
EXPOSE 4000
CMD ["./websocket-app"]
