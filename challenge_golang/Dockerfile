FROM golang:1.17

WORKDIR /go/src

COPY . .

RUN go build cmd/main.go && ./main

CMD ["tail","-f","/dev/null"]
