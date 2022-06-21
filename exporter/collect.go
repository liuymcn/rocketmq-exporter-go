package exporter

import (
	"context"
	"time"

	"github.com/rocketmq-exporter-go/admin"

	"strings"
	"sync"

	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *RocketmqExporter) collect(ch chan<- prometheus.Metric) {

	var startTime = time.Now().UnixMilli()

	e.UpdateBrokerClusterInfo()
	e.UpdateTopicRoute()

	topics := e.topics

	var waitGroup = sync.WaitGroup{}

	topicChannel := make(chan string)

	collectByTopic := func(topic string) {
		defer waitGroup.Done()

		e.CollectTopicOffset(ch, topic)

		if strings.HasPrefix(topic, RetryGroupTopicPrefix) || strings.HasPrefix(topic, DlqGroupTopicPrefix) {
			return
		}

		e.CollectBrokerStatsByTopic(ch, topic)

		groups, err := e.admin.ExamineTopicConsumeByWho(context.Background(), topic)

		if err != nil {
			return
		}

		for _, group := range groups {
			var retryTopic = RetryGroupTopicPrefix + group
			var broker = e.getRandBrokerByTopic(retryTopic)
			var brokerAddress = broker.SelectBrokerAddr()
			onlineConsumerConnection, err := e.admin.QueryConsumerConnectionInfo(context.Background(), group, brokerAddress)
			if err != nil {
				return
			}
			e.CollectConsumerOffset(ch, topic, group, onlineConsumerConnection)
			e.CollectOnlineConsumerMetric(ch, group, onlineConsumerConnection)
		}
	}

	loopTopics := func(id int) {
		ok := true
		for ok {
			topic, open := <-topicChannel
			ok = open
			if open {
				collectByTopic(topic)
			}
		}
	}

	topicSize := len(topics)
	if topicSize > 1 {
		topicSize = minx(topicSize/2, e.workers)
	}

	for w := 1; w <= topicSize; w++ {
		go loopTopics(w)
	}

	for _, topic := range topics {
		waitGroup.Add(1)
		topicChannel <- topic
	}

	close(topicChannel)

	waitGroup.Wait()

	collectByBroker := func(broker *admin.BrokerData) {
		defer waitGroup.Done()

		e.CollectBrokerRuntimeInfo(ch, broker)
		e.CollectBrokerStats(ch, broker)
	}

	brokerChannel := make(chan *admin.BrokerData)

	loopBrokers := func(id int) {
		ok := true
		for ok {
			broker, open := <-brokerChannel
			ok = open
			if open {
				collectByBroker(broker)
			}
		}
	}

	brokerSize := len(e.brokerTable)
	if brokerSize > 1 {
		brokerSize = minx(brokerSize/2, e.workers)
	}

	for w := 1; w <= brokerSize; w++ {
		go loopBrokers(w)
	}

	for _, broker := range e.brokerTable {
		waitGroup.Add(1)
		brokerChannel <- broker
	}

	close(brokerChannel)

	waitGroup.Wait()

	rlog.Info("collect finish", map[string]interface{}{
		"totalTime(ms)":      time.Now().UnixMilli() - startTime,
		"topic.size":         len(topics),
		"broker.master.size": len(e.brokerTable),
	})

}

func minx(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
