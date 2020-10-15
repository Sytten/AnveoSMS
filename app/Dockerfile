# BUILDER STAGE
FROM golang:1.15-alpine as builder

ARG BUILD_VERSION

WORKDIR /build

# Install dependencies
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

# Create appuser
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 10001 \
    appuser

# Build binary
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . .

RUN go build -ldflags="-w -s -X main.BuildVersion=$BUILD_VERSION" -o main .

# RUNTIME STAGE
FROM scratch

WORKDIR /app

# Copy user and cert
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy binary and assets
COPY --from=builder /build/assets ./assets
COPY --from=builder /build/main .

USER appuser:appuser

CMD ["./main"]
