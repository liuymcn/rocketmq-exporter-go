FROM quay.io/prometheus/busybox:latest

LABEL  maintainer="The Authors <liuym.1225836327@gmail.com>"

ARG TARGETOS
ARG TARGETARCH

COPY ./${TARGETOS}-${TARGETARCH}/rocketmq-exporter-go /home/exporter/rocketmq-exporter

WORKDIR /home/exporter/

ENTRYPOINT ["/home/exporter/rocketmq-exporter"]
