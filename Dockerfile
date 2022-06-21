FROM quay.io/prometheus/busybox:latest

COPY ./rocketmq-exporter /home/exporter/rocketmq-exporter

WORKDIR /home/exporter/

ENTRYPOINT ["/home/exporter/rocketmq-exporter"]