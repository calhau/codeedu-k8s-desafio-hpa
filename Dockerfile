FROM golang:1.14.1-alpine3.11 as builder
WORKDIR /go/src/app
COPY ./app .
RUN CGO_ENABLED=0 GOOS=linux go build bruno.go 


FROM scratch 
WORKDIR /go/src/app
COPY --from=builder /go/src/app/bruno .
CMD ["./bruno"]


