FROM alpine:3.8
RUN apk add --no-cache ca-certificates tzdata
COPY build/* /bin/
ENTRYPOINT cron-get
CMD [cron-get]