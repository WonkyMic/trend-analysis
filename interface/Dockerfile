FROM golang:1.18.0-armv7 AS builder
LABEL maintainer="Michael Barefield <wonkymic@gmail.com>"
WORKDIR /go
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o app

FROM alpine:3.15-armv7
COPY --from=builder /go/app /app
CMD ["/app"]
EXPOSE 8080