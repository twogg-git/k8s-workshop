FROM golang:1.10-alpine3.7 as builder
WORKDIR /go/src/k8s-workshop/k8s
COPY k8s.go .
RUN go get -d ./... && go build -o k8s .

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/k8s-workshop/k8s .

EXPOSE 9090
ENTRYPOINT ./k8s