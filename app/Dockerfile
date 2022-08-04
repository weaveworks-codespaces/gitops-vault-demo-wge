FROM --platform=$BUILDPLATFORM golang:alpine AS gobuilder

WORKDIR /tmp/build

COPY main.go go.mod go.sum dummy_creds ./

RUN go build -o app

FROM alpine

COPY --from=gobuilder /tmp/build/app /usr/local/bin/app

ENTRYPOINT [ "/usr/local/bin/app" ]
