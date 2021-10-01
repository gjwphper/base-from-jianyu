FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /tt/src/github.com/gjwphper/base
COPY . /tt/src/github.com/gjwphper/base
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./main"]