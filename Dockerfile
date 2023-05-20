FROM golang:1.18.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 8080

CMD ["/main"]

# docker commands (Ubuntu)
# docker build -t chloe2330/linklynx:1.0 .
# docker run -p 8080:8080 docker.io/chloe2330/linklynx:1.0