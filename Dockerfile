FROM golang:1.18 

WORKDIR /go/src/github.com/magicworld2020/rest-api-sample/rest-api-sample
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]