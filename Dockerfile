FROM golang:1.12-stretch AS base

WORKDIR /src/
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY main.go go.mod go.sum /src/

RUN go mod download
RUN go build -o /src/blogapi

FROM scratch

WORKDIR /src/
COPY --from=base /src/blogapi /src/blogapi

EXPOSE 8080
ENTRYPOINT ["/src/blogapi"]
