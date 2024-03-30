FROM golang:1.22-alpine

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.7

RUN apk add --no-cache ca-certificates

COPY --from=0 /app .

CMD ["./main"]
