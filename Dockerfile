FROM golang:1.18.2-alpine AS build
COPY ./ /go/src/github.com/hansrajrami
WORKDIR /go/src/github.com/hansrajrami
RUN go build -o log-occur-restart -v .

FROM alpine:3.12 as prod
COPY --from=build /go/src/github.com/hansrajrami/log-occur-restart /app/log-occur-restart
USER 1000
WORKDIR /app
CMD ./log-occur-restart