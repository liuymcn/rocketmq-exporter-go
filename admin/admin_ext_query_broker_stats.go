package admin

import (
	"context"
	"encoding/json"

	"github.com/rocketmq-exporter-go/internal"
	"github.com/rocketmq-exporter-go/internal/remote"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type BrokerStats struct {
	StatsMinute BrokerStatsItem `json:"statsMinute"`
	StatsHour   BrokerStatsItem `json:"statsHour"`
	StatsDay    BrokerStatsItem `json:"statsDay"`
}

type BrokerStatsItem struct {
	Sum   int64   `json:"sum"`
	Tps   float64 `json:"tps"`
	Avgpt float64 `json:"avgpt"`
}

func (stat *BrokerStats) decode(data []byte) error {
	return json.Unmarshal(data, stat)
}

func (admin *adminExt) QueryBrokerStats(ctx context.Context, statsName string, statsKey string, brokerAddress string) (*BrokerStats, error) {

	var requestHeader = &internal.GetBrokerStatsDataRequestHeader{
		StatsName: statsName,
		StatsKey:  statsKey,
	}

	cmd := remote.NewRemotingCommand(internal.ReqQueryBrokerStats, requestHeader, nil)
	response, err := admin.cli.InvokeSync(ctx, brokerAddress, cmd, requestTimeout)

	if err != nil {
		return nil, err
	}

	switch response.Code {
	case internal.ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		brokerStats := &BrokerStats{}
		brokerStats.decode(response.Body)
		return brokerStats, nil

	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}

}
