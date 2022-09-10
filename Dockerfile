FROM golang:1.19-buster

RUN go version

ENV GOPATH=/
COPY ./ ./
# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]
