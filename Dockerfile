FROM golang:1.21 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

ENV HOST 0.0.0.0

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/main /app/
COPY --from=builder /build/.env /app/
COPY --from=builder /build/internal/prompt /app/internal/prompt

# ca-certificates for calling gemini api
RUN apk --no-cache add ca-certificates

ENTRYPOINT ["/app/main"]

