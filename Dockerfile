FROM golang:1.26-alpine AS builder

# curl fetches the standalone tailwindcss CLI; libstdc++/libgcc are its runtime
# deps. No CGO toolchain needed (SQLite driver is pure Go via modernc.org/sqlite).
RUN apk --no-cache add curl libstdc++ libgcc

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Pin templ to match go.mod
RUN go install github.com/a-h/templ/cmd/templ@v0.3.1001

# Tailwindcss standalone CLI. Use the musl build (this is an Alpine/musl image)
# and pick the arch from the buildx TARGETARCH so it works on amd64 and arm64.
ARG TARGETARCH
RUN ARCH=$([ "$TARGETARCH" = "arm64" ] && echo arm64 || echo x64) && \
    curl -sL "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-${ARCH}-musl" \
    -o /usr/local/bin/tailwindcss && chmod +x /usr/local/bin/tailwindcss

COPY . .

RUN templ generate && \
    tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o server ./cmd/api/main.go

FROM alpine:3.22

RUN apk --no-cache add wget ca-certificates

COPY --from=builder /app/server /server

# The app listens on $PORT (random if unset); default to 8080 to match EXPOSE,
# the health check, and the DO app spec. DO can still override PORT.
ENV PORT=8080

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -qO- http://localhost:8080/ || exit 1

USER nobody

CMD ["/server"]
