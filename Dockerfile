FROM quay.io/prometheus/busybox:latest

LABEL  maintainer="The Authors <liuym.1225836327@gmail.com>"

ARG TARGETARCH

COPY ./rocketmq-exporter /home/linux-${TARGETARCH}/exporter/rocketmq-exporter

WORKDIR /home/linux-${TARGETARCH}/exporter/

ENTRYPOINT ["/home/exporter/rocketmq-exporter"]
