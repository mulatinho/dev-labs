ARG GIN_MODE=debug

FROM alpine:latest
WORKDIR /app

COPY . .

RUN addgroup -S tasksgrp && adduser -S tasksuser -G tasksgrp
RUN apk update
RUN apk add --no-cache musl-dev git go ca-certificates
RUN go get -v && go build -o tasks
RUN chown -R tasksuser:tasksgrp /app && chmod +x /app/tasks

USER tasksuser:tasksgrp
EXPOSE 8181

CMD ["/app/tasks"]
