FROM golang:1.19

WORKDIR /app

COPY . .

RUN go get
RUN go build

ENTRYPOINT ["/app/tasks"]
