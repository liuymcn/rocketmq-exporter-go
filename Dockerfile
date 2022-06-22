FROM quay.io/prometheus/busybox:latest
LABEL  maintainer="The Authors <liuym.1225836327@gmail.com>"

COPY ./rocketmq-exporter /home/exporter/rocketmq-exporter

WORKDIR /home/exporter/

ENTRYPOINT ["/home/exporter/rocketmq-exporter"]
