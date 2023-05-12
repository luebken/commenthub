FROM golang:1.18 AS build 

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=1 go build -ldflags '-extldflags "-static"' cmd/main/main.go 

FROM alpine:latest
COPY --from=build /src/main /bin/main
ENTRYPOINT ["/bin/main"]