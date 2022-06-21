package admin

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/rocketmq-exporter-go/internal"
	"github.com/rocketmq-exporter-go/internal/remote"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

type BrokerRuntimeKVInfo struct {
	Info map[string]string `json:"table"`
}

func (info *BrokerRuntimeKVInfo) decode(data []byte) error {
	return json.Unmarshal(data, info)
}

func (admin *adminExt) QueryBrokerRuntimeInfo(ctx context.Context, brokerAddress string) (*BrokerRuntimeInfo, error) {

	cmd := remote.NewRemotingCommand(internal.ReqGetBrokerRuntimeInfo, nil, nil)
	response, err := admin.cli.InvokeSync(ctx, brokerAddress, cmd, requestTimeout)

	if err != nil {
		return nil, err
	}

	switch response.Code {
	case internal.ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		kvInfo := &BrokerRuntimeKVInfo{}
		kvInfo.decode(response.Body)
		return newBrokerRuntimeInfo(*kvInfo), nil

	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}

}

type BrokerRuntimeInfo struct {
	MsgPutTotalTodayNow                   int64
	MsgGetTotalTodayNow                   int64
	MsgPutTotalTodayMorning               int64
	MsgGetTotalTodayMorning               int64
	MsgPutTotalYesterdayMorning           int64
	MsgGetTotalYesterdayMorning           int64
	SendThreadPoolQueueHeadWaitTimeMills  int64
	QueryThreadPoolQueueHeadWaitTimeMills int64
	PullThreadPoolQueueHeadWaitTimeMills  int64
	QueryThreadPoolQueueSize              int64
	PullThreadPoolQueueSize               int64
	SendThreadPoolQueueCapacity           int64
	PullThreadPoolQueueCapacity           int64
	CommitLogMinOffset                    int64
	CommitLogMaxOffset                    int64
	BootTimestamp                         int64
	DispatchMaxBuffer                     int64
	PageCacheLockTimeMills                int64
	GetMessageEntireTimeMax               int64
	PutMessageTimesTotal                  int64
	SendThreadPoolQueueSize               int64
	StartAcceptSendRequestTimestamp       int64
	PutMessageEntireTimeMax               int64
	EarliestMessageTimestamp              int64
	RemainTransientStoreBufferNumbs       int64
	QueryThreadPoolQueueCapacity          int64
	DispatchBehindBytes                   int64
	PutMessageSizeTotal                   int64
	PutMessageAverageSize                 float64
	RemainHowManyDataToFlush              float64
	CommitLogDirCapacityTotal             float64
	CommitLogDirCapacityFree              float64
	ConsumeQueueDiskRatio                 float64
	CommitLogDiskRatio                    float64
	Runtime                               string
	BrokerVersion                         int64
	BrokerVersionDesc                     string
	PutTps                                *PutTps
	GetMissTps                            *PutTps
	GetTransferedTps                      *PutTps
	GetTotalTps                           *PutTps
	GetFoundTps                           *PutTps
	PutMessageDistributeTimeMap           map[string]int64
}

type PutTps struct {
	Ten        float64
	Sixty      float64
	SixHundred float64
}

func newBrokerRuntimeInfo(kvInfo BrokerRuntimeKVInfo) *BrokerRuntimeInfo {
	var table = kvInfo.Info
	commitLogDiskRatio, _ := strconv.ParseFloat(table["commitLogDiskRatio"], 64)
	commitLogDirCapacityArray := strings.Split(table["commitLogDirCapacity"], " ")
	commitLogDirCapacityTotal, _ := strconv.ParseFloat(commitLogDirCapacityArray[2], 64)
	commitLogDirCapacityFree, _ := strconv.ParseFloat(commitLogDirCapacityArray[6], 64)

	commitLogMaxOffset, _ := strconv.ParseInt(table["commitLogMaxOffset"], 10, 64)
	commitLogMinOffset, _ := strconv.ParseInt(table["commitLogMinOffset"], 10, 64)
	msgPutTotalTodayNow, _ := strconv.ParseInt(table["msgPutTotalTodayNow"], 10, 64)
	msgGetTotalTodayNow, _ := strconv.ParseInt(table["msgGetTotalTodayNow"], 10, 64)
	dispatchBehindBytes, _ := strconv.ParseInt(table["dispatchBehindBytes"], 10, 64)
	putMessageAverageSize, _ := strconv.ParseFloat(table["putMessageAverageSize"], 64)
	queryThreadPoolQueueCapacity, _ := strconv.ParseInt(table["queryThreadPoolQueueCapacity"], 10, 64)
	remainTransientStoreBufferNumbs, _ := strconv.ParseInt(table["remainTransientStoreBufferNumbs"], 10, 64)
	earliestMessageTimestamp, _ := strconv.ParseInt(table["earliestMessageTimeStamp"], 10, 64)
	putMessageEntireTimeMax, _ := strconv.ParseInt(table["putMessageEntireTimeMax"], 10, 64)
	startAcceptSendRequestTimestamp, _ := strconv.ParseInt(table["startAcceptSendRequestTimeStamp"], 10, 64)
	sendThreadPoolQueueSize, _ := strconv.ParseInt(table["sendThreadPoolQueueSize"], 10, 64)
	getMessageEntireTimeMax, _ := strconv.ParseInt(table["getMessageEntireTimeMax"], 10, 64)
	pageCacheLockTimeMills, _ := strconv.ParseInt(table["pageCacheLockTimeMills"], 10, 64)
	consumeQueueDiskRatio, _ := strconv.ParseFloat(table["consumeQueueDiskRatio"], 64)
	dispatchMaxBuffer, _ := strconv.ParseInt(table["dispatchMaxBuffer"], 10, 64)
	pullThreadPoolQueueCapacity, _ := strconv.ParseInt(table["pullThreadPoolQueueCapacity"], 10, 64)
	sendThreadPoolQueueCapacity, _ := strconv.ParseInt(table["sendThreadPoolQueueCapacity"], 10, 64)
	pullThreadPoolQueueSize, _ := strconv.ParseInt(table["pullThreadPoolQueueSize"], 10, 64)
	queryThreadPoolQueueSize, _ := strconv.ParseInt(table["queryThreadPoolQueueSize"], 10, 64)
	pullThreadPoolQueueHeadWaitTimeMills, _ := strconv.ParseInt(table["pullThreadPoolQueueHeadWaitTimeMills"], 10, 64)
	queryThreadPoolQueueHeadWaitTimeMills, _ := strconv.ParseInt(table["queryThreadPoolQueueHeadWaitTimeMills"], 10, 64)
	sendThreadPoolQueueHeadWaitTimeMills, _ := strconv.ParseInt(table["sendThreadPoolQueueHeadWaitTimeMills"], 10, 64)
	msgGetTotalYesterdayMorning, _ := strconv.ParseInt(table["msgGetTotalYesterdayMorning"], 10, 64)
	msgPutTotalYesterdayMorning, _ := strconv.ParseInt(table["msgPutTotalYesterdayMorning"], 10, 64)
	msgGetTotalTodayMorning, _ := strconv.ParseInt(table["msgGetTotalTodayMorning"], 10, 64)
	msgPutTotalTodayMorning, _ := strconv.ParseInt(table["msgPutTotalTodayMorning"], 10, 64)
	remainHowManyDataToFlush, _ := strconv.ParseFloat(table["remainHowManyDataToFlush"], 64)
	putMessageTimesTotal, _ := strconv.ParseInt(table["putMessageTimesTotal"], 10, 64)

	bootTimestamp, _ := strconv.ParseInt(table["bootTimestamp"], 10, 64)
	brokerVersion, _ := strconv.ParseInt(table["brokerVersion"], 10, 64)

	getFoundTps := newPutTps(table["getFoundTps"])
	getTotalTps := newPutTps(table["getTotalTps"])
	getTransferedTps := newPutTps(table["getTransferedTps"])
	getMissTps := newPutTps(table["getMissTps"])
	putTps := newPutTps(table["putTps"])

	putMessageDistributeTime := buildPutMessageDistributeTime(table["putMessageDistributeTime"])

	info := &BrokerRuntimeInfo{
		BrokerVersionDesc:                     table["brokerVersionDesc"],
		BootTimestamp:                         bootTimestamp,
		BrokerVersion:                         brokerVersion,
		CommitLogDirCapacityTotal:             commitLogDirCapacityTotal,
		CommitLogDirCapacityFree:              commitLogDirCapacityFree,
		CommitLogMaxOffset:                    commitLogMaxOffset,
		CommitLogMinOffset:                    commitLogMinOffset,
		MsgPutTotalTodayNow:                   msgPutTotalTodayNow,
		MsgGetTotalTodayNow:                   msgGetTotalTodayNow,
		DispatchBehindBytes:                   dispatchBehindBytes,
		PutMessageAverageSize:                 putMessageAverageSize,
		QueryThreadPoolQueueCapacity:          queryThreadPoolQueueCapacity,
		RemainTransientStoreBufferNumbs:       remainTransientStoreBufferNumbs,
		EarliestMessageTimestamp:              earliestMessageTimestamp,
		PutMessageEntireTimeMax:               putMessageEntireTimeMax,
		StartAcceptSendRequestTimestamp:       startAcceptSendRequestTimestamp,
		SendThreadPoolQueueSize:               sendThreadPoolQueueSize,
		GetMessageEntireTimeMax:               getMessageEntireTimeMax,
		PageCacheLockTimeMills:                pageCacheLockTimeMills,
		ConsumeQueueDiskRatio:                 consumeQueueDiskRatio,
		DispatchMaxBuffer:                     dispatchMaxBuffer,
		PullThreadPoolQueueCapacity:           pullThreadPoolQueueCapacity,
		SendThreadPoolQueueCapacity:           sendThreadPoolQueueCapacity,
		PullThreadPoolQueueSize:               pullThreadPoolQueueSize,
		QueryThreadPoolQueueSize:              queryThreadPoolQueueSize,
		PullThreadPoolQueueHeadWaitTimeMills:  pullThreadPoolQueueHeadWaitTimeMills,
		QueryThreadPoolQueueHeadWaitTimeMills: queryThreadPoolQueueHeadWaitTimeMills,
		SendThreadPoolQueueHeadWaitTimeMills:  sendThreadPoolQueueHeadWaitTimeMills,
		MsgGetTotalYesterdayMorning:           msgGetTotalYesterdayMorning,
		MsgPutTotalYesterdayMorning:           msgPutTotalYesterdayMorning,
		MsgGetTotalTodayMorning:               msgGetTotalTodayMorning,
		MsgPutTotalTodayMorning:               msgPutTotalTodayMorning,
		RemainHowManyDataToFlush:              remainHowManyDataToFlush,
		PutMessageTimesTotal:                  putMessageTimesTotal,
		CommitLogDiskRatio:                    commitLogDiskRatio,
		PutTps:                                putTps,
		GetFoundTps:                           getFoundTps,
		GetTotalTps:                           getTotalTps,
		GetTransferedTps:                      getTransferedTps,
		GetMissTps:                            getMissTps,
		PutMessageDistributeTimeMap:           putMessageDistributeTime,
	}

	return info
}

func newPutTps(tpsStr string) *PutTps {

	tps := strings.Split(tpsStr, " ")

	ten, _ := strconv.ParseFloat(tps[0], 64)
	sixty, _ := strconv.ParseFloat(tps[1], 64)
	sixHundred, _ := strconv.ParseFloat(tps[2], 64)

	return &PutTps{
		Ten:        ten,
		Sixty:      sixty,
		SixHundred: sixHundred,
	}
}

func buildPutMessageDistributeTime(putMessageTime string) map[string]int64 {
	var putMessageDistributeTime = make(map[string]int64)
	putTimes := strings.Split(strings.Trim(putMessageTime, " "), " ")
	for _, value := range putTimes {
		elements := strings.Split(value, ":")
		key := strings.Replace(strings.Replace(elements[0], "[", "", 1), "]", "", 1)
		// TODO strange logs because of the invalid of the data
		if len(elements) < 2 {
			rlog.Error("buildPutMessageDistributeTime err ", map[string]interface{}{
				"elements": elements,
			})
			putMessageDistributeTime[key] = 0
		} else {
			number, _ := strconv.ParseInt(elements[1], 10, 64)
			putMessageDistributeTime[key] = number
		}

	}

	return putMessageDistributeTime
}
