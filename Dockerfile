FROM golang:1.17.7-alpine3.15 as build

RUN mkdir /build
COPY go.mod go.sum error.go main.go model.go /build/
COPY levenshtein /build/levenshtein

WORKDIR /build
RUN go mod download

RUN go build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /build/tokogambar /app/
COPY web/ app/web
COPY images/ app/images
COPY input/ app/input
WORKDIR /app

ENTRYPOINT [ "./tokogambar" ]