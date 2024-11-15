FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY app/server.go .
COPY app/template.html .
COPY VERSION .
RUN GOOS=linux GOARCH=amd64 go build -o server server.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/template.html .
COPY --from=builder /app/VERSION .
RUN chmod +x ./server
EXPOSE 8080
CMD ["./server"]