FROM golang:1.13.4-buster
RUN go get github.com/cespare/reflex
WORKDIR /app
ENTRYPOINT ["reflex", "-c", "reflex.conf"]