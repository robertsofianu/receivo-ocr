FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY ../ ../go.mod ../../go.sum ./
RUN go mod download

COPY . .

RUN go build -o ocr_service ./services/ocr_service/ocr_service.go

COPY --from=builder /app/ocr_service .

EXPOSE 8080

CMD ["./ocr_svc"]