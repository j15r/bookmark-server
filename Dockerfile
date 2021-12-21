# syntax=docker/dockerfile:1

################################
#             build            #
################################
FROM golang:1.17-alpine as build

WORKDIR /srv

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o bookmark-server ./cmd/server.go

################################
#             prod             #
################################
FROM gcr.io/distroless/base-debian10

WORKDIR /srv

COPY --from=build /srv/bookmark-server ./bookmark-server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/bookmark-server"]