FROM golang:1.22.0-alpine

WORKDIR /go/src/app
COPY . .

RUN apk upgrade --update && apk --no-cache add git

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]
