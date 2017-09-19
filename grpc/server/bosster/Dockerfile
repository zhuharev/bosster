FROM golang:alpine

RUN apk add --update --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]

# Document that the service listens on port 2020.
EXPOSE 2020
