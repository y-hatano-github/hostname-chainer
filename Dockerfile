FROM golang:latest as builder
WORKDIR /go/src/recursive-echo
COPY . /go/src/recursive-echo
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata
WORKDIR /work
COPY --from=builder /go/src/recursive-echo/main /work
ENTRYPOINT ["./main"]
