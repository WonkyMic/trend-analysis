FROM --platform=arm64 golang:latest AS builder
LABEL maintainer="Michael Barefield <wonkymic@gmail.com>"
WORKDIR /src
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm64
RUN go build -o app

FROM --platform=arm64 alpine:latest
COPY --from=builder /src/app /app
CMD ["/app"]
EXPOSE 8080