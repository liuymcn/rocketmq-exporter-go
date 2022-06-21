package exporter

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/rocketmq-exporter-go/admin"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/prometheus/client_golang/prometheus"
)

// exclude the topic which has the prefix of "%DLQ%"
func (e *RocketmqExporter) CollectConsumerOffset(
	ch chan<- prometheus.Metric,
	topic string,
	group string,
	onlineConsumerConnection *admin.ConsumerConnection,

) {

	if strings.HasPrefix(topic, DlqGroupTopicPrefix) {
		return
	}

	var countOfOnlineConsumers = 0
	var messageModel = admin.Clustering

	if onlineConsumerConnection.MessageModel == "" {
		messageModel = onlineConsumerConnection.MessageModel
	}

	if onlineConsumerConnection.Connections != nil {
		countOfOnlineConsumers = len(onlineConsumerConnection.Connections)
	}

	if countOfOnlineConsumers > 0 {

		var clientAddresses = make([]string, countOfOnlineConsumers)
		var clientIds = make([]string, countOfOnlineConsumers)

		for _, connection := range onlineConsumerConnection.Connections {
			clientAddresses = append(clientAddresses, connection.ClientAddress)
			clientIds = append(clientIds, connection.ClientId)
		}

		ch <- prometheus.MustNewConstMetric(
			rocketmqGroupCount,
			prometheus.GaugeValue,
			float64(countOfOnlineConsumers),
			strings.Join(clientAddresses, ","),
			strings.Join(clientIds, ","),
			topic,
			group,
		)

	}

	consumeStats, err := e.admin.ExamineConsumeStats(context.Background(), group, topic)

	if err != nil {
		rlog.Error("CollectConsumerOffset ExamineConsumeStats ", map[string]interface{}{
			"group": group,
			"topic": topic,
		})
		return
	}

	var diffTotal int64
	var consumerOffsetMap = make(map[string]int64)
	var consumerLatencyMap = make(map[string]int64)
	for queue, offset := range consumeStats.OffsetTable {

		var diff = offset.BrokerOffset - offset.ConsumerOffset
		diffTotal = diff

		var brokerName = queue.BrokerName
		if consumerOffset, ok := consumerOffsetMap[brokerName]; ok {
			consumerOffsetMap[brokerName] = consumerOffset + offset.ConsumerOffset
		} else {
			consumerOffsetMap[brokerName] = offset.ConsumerOffset
		}

		var consumerLatency int64 = 0

		pullResult, err := e.consumer.PullFrom(context.Background(), queue, offset.ConsumerOffset, 1)
		if err != nil {
			rlog.Error("CollectConsumerOffset PullFrom", map[string]interface{}{
				"queue":          queue,
				"consumerOffset": offset.ConsumerOffset,
				"err":            err,
			})
			return
		}

		// strange logic
		if pullResult.Status == primitive.PullFound {
			if diff != 0 {
				consumerLatency = time.Now().UnixMilli() - pullResult.GetMessageExts()[0].StoreTimestamp
			}
		} else if pullResult.Status == primitive.PullOffsetIllegal {
			pullResult, err = e.consumer.PullFrom(context.Background(), queue, pullResult.MinOffset, 1)
			if err != nil {
				rlog.Error("CollectConsumerOffset PullFrom ", map[string]interface{}{
					"queue":          queue,
					"consumerOffset": pullResult.MinOffset,
					"err":            err,
				})
				return
			}

			consumerLatency = time.Now().UnixMilli() - pullResult.GetMessageExts()[0].StoreTimestamp
		} else {
			continue
		}

		if latency, ok := consumerLatencyMap[brokerName]; ok {
			if consumerLatency > latency {
				consumerLatencyMap[brokerName] = consumerLatency
			}
		} else {
			consumerLatencyMap[brokerName] = consumerLatency
		}

	}

	if messageModel == admin.Clustering {

		if strings.HasPrefix(topic, RetryGroupTopicPrefix) {
			ch <- prometheus.MustNewConstMetric(
				rocketmqGroupRetryDiff,
				prometheus.GaugeValue,
				float64(diffTotal),
				group,
				topic,
				strconv.Itoa(countOfOnlineConsumers),
				messageModel,
			)
		} else if strings.HasPrefix(topic, DlqGroupTopicPrefix) {
			// TODO this topic has been already excluded
			ch <- prometheus.MustNewConstMetric(
				rocketmqGroupDlqDiff,
				prometheus.GaugeValue,
				float64(diffTotal),
				group,
				topic,
				strconv.Itoa(countOfOnlineConsumers),
				messageModel,
			)
		} else {
			ch <- prometheus.MustNewConstMetric(
				rocketmqGroupDiff,
				prometheus.GaugeValue,
				float64(diffTotal),
				group,
				topic,
				strconv.Itoa(countOfOnlineConsumers),
				messageModel,
			)
		}

	}

	// get consumer offset
	for brokerName, offset := range consumerOffsetMap {
		ch <- prometheus.MustNewConstMetric(
			rocketmqConsumerOffset,
			prometheus.GaugeValue,
			float64(offset),
			e.GetClusterByBroker(brokerName),
			brokerName,
			topic,
			group,
		)
	}

	// get consumer latency
	for brokerName, lactency := range consumerLatencyMap {
		ch <- prometheus.MustNewConstMetric(
			rocketmqGroupGetLatencyByStoreTime,
			prometheus.GaugeValue,
			float64(lactency),
			e.GetClusterByBroker(brokerName),
			brokerName,
			topic,
			group,
		)
	}

}
