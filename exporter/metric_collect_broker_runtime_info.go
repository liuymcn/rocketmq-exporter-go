package exporter

import (
	"context"
	"strconv"

	"github.com/rocketmq-exporter-go/admin"

	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *RocketmqExporter) CollectBrokerRuntimeInfo(ch chan<- prometheus.Metric, broker *admin.BrokerData) {

	var brokerAddress = broker.SelectBrokerAddr()

	runtime, err := e.admin.QueryBrokerRuntimeInfo(context.Background(), brokerAddress)

	if err != nil {
		rlog.Error("CollectBrokerRuntimeInfo QueryBrokerRuntimeInfo broker:%s ", map[string]interface{}{
			"brokerName": broker.BrokerName,
			"err":        err,
		})
		return
	}

	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt0ms,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["<=0ms"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt0to10ms,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["0~10ms"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt10to50ms,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["10~50ms"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt50to100ms,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["50~100ms"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt100to200ms,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["100~200ms"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt200to500ms,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["200~500ms"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt500to1s,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["500ms~1s"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt1to2s,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["1~2s"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt2to3s,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["2~3s"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt3to4s,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["3~4s"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt4to5s,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["4~5s"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt5to10s,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["5~10s"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePmdt10stomore,
		prometheus.GaugeValue,
		float64(runtime.PutMessageDistributeTimeMap["10s~"]),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeMsgPutTotalTodayNow,
		prometheus.GaugeValue,
		float64(runtime.MsgPutTotalTodayNow),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeMsgGetTotalTodayNow,
		prometheus.GaugeValue,
		float64(runtime.MsgGetTotalTodayNow),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeDispatchBehindBytes,
		prometheus.GaugeValue,
		float64(runtime.DispatchBehindBytes),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePutMessageSizeTotal,
		prometheus.GaugeValue,
		float64(runtime.PutMessageSizeTotal),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePutMessageAverageSize,
		prometheus.GaugeValue,
		runtime.PutMessageAverageSize,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeQueryThreadPoolQueueCapacity,
		prometheus.GaugeValue,
		float64(runtime.QueryThreadPoolQueueCapacity),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeRemainTransientStoreBufferNumbs,
		prometheus.GaugeValue,
		float64(runtime.RemainTransientStoreBufferNumbs),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeEarliestMessageTimestamp,
		prometheus.GaugeValue,
		float64(runtime.EarliestMessageTimestamp),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePutMessageEntireTimeMax,
		prometheus.GaugeValue,
		float64(runtime.PutMessageEntireTimeMax),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeStartAcceptSendrequestTime,
		prometheus.GaugeValue,
		float64(runtime.StartAcceptSendRequestTimestamp),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeSendThreadPoolQueueSize,
		prometheus.GaugeValue,
		float64(runtime.SendThreadPoolQueueSize),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePutMessageTimesTotal,
		prometheus.GaugeValue,
		float64(runtime.PutMessageTimesTotal),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetMessageEntireTimeMax,
		prometheus.GaugeValue,
		float64(runtime.GetMessageEntireTimeMax),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePageCacheLockTimeMills,
		prometheus.GaugeValue,
		float64(runtime.PageCacheLockTimeMills),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeCommitLogDiskRatio,
		prometheus.GaugeValue,
		runtime.CommitLogDiskRatio,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeConsumeQueueDiskRatio,
		prometheus.GaugeValue,
		runtime.ConsumeQueueDiskRatio,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetFoundTps600,
		prometheus.GaugeValue,
		runtime.GetFoundTps.SixHundred,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetFoundTps60,
		prometheus.GaugeValue,
		runtime.GetFoundTps.Sixty,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetFoundTps10,
		prometheus.GaugeValue,
		runtime.GetFoundTps.Ten,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetTotalTps600,
		prometheus.GaugeValue,
		runtime.GetTotalTps.SixHundred,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetTotalTps60,
		prometheus.GaugeValue,
		runtime.GetTotalTps.Sixty,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetTotalTps10,
		prometheus.GaugeValue,
		runtime.GetTotalTps.Ten,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetTransferedTps600,
		prometheus.GaugeValue,
		runtime.GetTransferedTps.SixHundred,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetTransferedTps60,
		prometheus.GaugeValue,
		runtime.GetTransferedTps.Sixty,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetTransferedTps10,
		prometheus.GaugeValue,
		runtime.GetTransferedTps.Ten,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetMissTps600,
		prometheus.GaugeValue,
		runtime.GetMissTps.SixHundred,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetMissTps60,
		prometheus.GaugeValue,
		runtime.GetMissTps.Sixty,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeGetMissTps10,
		prometheus.GaugeValue,
		runtime.GetMissTps.Ten,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePutTps600,
		prometheus.GaugeValue,
		runtime.PutTps.SixHundred,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePutTps60,
		prometheus.GaugeValue,
		runtime.PutTps.Sixty,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePutTps10,
		prometheus.GaugeValue,
		runtime.PutTps.Ten,
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeDispatchMaxBuffer,
		prometheus.GaugeValue,
		float64(runtime.DispatchMaxBuffer),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePullThreadPoolQueueCapacity,
		prometheus.GaugeValue,
		float64(runtime.PullThreadPoolQueueCapacity),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeSendThreadPoolQueueCapacity,
		prometheus.GaugeValue,
		float64(runtime.SendThreadPoolQueueCapacity),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePullThreadPoolQueueSize,
		prometheus.GaugeValue,
		float64(runtime.PullThreadPoolQueueSize),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeQueryThreadPoolQueueSize,
		prometheus.GaugeValue,
		float64(runtime.QueryThreadPoolQueueSize),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimePullThreadPoolQueueHeadWaitTimeMills,
		prometheus.GaugeValue,
		float64(runtime.PullThreadPoolQueueHeadWaitTimeMills),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeQueryThreadPoolQueueHeadWaitTimeMills,
		prometheus.GaugeValue,
		float64(runtime.QueryThreadPoolQueueHeadWaitTimeMills),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeSendThreadPoolQueueHeadWaitTimeMills,
		prometheus.GaugeValue,
		float64(runtime.SendThreadPoolQueueHeadWaitTimeMills),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeMsgGetTotalYesterdayMorning,
		prometheus.GaugeValue,
		float64(runtime.MsgGetTotalYesterdayMorning),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeMsgPutTotalYesterdayMorning,
		prometheus.GaugeValue,
		float64(runtime.MsgPutTotalYesterdayMorning),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeMsgGetTotalTodayMorning,
		prometheus.GaugeValue,
		float64(runtime.MsgGetTotalTodayMorning),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeMsgPutTotalTodayMorning,
		prometheus.GaugeValue,
		float64(runtime.MsgPutTotalTodayMorning),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeCommitLogDirCapacityFree,
		prometheus.GaugeValue,
		float64(runtime.CommitLogDirCapacityFree),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeCommitLogDirCapacityTotal,
		prometheus.GaugeValue,
		float64(runtime.CommitLogDirCapacityTotal),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeCommitLogMaxOffset,
		prometheus.GaugeValue,
		float64(runtime.CommitLogMaxOffset),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeCommitLogMinOffset,
		prometheus.GaugeValue,
		float64(runtime.CommitLogMinOffset),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)
	ch <- prometheus.MustNewConstMetric(
		rocketmqBrokeRuntimeRemainHowManyDataToFlush,
		prometheus.GaugeValue,
		float64(runtime.RemainHowManyDataToFlush),
		broker.Cluster,
		brokerAddress,
		strconv.Itoa(int(runtime.BrokerVersion)),
		runtime.BrokerVersionDesc,
		strconv.Itoa(int(runtime.BootTimestamp)),
	)

}
