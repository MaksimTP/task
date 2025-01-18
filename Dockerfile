FROM golang:1.22.2-alpine AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .
ENV GOOS=linux
RUN go mod download
RUN go build -o crypto ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/crypto /app/crypto
COPY --from=builder /app/cryptosymbols.json /app/cryptosymbols.json
EXPOSE 8081
RUN chmod +x /app/crypto
CMD ["./crypto"]