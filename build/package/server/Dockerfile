FROM golang:1.17.7-alpine3.15 as build
WORKDIR /go/src/github.com/riandyrn/tokogambar

ADD go.mod go.sum .env ./
RUN go mod download

ADD ./cmd ./cmd
ADD ./internal ./internal
ADD ./static ./static

WORKDIR /go/src/github.com/riandyrn/tokogambar/cmd/server
RUN go build -o server

FROM alpine:3.15
RUN apk add ca-certificates
COPY --from=build /go/src/github.com/riandyrn/tokogambar/cmd/server .
COPY ./cmd/server .
COPY .env .
COPY ./static ./static

ENTRYPOINT [ "./server" ]