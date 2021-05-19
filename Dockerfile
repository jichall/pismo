FROM golang:latest

WORKDIR /pismo
COPY . .

EXPOSE 8090 5432

RUN go mod download
RUN go build -o "pismo.out" src/*.go

CMD ["./pismo.out"]