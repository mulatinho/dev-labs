FROM golang:1.19

WORKDIR /app

COPY . .

RUN go get
RUN go build

EXPOSE 8181

ENTRYPOINT ["/app/tasks"]
