# Build Stage
FROM golang:latest AS builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /todorestapi

# Run Stage
FROM alpine
WORKDIR /
COPY --from=builder /todorestapi /todorestapi
COPY .env .

EXPOSE 8080
CMD ["/todorestapi"]