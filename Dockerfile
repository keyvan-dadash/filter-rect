FROM golang AS builder

WORKDIR /go/src/frect

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o frect .

FROM ubuntu:latest  

WORKDIR /root/

COPY --from=builder /go/src/frect .

EXPOSE 8080

CMD ["./frect"]