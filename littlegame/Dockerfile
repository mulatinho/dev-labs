ARG GIN_MODE=debug

FROM alpine:latest
WORKDIR /app

COPY . .

RUN addgroup -S littlegamegrp && adduser --home /app -S littlegameuser -G littlegamegrp
RUN apk update
RUN apk add --no-cache musl-dev git go ca-certificates
RUN go get -v && go build -o littlegame
RUN chown -R littlegameuser:littlegamegrp /app && chmod +x /app/littlegame

USER littlegameuser:littlegamegrp
EXPOSE 8181

CMD ["/app/littlegame"]
