FROM golang:1.15.6 as builder

WORKDIR /go/src/mailer-api
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM scratch
ENV ADMIN_USER=admin \
    ADMIN_PASSWORD=admin
COPY --from=builder /go/src/mailer-api/app .
CMD ["./app"]

EXPOSE 8080