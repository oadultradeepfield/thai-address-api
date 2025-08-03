# Builder Stage
FROM golang:1.24.5-bullseye AS builder
WORKDIR /app

# Install build dependencies including gcc
RUN apt-get update && apt-get install -y --no-install-recommends \
  build-essential \
  sqlite3 \
  libsqlite3-dev \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with CGO enabled for SQLite support
# Remove GOARCH=amd64 to let it auto-detect or use build-essential's gcc
RUN CGO_ENABLED=1 GOOS=linux \
  go build -a -installsuffix cgo \
  -o thai-address-server ./cmd/server

# Final Stage
FROM debian:bookworm-slim
WORKDIR /app

# Install SQLite runtime (minimal)
RUN apt-get update && apt-get install -y --no-install-recommends \
  sqlite3 \
  && rm -rf /var/lib/apt/lists/*

# Copy the built application
COPY --from=builder /app/thai-address-server .

# Copy the SQLite database file
COPY thai_address.sqlite ./

# Create a directory for SQLite data persistence (optional)
RUN mkdir -p /app/data && \
  chown -R nobody:nogroup /app

EXPOSE 8080

# Switch to non-root user
USER nobody:nogroup

ENTRYPOINT ["/app/thai-address-server"]
