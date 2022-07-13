FROM docker.io/golang:1.18 as builder

WORKDIR /usr/src/intech-bot
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/intech-bot


FROM alpine

WORKDIR /bin/intech

COPY --from=builder /bin/intech-bot .

CMD ["./intech-bot"]
