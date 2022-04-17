FROM golang:1.18.1-stretch
RUN go install github.com/cespare/reflex@latest
WORKDIR /app
ENTRYPOINT ["reflex", "-c", "reflex.conf"]