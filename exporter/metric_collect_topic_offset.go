package exporter

import (
	"context"
	"strconv"
	"strings"

	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *RocketmqExporter) CollectTopicOffset(ch chan<- prometheus.Metric, topic string) {

	topicStatsTable, err := e.admin.ExamineTopicStats(context.Background(), topic)

	if err != nil {
		rlog.Error("CollectOnlineConsumerMetric ExamineConsumerRunningInfo ", map[string]interface{}{
			"topic": topic,
			"err":   err,
		})
		return
	}

	brokerOffsetMap := make(map[string]int64)
	brokerUpdateTimestampMap := make(map[string]int64)

	for messageQueue, topicOffset := range topicStatsTable.OffsetTable {

		var brokerName = messageQueue.BrokerName

		if maxOffset, ok := brokerOffsetMap[brokerName]; ok {
			brokerOffsetMap[brokerName] = maxOffset + topicOffset.MaxOffset
		} else {
			brokerOffsetMap[brokerName] = topicOffset.MaxOffset
		}

		if lastUpdateTimestamp, ok := brokerUpdateTimestampMap[brokerName]; ok {
			if topicOffset.LastUpdateTimestamp > lastUpdateTimestamp {
				brokerUpdateTimestampMap[brokerName] = topicOffset.LastUpdateTimestamp
			}
		} else {
			brokerUpdateTimestampMap[brokerName] = topicOffset.LastUpdateTimestamp
		}

	}

	for brokerName, offset := range brokerOffsetMap {
		lastUpdateTimestamp := brokerUpdateTimestampMap[brokerName]
		if strings.HasPrefix(topic, RetryGroupTopicPrefix) {
			ch <- prometheus.MustNewConstMetric(
				rocketmqTopicRetryOffset,
				prometheus.GaugeValue,
				float64(offset),
				e.GetClusterByBroker(brokerName),
				brokerName,
				topic,
				strconv.FormatInt(lastUpdateTimestamp, 10),
			)
		} else if strings.HasPrefix(topic, DlqGroupTopicPrefix) {
			group := strings.Replace(topic, DlqGroupTopicPrefix, "", 1)
			ch <- prometheus.MustNewConstMetric(
				rocketmqTopicDlqOffset,
				prometheus.GaugeValue,
				float64(offset),
				e.GetClusterByBroker(brokerName),
				brokerName,
				group,
				strconv.FormatInt(lastUpdateTimestamp, 10),
			)
		} else {
			ch <- prometheus.MustNewConstMetric(
				rocketmqProducerOffset,
				prometheus.GaugeValue,
				float64(offset),
				e.GetClusterByBroker(brokerName),
				brokerName,
				topic,
				strconv.FormatInt(lastUpdateTimestamp, 10),
			)
		}

	}

}
