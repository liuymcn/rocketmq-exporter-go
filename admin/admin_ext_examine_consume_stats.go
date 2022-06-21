package admin

import (
	"context"
	"strconv"

	"github.com/rocketmq-exporter-go/internal"
	"github.com/rocketmq-exporter-go/internal/remote"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/tidwall/gjson"
)

type ConsumeStats struct {
	OffsetTable map[*primitive.MessageQueue]*OffsetWrapper `json:"offsetTable"`
	ConsumeTps  float64                                    `json:"consumeTps"`
}

type OffsetWrapper struct {
	BrokerOffset   int64 `json:"brokerOffset"`
	ConsumerOffset int64 `json:"consumerOffset"`
	LastTimestamp  int64 `json:"lastTimestamp"`
}

const (
	BrokerOffsetIndex   = 3
	ConsumerOffsetIndex = 4
	LastTimestampIndex  = 5
)

func (consumeStats *ConsumeStats) decode(data []byte) {
	var dataStr = string(data)

	var dataResult = gjson.Parse(dataStr)

	var consumeTps = dataResult.Get("consumeTps").Float()
	consumeStats.ConsumeTps = consumeTps

	var offsetTableResult = dataResult.Get("offsetTable")

	values := make([]string, 0)
	offsetTableResult.ForEach(func(key, value gjson.Result) bool {
		values = append(values, value.String())
		return true
	})

	for i := 0; i < len(values); i += 6 {
		queueId, _ := strconv.Atoi(values[i+QueueIdIndex])
		var messageQueue = &primitive.MessageQueue{
			Topic:      values[i+TopicIndex],
			BrokerName: values[i+BrokerNameIndex],
			QueueId:    queueId,
		}
		brokerOffset, _ := convertToInt64(values[i+BrokerOffsetIndex])
		consumerOffset, _ := convertToInt64(values[i+ConsumerOffsetIndex])
		lastTimestamp, _ := convertToInt64(values[i+LastTimestampIndex])
		var offsetWrapper = &OffsetWrapper{
			BrokerOffset:   brokerOffset,
			ConsumerOffset: consumerOffset,
			LastTimestamp:  lastTimestamp,
		}
		if consumeStats.OffsetTable == nil {
			consumeStats.OffsetTable = make(map[*primitive.MessageQueue]*OffsetWrapper)
		}
		consumeStats.OffsetTable[messageQueue] = offsetWrapper
	}

}

func (admin *adminExt) ExamineConsumeStats(ctx context.Context, consumerGroup string, topic string) (*ConsumeStats, error) {
	var retryTopic = internal.RetryGroupTopicPrefix + consumerGroup

	retryTopicRouteData, err := admin.namesrv.QueryTopicRouteInfoFromServer(retryTopic)
	if err != nil {
		return nil, err
	}

	var finalConsumeStats = &ConsumeStats{
		OffsetTable: make(map[*primitive.MessageQueue]*OffsetWrapper),
		ConsumeTps:  0,
	}

	for _, broker := range retryTopicRouteData.BrokerDataList {
		var address = broker.SelectBrokerAddr()
		consumeStats, err := admin.QueryConsumeStats(ctx, consumerGroup, topic, address)

		if err != nil {
			return nil, err
		}

		finalConsumeStats.ConsumeTps = finalConsumeStats.ConsumeTps + consumeStats.ConsumeTps

		for k, v := range consumeStats.OffsetTable {
			finalConsumeStats.OffsetTable[k] = v
		}
	}

	return finalConsumeStats, nil
}

func (admin *adminExt) QueryConsumeStats(ctx context.Context, consumerGroup string, topic string, brokerAddress string) (*ConsumeStats, error) {

	var requestHeader = &internal.GetConsumerStatsRequestHeader{
		Topic:         topic,
		ConsumerGroup: consumerGroup,
	}

	cmd := remote.NewRemotingCommand(internal.ReqQueryConsumeStats, requestHeader, nil)
	response, err := admin.cli.InvokeSync(ctx, brokerAddress, cmd, requestTimeout)

	if err != nil {
		return nil, err
	}

	switch response.Code {
	case internal.ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		consumeStats := &ConsumeStats{}
		consumeStats.decode(response.Body)
		return consumeStats, nil

	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}

}
