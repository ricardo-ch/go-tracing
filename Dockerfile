FROM alpine

RUN apk add --no-cache ca-certificates

COPY app /
ENTRYPOINT ["./app"]
EXPOSE 5000