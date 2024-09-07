FROM golang:latest

WORKDIR /app
COPY . .
RUN go mod tidy

ENV TERM=xterm

CMD ["go", "run", "./server/server.go"]
