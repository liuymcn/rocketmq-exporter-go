package admin

import (
	"context"
	"strconv"

	"github.com/rocketmq-exporter-go/internal"
	"github.com/rocketmq-exporter-go/internal/remote"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/tidwall/gjson"
)

type TopicOffset struct {
	MinOffset           int64 `json:"minOffset"`
	MaxOffset           int64 `json:"maxOffset"`
	LastUpdateTimestamp int64 `json:"lastUpdateTimestamp"`
}

type TopicStatsTable struct {
	OffsetTable map[*primitive.MessageQueue]*TopicOffset `json:"offsetTable"`
}

const (
	BrokerNameIndex          = 0
	QueueIdIndex             = 1
	TopicIndex               = 2
	LastUpdateTimestampIndex = 3
	MaxOffsetIndex           = 4
	MinOffsetIndex           = 5
)

// TODO
// Note: because of lacking of the method of
// deserializing the type of map[interface{}]interface{}
func (table *TopicStatsTable) decode(data []byte) {

	offsetTableResult := gjson.Parse(string(data)).Get("offsetTable")

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
		minOffset, _ := convertToInt64(values[i+MinOffsetIndex])
		maxOffset, _ := convertToInt64(values[i+MaxOffsetIndex])
		lastUpdateTimestamp, _ := convertToInt64(values[i+LastUpdateTimestampIndex])
		var topicOffset = &TopicOffset{
			MinOffset:           minOffset,
			MaxOffset:           maxOffset,
			LastUpdateTimestamp: lastUpdateTimestamp,
		}
		if table.OffsetTable == nil {
			table.OffsetTable = make(map[*primitive.MessageQueue]*TopicOffset)
		}
		table.OffsetTable[messageQueue] = topicOffset
	}
}

func convertToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 64)
}

func (admin *adminExt) ExamineTopicStats(ctx context.Context, topic string) (*TopicStatsTable, error) {

	routeData, err := admin.namesrv.QueryTopicRouteInfoFromServer(topic)

	if err != nil {
		return nil, err
	}

	topicStatsTable := &TopicStatsTable{
		OffsetTable: make(map[*primitive.MessageQueue]*TopicOffset),
	}

	for _, broker := range routeData.BrokerDataList {

		var address = broker.SelectBrokerAddr()

		topicStatsTableItem, err := admin.QueryTopicStats(ctx, topic, address)
		if err != nil {
			return nil, err
		}
		for k, v := range topicStatsTableItem.OffsetTable {
			topicStatsTable.OffsetTable[k] = v
		}
	}

	return topicStatsTable, nil
}

func (admin *adminExt) QueryTopicStats(ctx context.Context, topic string, brokerAddress string) (*TopicStatsTable, error) {
	var requestHeader = &internal.GetTopicRelateRequestHeader{
		Topic: topic,
	}

	cmd := remote.NewRemotingCommand(internal.ReqGetTopicStatsInfo, requestHeader, nil)
	response, err := admin.cli.InvokeSync(ctx, brokerAddress, cmd, requestTimeout)

	if err != nil {
		return nil, err
	}

	switch response.Code {
	case internal.ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		// TOFIX: cannot deserialize the type map[interface{}]interface{}
		topicStatsTable := &TopicStatsTable{}
		topicStatsTable.decode(response.Body)

		return topicStatsTable, nil

	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}
}
