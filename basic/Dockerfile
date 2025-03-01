FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get -u github.com/gin-gonic/gin

# RUN go get -u github.com/cosmtrek/air
# ENV AIR_PORT 8080
# ENV AIR_HOT_RELOAD 1

CMD ["go", "run", "main.go"]