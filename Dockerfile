FROM golang:1.16.3-buster as builder
WORKDIR /go/src/app
COPY . /go/src/app
RUN go build -o psping /go/src/app/cmd/psping/*.go

FROM gcr.io/distroless/base
COPY --from=builder /go/src/app/psping /
ENTRYPOINT [ "/psping" ]
