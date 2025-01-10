FROM golang:1.22-alpine AS builder

COPY . .
RUN go build url-shortener-db-migrate/cmd/url-shortener-db-migrate 

FROM alpine:3.20 AS target

COPY --from=builder /go/url-shortener-db-migrate /usr/local/bin/url-shortener-db-migrate

CMD [ "/usr/local/bin/url-shortener-db-migrate" ]