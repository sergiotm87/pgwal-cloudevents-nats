FROM golang:1.18.1-stretch as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/pgwalstreams

FROM scratch

WORKDIR /app

COPY --from=build /bin/pgwalstreams /bin/pgwalstreams
COPY --from=build /app/config.yml /app/config.yml

ENTRYPOINT ["/bin/pgwalstreams"]
