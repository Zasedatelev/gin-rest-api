FROM golang:1.20-buster AS stage

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh



RUN go mod download
RUN go build -o gin-rest-api ./cmd/app/main.go

CMD ["./git-rest-api"]
