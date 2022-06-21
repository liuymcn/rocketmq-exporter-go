package internal

import "strconv"

const (
	ReqGetBrokerRuntimeInfo   = int16(28)
	ReqGetTopicStatsInfo      = int16(202)
	ReqListConsumerConnection = int16(203)
	ReqQueryConsumeStats      = int16(208)
	ReqQueryTopicConsumeByWho = int16(300)
	ReqQueryBrokerStats       = int16(315)
)

type GetTopicRelateRequestHeader struct {
	Topic string
}

func (request *GetTopicRelateRequestHeader) Encode() map[string]string {
	maps := make(map[string]string)
	maps["topic"] = request.Topic
	return maps
}

type GetConsumerGroupRelateRequestHeader struct {
	ConsumerGroup string
}

func (request *GetConsumerGroupRelateRequestHeader) Encode() map[string]string {
	maps := make(map[string]string)
	maps["consumerGroup"] = request.ConsumerGroup
	return maps
}

type GetConsumerStatsRequestHeader struct {
	Topic         string
	ConsumerGroup string
}

func (request *GetConsumerStatsRequestHeader) Encode() map[string]string {
	maps := make(map[string]string)
	maps["topic"] = request.Topic
	maps["consumerGroup"] = request.ConsumerGroup
	return maps
}

type GetConsumerRunningInfoRequestHeader struct {
	ConsumerGroup string
	ClientId      string
	Jstack        bool
}

func (request *GetConsumerRunningInfoRequestHeader) Encode() map[string]string {
	maps := make(map[string]string)
	maps["consumerGroup"] = request.ConsumerGroup
	maps["clientId"] = request.ClientId
	maps["jstackEnable"] = strconv.FormatBool(request.Jstack)
	return maps
}

type GetBrokerStatsDataRequestHeader struct {
	StatsName string
	StatsKey  string
}

func (request *GetBrokerStatsDataRequestHeader) Encode() map[string]string {
	maps := make(map[string]string)
	maps["statsName"] = request.StatsName
	maps["statsKey"] = request.StatsKey
	return maps
}
