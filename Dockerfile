## Builder image
#FROM golang:1.12-alpine as builder
#
#RUN apk add --no-cache git
#
#RUN mkdir -p /go/src/lottery_backend
#WORKDIR /go/src/lottery_backend
#
## Cache dependencies
#COPY go.mod .
#COPY go.sum .
#
##RUN make
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.io go build -o lottery ./src

# Executable image
FROM alpine

#COPY --from=builder /go/src/lottery_backend/lottery /lottery
COPY lottery /lottery

WORKDIR /

EXPOSE 8888

ENTRYPOINT ["/lottery"]
