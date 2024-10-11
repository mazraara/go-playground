FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod init notes-app
RUN go get -u github.com/gin-gonic/gin
RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]