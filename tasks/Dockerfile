ARG GIN_MODE=debug

FROM --platform=linux/amd64 golang:1.19-alpine as build
WORKDIR /app

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o tasks .

FROM --platform=linux/amd64 alpine:3.18.2
COPY --from=build /app /app
WORKDIR /app

RUN apk add --no-cache libc6-compat ca-certificates file
RUN addgroup -S tasksgrp && adduser -S tasksuser -G tasksgrp
RUN chown -R tasksuser:tasksgrp /app && chmod +x /app/tasks

USER tasksuser
EXPOSE 8181

ENTRYPOINT ["/app/tasks"]
