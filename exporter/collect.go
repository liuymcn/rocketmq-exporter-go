package exporter

import (
	"context"
	"time"

	"github.com/rocketmq-exporter-go/admin"

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

	// unique group
	// a group match normal topic and retry topic
	var groupMap = sync.Map{}

	collectByTopic := func(topic string) {
		defer waitGroup.Done()

		e.CollectTopicOffset(ch, topic)

		e.CollectBrokerStatsByTopic(ch, topic)

		groups, err := e.admin.ExamineTopicConsumeByWho(context.Background(), topic)

		if err != nil {
			rlog.Error("ExamineTopicConsumeByWho err ", map[string]interface{}{
				"topic": topic,
			})
			return
		}

		for _, group := range groups {
			var retryTopic = RetryGroupTopicPrefix + group
			var broker = e.getRandBrokerByTopic(retryTopic)
			var brokerAddress = broker.SelectBrokerAddr()
			onlineConsumerConnection, err := e.admin.QueryConsumerConnectionInfo(context.Background(), group, brokerAddress)
			if err == nil {
				if _, ok := groupMap.Load(group); !ok {
					groupMap.Store(group, onlineConsumerConnection)
				}
			}

			e.CollectConsumerOffset(ch, topic, group, onlineConsumerConnection)
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

	for w := 1; w <= e.workers; w++ {
		go loopTopics(w)
	}

	for _, topic := range topics {
		waitGroup.Add(1)
		topicChannel <- topic
	}

	close(topicChannel)

	waitGroup.Wait()

	groupChannel := make(chan string)

	collectByGroup := func(group string) {
		defer waitGroup.Done()

		onlineConsumerConnection, _ := groupMap.Load(group)

		e.CollectOnlineConsumerMetric(ch, group, onlineConsumerConnection.(*admin.ConsumerConnection))
	}

	loopGroup := func(id int) {
		ok := true
		for ok {
			group, open := <-groupChannel
			ok = open
			if open {
				collectByGroup(group)
			}
		}
	}

	for w := 1; w <= e.workers; w++ {
		go loopGroup(w)
	}

	groupMap.Range(func(key, value any) bool {
		waitGroup.Add(1)
		groupChannel <- key.(string)
		return true
	})

	close(groupChannel)

	waitGroup.Wait()

	brokerChannel := make(chan *admin.BrokerData)

	collectByBroker := func(broker *admin.BrokerData) {
		defer waitGroup.Done()

		e.CollectBrokerRuntimeInfo(ch, broker)
		e.CollectBrokerStats(ch, broker)
	}

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
