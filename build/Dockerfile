FROM alpine:3.15

RUN mkdir -p /etc/account/{bin,conf}
COPY bin/busi /etc/account/bin/busi

CMD ["--conf", "/etc/account/conf/service.conf"]
ENTRYPOINT ["/etc/account/bin/busi"]