# Build stage
FROM golang:1.26.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o kai ./cmd/kai

# Final stage
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

COPY --from=builder /app/kai /kai

USER nonroot:nonroot

ENTRYPOINT ["/kai"]
