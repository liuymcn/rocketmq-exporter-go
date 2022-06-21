package admin

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/rocketmq-exporter-go/internal"
	"github.com/rocketmq-exporter-go/internal/remote"
	"github.com/rocketmq-exporter-go/internal/utils"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

const (
	Broadcast  string = "BROADCASTING"
	Clustering string = "CLUSTERING"
)

type ConsumerConnection struct {
	Connections       []Connection                `json:"connectionSet"`
	SubscriptionTable map[string]SubscriptionData `json:"subscriptionTable"`
	ConsumeType       string                      `json:"consumeType"`
	MessageModel      string                      `json:"messageModel"`
	ConsumeFromWhere  string                      `json:"consumeFromWhere"`
}

type Connection struct {
	ClientId      string `json:"clientId"`
	ClientAddress string `json:"clientAddr"`
	Language      string `json:"language"`
	Version       int    `json:"version"`
}

type SubscriptionData struct {
	ClassFilterMode bool     `json:"classFilterMode"`
	Topic           string   `json:"topic"`
	SubString       string   `json:"subString"`
	Tags            []string `json:"tagsSet"`
	Codes           []int    `json:"codeSet"`
	SubVersion      int64    `json:"subVersion"`
	ExpressionType  string   `json:"expressionType"`
}

func (consumerConnection *ConsumerConnection) decode(data []byte) error {
	return json.Unmarshal(data, consumerConnection)
}

func (admin *adminExt) ExamineConsumerConnectionInfo(ctx context.Context, consumerGroup string) (*ConsumerConnection, error) {
	var retryTopic = internal.RetryGroupTopicPrefix + consumerGroup

	topicRouteData, err := admin.namesrv.QueryTopicRouteInfoFromServer(retryTopic)
	if err != nil {
		return nil, err
	}

	i := utils.AbsInt(rand.Int())
	i = i % len(topicRouteData.BrokerDataList)
	var broker = topicRouteData.BrokerDataList[i]

	var address = broker.SelectBrokerAddr()
	return admin.QueryConsumerConnectionInfo(ctx, consumerGroup, address)
}

func (admin *adminExt) QueryConsumerConnectionInfo(ctx context.Context, consumerGroup string, brokerAddress string) (*ConsumerConnection, error) {

	var requestHeader = &internal.GetConsumerGroupRelateRequestHeader{
		ConsumerGroup: consumerGroup,
	}

	cmd := remote.NewRemotingCommand(internal.ReqListConsumerConnection, requestHeader, nil)
	response, err := admin.cli.InvokeSync(ctx, brokerAddress, cmd, requestTimeout)

	if err != nil {
		return nil, err
	}

	switch response.Code {
	case internal.ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		consumerConnection := &ConsumerConnection{}
		err := consumerConnection.decode(response.Body)
		if err != nil {
			return nil, err
		}
		return consumerConnection, nil

	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}

}
