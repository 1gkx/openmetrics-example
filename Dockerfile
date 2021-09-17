FROM golang:1.17-alpine as build

ARG CGO_ENABLED=0

WORKDIR /out/app
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w" -o /out/app

FROM scratch AS bin
COPY --from=build /out/app /app
ENTRYPOINT ["/usr/app"]
