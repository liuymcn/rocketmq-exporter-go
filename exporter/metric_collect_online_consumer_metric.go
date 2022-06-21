package exporter

import (
	"context"

	"github.com/rocketmq-exporter-go/admin"

	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *RocketmqExporter) CollectOnlineConsumerMetric(
	ch chan<- prometheus.Metric,
	group string,
	onlineConsumerConnection *admin.ConsumerConnection) {

	for _, connection := range onlineConsumerConnection.Connections {

		runningInfo, err := e.admin.ExamineConsumerRunningInfo(context.Background(), group, connection.ClientId, false)

		if err != nil {
			rlog.Error("CollectOnlineConsumerMetric ExamineConsumerRunningInfo ", map[string]interface{}{
				"group":    group,
				"clientId": connection.ClientId,
				"err":      err,
			})
			continue
		}

		for topic, consumeStatus := range runningInfo.StatusTable {
			ch <- prometheus.MustNewConstMetric(
				rocketmqClientConsumeFailMsgCount,
				prometheus.GaugeValue,
				float64(consumeStatus.ConsumeFailedMsgs),
				group,
				topic,
				connection.ClientAddress,
				connection.ClientId,
			)

			ch <- prometheus.MustNewConstMetric(
				rocketmqClientConsumeFailMsgTps,
				prometheus.GaugeValue,
				float64(consumeStatus.ConsumeFailedTPS),
				group,
				topic,
				connection.ClientAddress,
				connection.ClientId,
			)

			ch <- prometheus.MustNewConstMetric(
				rocketmqClientConsumeOkMsgTps,
				prometheus.GaugeValue,
				float64(consumeStatus.ConsumeOKTPS),
				group,
				topic,
				connection.ClientAddress,
				connection.ClientId,
			)

			ch <- prometheus.MustNewConstMetric(
				rocketmqClientConsumeRt,
				prometheus.GaugeValue,
				float64(consumeStatus.ConsumeRT),
				group,
				topic,
				connection.ClientAddress,
				connection.ClientId,
			)

			ch <- prometheus.MustNewConstMetric(
				rocketmqClientConsumerPullRt,
				prometheus.GaugeValue,
				float64(consumeStatus.PullRT),
				group,
				topic,
				connection.ClientAddress,
				connection.ClientId,
			)

			ch <- prometheus.MustNewConstMetric(
				rocketmqClientConsumerPullTps,
				prometheus.GaugeValue,
				float64(consumeStatus.PullTPS),
				group,
				topic,
				connection.ClientAddress,
				connection.ClientId,
			)
		}

	}

}
