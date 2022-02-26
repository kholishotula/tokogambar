FROM golang:1.17.7-alpine3.15 as build
WORKDIR /tokogambar

COPY go.mod go.sum .env ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./static ./static

WORKDIR /tokogambar/cmd/server
RUN go build -o server

FROM alpine:3.15
RUN apk add ca-certificates
COPY --from=build /tokogambar/cmd/server .
COPY ./cmd/server .
COPY .env .
COPY ./static ./static

ENTRYPOINT [ "./server" ]