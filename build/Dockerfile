FROM golang:1.11.4-alpine3.7 AS builder
RUN apk update && apk add --no-cache git \
    && mkdir -p /go/bin

COPY ./server/go.* $GOPATH/src/seibiki/
WORKDIR $GOPATH/src/seibiki
RUN GO111MODULE=on go mod download

COPY ./server $GOPATH/src/seibiki
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/app

FROM node:11.7.0-alpine as uibuilder
RUN mkdir /app
WORKDIR /app
COPY ./web/package.json /app
RUN npm install --production
COPY ./web/ /app
RUN npm run build

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
COPY --from=uibuilder /app/build /go/bin/wwwroot
EXPOSE 3000
ENTRYPOINT ["/go/bin/app"]