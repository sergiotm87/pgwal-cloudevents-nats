FROM alpine:3.10
RUN adduser -D developer
WORKDIR /app
COPY wal-listener .
USER developer

ENTRYPOINT ["./wal-listener"]