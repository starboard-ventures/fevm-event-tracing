FROM golang:1.18.3-bullseye as builder

COPY . /opt
RUN cd /opt && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/busi cmd/account/main.go

FROM alpine:3.15.4
RUN mkdir -p /app/account-backend
RUN adduser -h /app/account-backend -D starboard
USER starboard
COPY --from=builder /opt/bin/busi /app/account-backend/busi

CMD ["--conf", "/app/account-backend/service.conf"]
ENTRYPOINT ["/app/account-backend/busi"]
