FROM golang:1.12.5-stretch as builder
RUN apt update && apt -y upgrade
COPY . ./cmd
RUN go build ./go/cmd/webapp

FROM amd64/debian:9.8 as production
RUN apt-get update && apt-get -y upgrade && apt-get -y install ca-certificates
COPY --from=builder /root/webapp /root/webapp
RUN go build app.go
EXPOSE 8080
CMD cd /root && ./app

