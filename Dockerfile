FROM golang:1.21.4-bullseye as builder

COPY . /opt
RUN cd /opt && go build -o bin/fevm-event-tracing cmd/busi/main.go

FROM debian:bullseye
RUN apt update && apt-get install ca-certificates -y
RUN adduser --gecos "Devops Starboard,Github,WorkPhone,HomePhone" --home /app/fevm-event-tracing --disabled-password starboard
USER starboard
COPY --from=builder /opt/bin/fevm-event-tracing /app/fevm-event-tracing/fevm-event-tracing

CMD ["--conf", "/app/fevm-event-tracing/service.conf"]
ENTRYPOINT ["/app/fevm-event-tracing/fevm-event-tracing"]
