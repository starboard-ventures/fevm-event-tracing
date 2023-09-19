FROM golang:1.20.8-bullseye as builder

COPY . /opt
RUN cd /opt && go build -o bin/fevm-event-tracking cmd/busi/main.go

FROM debian:bullseye
RUN apt update && apt-get install ca-certificates -y
RUN adduser --gecos "Devops Starboard,Github,WorkPhone,HomePhone" --home /app/fevm-event-tracking --disabled-password starboard
USER starboard
COPY --from=builder /opt/bin/fevm-event-tracking /app/fevm-event-tracking/fevm-event-tracking

CMD ["--conf", "/app/fevm-event-tracking/service.conf"]
ENTRYPOINT ["/app/fevm-event-tracking/fevm-event-tracking"]
