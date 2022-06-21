package admin

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/rocketmq-exporter-go/internal"
	"github.com/rocketmq-exporter-go/internal/remote"
	"github.com/rocketmq-exporter-go/internal/utils"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/tidwall/gjson"
)

// type ConsumerRunningInfo struct {
// 	StatusTable map[string]ConsumeStatus `json:"statusTable"`
// 	Jstack      string                   `json:"jstack"`
// }

// type ConsumeStatus struct {
// 	PullRT           float64 `json:"pullRT"`
// 	PullTPS          float64 `json:"pullTPS"`
// 	ConsumeRT        float64 `json:"consumeRT"`
// 	ConsumeOKTPS     float64 `json:"consumeOKTPS"`
// 	ConsumeFailedTPS float64 `json:"consumeFailedTPS"`
// }

// func (info *internal.ConsumerRunningInfo) decode(data []byte) error {
// 	return json.Unmarshal(data, info)
// }

func (admin *adminExt) ExamineConsumerRunningInfo(ctx context.Context, consumerGroup string, clientId string, jstack bool) (*ConsumerRunningInfo, error) {
	var topic = internal.RetryGroupTopicPrefix + consumerGroup

	topicRouteData, err := admin.namesrv.QueryTopicRouteInfoFromServer(topic)
	if err != nil {
		return nil, err
	}

	i := utils.AbsInt(rand.Int())
	i = i % len(topicRouteData.BrokerDataList)
	var broker = topicRouteData.BrokerDataList[i]

	var address = broker.SelectBrokerAddr()
	return admin.QueryConsumerRunningInfo(ctx, consumerGroup, clientId, jstack, address)
}

func (admin *adminExt) QueryConsumerRunningInfo(ctx context.Context, consumerGroup string, clientId string, jstack bool, brokerAddress string) (*ConsumerRunningInfo, error) {
	var requestHeader = &internal.GetConsumerRunningInfoRequestHeader{
		ConsumerGroup: consumerGroup,
		ClientId:      clientId,
		Jstack:        jstack,
	}

	cmd := remote.NewRemotingCommand(internal.ReqGetConsumerRunningInfo, requestHeader, nil)
	response, err := admin.cli.InvokeSync(ctx, brokerAddress, cmd, requestTimeout)

	if err != nil {
		return nil, err
	}

	switch response.Code {
	case internal.ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		var statusTableResult = gjson.Parse(string(response.Body)).Get("statusTable").Map()
		var consumeStatusMap = make(map[string]internal.ConsumeStatus)

		for key, value := range statusTableResult {
			var consumeStatus internal.ConsumeStatus
			err := json.Unmarshal([]byte(value.String()), &consumeStatus)
			if err != nil {
				return nil, err
			}
			consumeStatusMap[key] = consumeStatus
		}

		if err != nil {
			return nil, err
		}

		var consumerRunningInfo = &ConsumerRunningInfo{
			StatusTable: consumeStatusMap,
		}

		return consumerRunningInfo, nil

	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}
}
