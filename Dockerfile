FROM golang:1.13 as builder

WORKDIR /build
ADD . /build/


ENV GOOS=linux GARCH=amd64 CGO_ENABLED=0
RUN go build -o api-server .

FROM alpine:3.7

#RUN adduser -S -D -H -h /app appuser
#USER appuser

WORKDIR /app
COPY --from=builder /build/api-server /app/api-server

CMD [ "/app/api-server" ]
