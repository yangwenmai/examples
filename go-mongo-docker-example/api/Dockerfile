FROM golang:1.10.2

WORKDIR /go/src/app
COPY . .

RUN go get github.com/pilu/fresh

CMD [ "fresh" ]
