package exporter

import (
	"math/rand"
	"net/http"
	"sync"

	"github.com/rocketmq-exporter-go/admin"
	"github.com/rocketmq-exporter-go/consumer"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	ToolConsumerGroup = "TOOLS_CONSUMER"
)

const (
	INFO  = 0
	DEBUG = 1
	TRACE = 2
)

type RocketmqExporter struct {
	brokerTable    map[string]*admin.BrokerData
	topicRouteInfo map[string]*admin.TopicRouteData
	topics         []string
	admin          admin.AdminExt
	consumer       consumer.PullConsumer
	workers        int
	signMutex      sync.Mutex
	signWaitCh     chan struct{}
	collectChans   []chan<- prometheus.Metric
}

func (e *RocketmqExporter) UpdateBrokerClusterInfo() {
	clusterInfo, err := e.admin.QueryBrokerClusterInfo()
	if err != nil {
		rlog.Error("UpdateBrokerClusterInfo", map[string]interface{}{
			"err": err,
		})
		return
	}

	e.brokerTable = clusterInfo.BrokerDataMap
}

func (e *RocketmqExporter) UpdateTopicRoute() {
	topics, err := e.admin.ListTopic()
	if err != nil {
		rlog.Error("ListTopic", map[string]interface{}{
			"err": err,
		})
		return
	}

	var currentTopicRouteInfo = make(map[string]*admin.TopicRouteData)

	for _, topic := range topics {
		routeData, err := e.admin.QueryTopicRouteInfo(topic)
		if err != nil {
			rlog.Error("QueryTopicRouteInfo", map[string]interface{}{
				"err": err,
			})
			continue
		}
		currentTopicRouteInfo[topic] = routeData
	}

	e.topics = topics
	e.topicRouteInfo = currentTopicRouteInfo
}

func (e *RocketmqExporter) getRouteInfoByTopic(topic string) *admin.TopicRouteData {
	return e.topicRouteInfo[topic]
}

func (e *RocketmqExporter) getRandBrokerByTopic(topic string) *admin.BrokerData {
	var routeInfo = e.getRouteInfoByTopic(topic)
	i := rand.Int()
	i = i % len(routeInfo.BrokerDataList)
	return routeInfo.BrokerDataList[i]
}

func (e *RocketmqExporter) GetClusterByBroker(brokerName string) string {
	broker, ok := e.brokerTable[brokerName]
	if ok {
		return broker.Cluster
	} else {
		rlog.Error("GetClusterByBroker", map[string]interface{}{
			"brokerName":  brokerName,
			"brokerTable": e.brokerTable,
		})

		return "unknown"
	}
}

type RocketmqOptions struct {
	NameServerAddress []string
	Workers           int
}

// NewExporter returns an initialized Exporter.
func NewExporter(opts *RocketmqOptions) (*RocketmqExporter, error) {

	var passthroughResolver = primitive.NewPassthroughResolver(opts.NameServerAddress)

	admin, err := admin.NewAdminExt(admin.WithResolver(passthroughResolver))
	if err != nil {
		return nil, err
	}

	pullConsumer, err := consumer.NewPullConsumer(
		consumer.WithGroupName(ToolConsumerGroup),
		consumer.WithNsResolver(passthroughResolver),
	)

	if err != nil {
		rlog.Error("create pullConsumer", map[string]interface{}{
			"err": err,
		})
		return nil, err
	}

	pullConsumer.Start()

	var exporter = &RocketmqExporter{
		admin:    admin,
		consumer: pullConsumer,
		workers:  opts.Workers,
	}

	exporter.UpdateBrokerClusterInfo()
	exporter.UpdateTopicRoute()

	return exporter, nil
}

func Fly(
	listenAddress string,
	metricsPath string,
	opts *RocketmqOptions,
) {

	InitMetricDesc()

	exporter, err := NewExporter(opts)
	if err != nil {
		rlog.Fatal("create Exporter", map[string]interface{}{
			"err": err,
		})
	}

	defer exporter.admin.Close()
	defer exporter.consumer.Shutdown()

	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
	        <head><title>RocketMQ Exporter</title></head>
	        <body>
	        <h1>RocketMQ Exporter</h1>
	        <p><a href='` + metricsPath + `'>Metrics</a></p>
	        </body>
	        </html>`))
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	rlog.Fatal("ListenAndServe", map[string]interface{}{
		"err": http.ListenAndServe(listenAddress, nil),
	})

}
