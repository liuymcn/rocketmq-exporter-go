package admin

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/rocketmq-exporter-go/internal"
	"github.com/rocketmq-exporter-go/internal/remote"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type Groups struct {
	Groups []string `json:"groupList"`
}

func (groups *Groups) decode(data []byte) error {
	return json.Unmarshal(data, groups)
}

func (admin *adminExt) ExamineTopicConsumeByWho(ctx context.Context, topic string) ([]string, error) {

	routeData, err := admin.namesrv.QueryTopicRouteInfoFromServer(topic)

	if err != nil {
		return nil, err
	}

	for _, broker := range routeData.BrokerDataList {

		var address = broker.SelectBrokerAddr()

		groups, err := admin.QueryTopicConsumeByWho(ctx, topic, address)

		if err != nil {
			return nil, err
		} else {
			return groups, nil
		}
	}

	return nil, primitive.NewRemotingErr("brokerData len is unexpect: " + strconv.Itoa(len(routeData.BrokerDataList)))
}

func (admin *adminExt) QueryTopicConsumeByWho(ctx context.Context, topic string, brokerAddress string) ([]string, error) {

	var requestHeader = &internal.GetTopicRelateRequestHeader{
		Topic: topic,
	}

	cmd := remote.NewRemotingCommand(internal.ReqQueryTopicConsumeByWho, requestHeader, nil)
	response, err := admin.cli.InvokeSync(ctx, brokerAddress, cmd, requestTimeout)

	if err != nil {
		return nil, err
	}

	switch response.Code {
	case internal.ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		group := &Groups{}
		err := group.decode(response.Body)

		if err != nil {
			return nil, err
		}

		return group.Groups, nil

	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}
}
