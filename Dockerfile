# Build stage
FROM golang:1.19.9-alpine3.18 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 8080
CMD ["/app/main"]
