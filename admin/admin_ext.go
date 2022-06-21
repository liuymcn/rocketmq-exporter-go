package admin

import (
	"context"
	"sync"
	"time"

	"github.com/rocketmq-exporter-go/internal"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

const (
	requestTimeout = 5 * time.Second
)

type ClusterInfo struct {
	*internal.ClusterInfo

	// key -> brokerName
	BrokerDataMap map[string]*BrokerData
}

type TopicRouteData struct {
	BrokerDataList []*BrokerData
}

type ConsumerRunningInfo struct {
	// *internal.ConsumerRunningInfo
	StatusTable map[string]internal.ConsumeStatus
}

type BrokerData struct {
	*internal.BrokerData
}

type AdminExt interface {
	QueryBrokerClusterInfo() (*ClusterInfo, error)

	QueryBrokerRuntimeInfo(ctx context.Context, brokerAddress string) (*BrokerRuntimeInfo, error)

	ListTopic() ([]string, error)
	QueryTopicRouteInfo(topic string) (*TopicRouteData, error)

	ExamineTopicStats(ctx context.Context, topic string) (*TopicStatsTable, error)
	QueryTopicStats(ctx context.Context, topic string, brokerAddress string) (*TopicStatsTable, error)

	ExamineTopicConsumeByWho(ctx context.Context, topic string) ([]string, error)
	QueryTopicConsumeByWho(ctx context.Context, topic string, brokerAddress string) ([]string, error)

	ExamineConsumerConnectionInfo(ctx context.Context, consumerGroup string) (*ConsumerConnection, error)
	QueryConsumerConnectionInfo(ctx context.Context, consumerGroup string, brokerAddress string) (*ConsumerConnection, error)

	ExamineConsumeStats(ctx context.Context, consumerGroup string, topic string) (*ConsumeStats, error)
	QueryConsumeStats(ctx context.Context, consumerGroup string, topic string, brokerAddress string) (*ConsumeStats, error)

	ExamineConsumerRunningInfo(ctx context.Context, consumerGroup string, clientId string, jstack bool) (*ConsumerRunningInfo, error)
	QueryConsumerRunningInfo(ctx context.Context, consumerGroup string, clientId string, jstack bool, brokerAddress string) (*ConsumerRunningInfo, error)

	QueryBrokerStats(ctx context.Context, statsName string, statsKey string, brokerAddress string) (*BrokerStats, error)

	Close() error
}

// TODO: move outdated context to ctx
type adminExtOptions struct {
	internal.ClientOptions
}

type AdminExtOption func(options *adminExtOptions)

func defaultAdminExtOptions() *adminExtOptions {
	opts := &adminExtOptions{
		ClientOptions: internal.DefaultClientOptions(),
	}
	opts.GroupName = "TOOLS_ADMIN"
	opts.InstanceName = time.Now().String()
	return opts
}

// WithResolver nameserver resolver to fetch nameserver addr
func WithResolver(resolver primitive.NsResolver) AdminExtOption {
	return func(options *adminExtOptions) {
		options.Resolver = resolver
	}
}

type adminExt struct {
	cli     internal.RMQClient
	namesrv internal.NamesrvsExt

	opts *adminExtOptions

	closeOnce sync.Once
}

// NewAdmin initialize admin
func NewAdminExt(opts ...AdminExtOption) (AdminExt, error) {
	defaultOpts := defaultAdminExtOptions()
	for _, opt := range opts {
		opt(defaultOpts)
	}

	cli := internal.GetOrNewRocketMQClient(defaultOpts.ClientOptions, nil)
	namesrv, err := internal.NewNamesrvExt(defaultOpts.Resolver)
	if err != nil {
		return nil, err
	}

	return &adminExt{
		cli:     cli,
		namesrv: namesrv,
		opts:    defaultOpts,
	}, nil
}

type ListTopicResponse struct {
	Topics []string `json:"topicList"`
}

func (admin *adminExt) QueryBrokerClusterInfo() (*ClusterInfo, error) {
	clusterInfo, err := admin.namesrv.QueryBrokerClusterInfoFromServer()

	var brokerDataTable = make(map[string]*BrokerData)

	for brokerName, broker := range clusterInfo.BrokerDataTable {
		brokerDataTable[brokerName] = &BrokerData{
			BrokerData: broker,
		}
	}

	return &ClusterInfo{
		ClusterInfo:   clusterInfo,
		BrokerDataMap: brokerDataTable,
	}, err
}

func (admin *adminExt) ListTopic() ([]string, error) {
	return admin.namesrv.ListTopic()
}

func (admin *adminExt) QueryTopicRouteInfo(topic string) (*TopicRouteData, error) {
	topicRouteData, err := admin.namesrv.QueryTopicRouteInfoFromServer(topic)
	var brokerDataList = make([]*BrokerData, 0, len(topicRouteData.BrokerDataList))

	for _, broker := range topicRouteData.BrokerDataList {
		brokerDataList = append(brokerDataList, &BrokerData{BrokerData: broker})
	}

	return &TopicRouteData{
		BrokerDataList: brokerDataList,
	}, err
}

func (admin *adminExt) Close() error {
	admin.closeOnce.Do(func() {
		admin.cli.Shutdown()
	})
	return nil
}
