FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Copy module files (from repository root -> api_server)
# Build MUST be run with the repository root as the build context.
COPY api_server/go.mod api_server/go.sum ./
RUN go mod download

# Copy only the OCR service source into the build context
COPY api_server/services/ocr_service ./

# Build the binary (output name matches the CMD below)
RUN go build -o ocr_svc ocr_svc.go

FROM alpine:3.18
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/ocr_svc ./ocr_svc

# The Go service listens on port 8081 (see source), so expose that
EXPOSE 8081

CMD ["./ocr_svc"]