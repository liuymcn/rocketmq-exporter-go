package exporter

import (
	"context"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	TopicPutNums    = "TOPIC_PUT_NUMS"
	TopicPutSize    = "TOPIC_PUT_SIZE"
	GroupGetNums    = "GROUP_GET_NUMS"
	GroupGetSize    = "GROUP_GET_SIZE"
	SendbackPutNums = "SNDBCK_PUT_NUMS"
)

func (e *RocketmqExporter) CollectBrokerStatsByTopic(ch chan<- prometheus.Metric, topic string) {

	groups, err := e.admin.ExamineTopicConsumeByWho(context.Background(), topic)

	if err != nil {
		glog.Errorf("CollectBrokerStatsByTopic ExamineTopicConsumeByWho topic:%s ", topic, err)
		return
	}

	for _, broker := range e.getRouteInfoByTopic(topic).BrokerDataList {

		var brokerAddress = broker.SelectBrokerAddr()

		topicPuNumBrokerStats, err := e.admin.QueryBrokerStats(context.Background(), TopicPutNums, topic, brokerAddress)
		if err != nil {
			// glog.Errorf("CollectBrokerStatsByTopic QueryBrokerStats statsName:%s, statsKey:%s, broker:%s ", TopicPutNums, topic, broker.BrokerName, err)
			continue
		}

		ch <- prometheus.MustNewConstMetric(
			rocketmqProducerTps,
			prometheus.GaugeValue,
			float64(topicPuNumBrokerStats.StatsMinute.Tps),
			broker.Cluster,
			broker.BrokerName,
			topic,
		)

		topicPutSizeBrokerStats, err := e.admin.QueryBrokerStats(context.Background(), TopicPutSize, topic, brokerAddress)
		if err != nil {
			// glog.Errorf("CollectBrokerStatsByTopic QueryBrokerStats statsName:%s, statsKey:%s, broker:%s ", TopicPutSize, topic, broker.BrokerName, err)
			continue
		}

		ch <- prometheus.MustNewConstMetric(
			rocketmqProducerMessageSize,
			prometheus.GaugeValue,
			float64(topicPutSizeBrokerStats.StatsMinute.Tps),
			broker.Cluster,
			broker.BrokerName,
			topic,
		)

		for _, group := range groups {
			var statsKey = topic + "@" + group

			groupGetNumBrokerStats, err := e.admin.QueryBrokerStats(context.Background(), GroupGetNums, statsKey, brokerAddress)
			if err != nil {
				// glog.Errorf("CollectBrokerStatsByTopic QueryBrokerStats statsName:%s, statsKey:%s, broker:%s ", GroupGetNums, statsKey, broker.BrokerName, err)
				continue
			}

			ch <- prometheus.MustNewConstMetric(
				rocketmqConsumerTps,
				prometheus.GaugeValue,
				float64(groupGetNumBrokerStats.StatsMinute.Tps),
				broker.Cluster,
				broker.BrokerName,
				topic,
				group,
			)

			groupGetSizeBrokerStats, err := e.admin.QueryBrokerStats(context.Background(), GroupGetSize, statsKey, brokerAddress)
			if err != nil {
				// glog.Errorf("CollectBrokerStatsByTopic QueryBrokerStats statsName:%s, statsKey:%s, broker:%s ", GroupGetSize, statsKey, broker.BrokerName, err)
				continue
			}

			ch <- prometheus.MustNewConstMetric(
				rocketmqConsumerMessageSize,
				prometheus.GaugeValue,
				float64(groupGetSizeBrokerStats.StatsMinute.Tps),
				broker.Cluster,
				broker.BrokerName,
				topic,
				group,
			)

			sendbackPutNumsBrokerStats, err := e.admin.QueryBrokerStats(context.Background(), SendbackPutNums, statsKey, brokerAddress)
			if err != nil {
				// glog.Errorf("CollectBrokerStatsByTopic QueryBrokerStats statsName:%s, statsKey:%s, broker:%s ", SendbackPutNums, statsKey, broker.BrokerName, err)
				continue
			}

			ch <- prometheus.MustNewConstMetric(
				rocketmqSendBackNums,
				prometheus.GaugeValue,
				float64(sendbackPutNumsBrokerStats.StatsMinute.Sum),
				broker.Cluster,
				broker.BrokerName,
				topic,
				group,
			)

		}

	}

}
