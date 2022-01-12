FROM golang:1.17-alpine as builder
WORKDIR /go/src/someProject
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o build/projectA

FROM alpine
RUN apk add --no-cache docker
COPY --from=builder /go/src/someProject/build/projectA /usr/bin/projectA
ENTRYPOINT [ "projectA" ]