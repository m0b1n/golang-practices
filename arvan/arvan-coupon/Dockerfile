FROM golang:1.16.8-alpine AS build

RUN apk update
RUN apk add build-base
RUN mkdir -p $GOPATH/src/arvan/
WORKDIR $GOPATH/src/arvan/

COPY . .

RUN go get -d -v ./...

RUN go build -o /out/main .

FROM alpine AS bin
COPY --from=build /out/main main
ENTRYPOINT [ "./main" ]