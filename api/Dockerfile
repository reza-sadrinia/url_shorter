FROM golang:alpine as builder

WORKDIR /build

ADD . .

RUN go build -o main .

FROM alpine

RUN adduser -S -D -h /app appuser

USER appuser

WORKDIR /app

COPY --from=builder /build/main .
COPY . .

EXPOSE 3000

CMD ["./main"]
