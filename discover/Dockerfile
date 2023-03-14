FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get -u github.com/gin-gonic/gin

CMD ["go", "run", "main.go"]