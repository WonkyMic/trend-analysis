FROM --platform=arm64 golang:1.21 AS builder
LABEL maintainer="Michael Barefield <wonkymic@gmail.com>"
WORKDIR /src
# COPY go.mod go.sum ./
# RUN go mod download
COPY . ./
RUN go build -o app

FROM --platform=arm64 alpine:latest
COPY --from=builder /src/app /app
CMD ["/app"]
EXPOSE 8080