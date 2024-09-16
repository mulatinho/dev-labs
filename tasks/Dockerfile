ARG GIN_MODE=debug

FROM alpine:latest
WORKDIR /app

COPY . .

RUN addgroup -S tasksgrp && adduser --home /app -S tasksuser -G tasksgrp
RUN apk update
RUN apk add --no-cache musl-dev git go ca-certificates
RUN go get -v && go build -o cmd/tasks
RUN chown -R tasksuser:tasksgrp /app && chmod +x /app/cmd/tasks

USER tasksuser:tasksgrp
EXPOSE 8181

CMD ["/app/cmd/tasks"]
