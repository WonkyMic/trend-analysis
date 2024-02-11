FROM golang:latest
LABEL maintainer="Michael Barefield <wonkymic@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main
EXPOSE 8080
CMD ["/main"]