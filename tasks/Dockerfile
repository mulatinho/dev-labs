FROM golang:1.19 as build

WORKDIR /app
COPY . .
RUN go get
RUN go build

FROM alpine:3.18.2
COPY --from=build /app/tasks /app/tasks

RUN apk add --no-cache libc6-compat
RUN addgroup -S tasksgrp && adduser -S tasksuser -G tasksgrp
RUN chown -R tasksuser:tasksgrp /app
USER tasksuser
EXPOSE 8181

ENTRYPOINT ["/app/tasks"]
