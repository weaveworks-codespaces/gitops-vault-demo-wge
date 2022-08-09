FROM --platform=$BUILDPLATFORM golang:alpine AS gobuilder

WORKDIR /tmp/build

COPY . .

RUN go build -o app

FROM alpine

COPY --from=gobuilder /tmp/build/ /demo/

WORKDIR /demo

ENTRYPOINT [ "/demo/app" ]
