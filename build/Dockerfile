FROM golang:1.11.4-alpine3.7 AS builder

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN apk update && apk add --no-cache git \
  && mkdir -p /go/bin

# Download modules
WORKDIR $GOPATH/src/seibiki
COPY go.mod go.sum $GOPATH/src/seibiki/
RUN go mod download

# Build binary
COPY cmd $GOPATH/src/seibiki/cmd
COPY internal $GOPATH/src/seibiki/internal
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/go-service ./cmd/*.go

FROM node:11.11.0-alpine as uibuilder
RUN mkdir -p /app
WORKDIR /app
COPY ./web/package.json /app
RUN npm install
COPY ./web/ /app
RUN NODE_ENV=production npm run build

FROM alpine:3.7
COPY --from=builder /go/bin/go-service /go/bin/go-service
COPY --from=uibuilder /app/build /go/bin/wwwroot
CMD ["./go/bin/go-service"]
