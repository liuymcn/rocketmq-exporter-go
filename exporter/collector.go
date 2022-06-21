package exporter

import (
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/prometheus/client_golang/prometheus"
)

// Collect fetches the stats from configured Rocketmq location and delivers them
// as Prometheus metrics. It implements prometheus.Collector.
func (e *RocketmqExporter) Collect(ch chan<- prometheus.Metric) {

	// Locking to avoid race add
	e.signMutex.Lock()
	e.collectChans = append(e.collectChans, ch)
	// Safe to compare length since we own the Lock
	if len(e.collectChans) == 1 {
		e.signWaitCh = make(chan struct{})
		go e.collectForChans(e.signWaitCh)
	} else {
		rlog.Info("concurrent calls detected, waiting for first to finish", nil)
	}
	// Put in another variable to ensure not overwriting it in another Collect once we wait
	waiter := e.signWaitCh
	e.signMutex.Unlock()
	// Released lock, we have insurance that our chan will be part of the collectChan slice
	<-waiter
	// collectChan finished

}

func (e *RocketmqExporter) collectForChans(quit chan struct{}) {
	original := make(chan prometheus.Metric)
	container := make([]prometheus.Metric, 0, 100)

	go func() {
		for metric := range original {
			container = append(container, metric)
		}
	}()

	e.collect(original)
	close(original)

	// Lock to avoid modification on the channel slice
	e.signMutex.Lock()

	// all collection task will be return the same result mestric once collection period
	for _, ch := range e.collectChans {
		for _, metric := range container {
			ch <- metric
		}
	}

	// Reset the slice
	e.collectChans = e.collectChans[:0]
	// Notify remaining waiting Collect they can return
	close(quit)
	// Release the lock so Collect can append to the slice again
	e.signMutex.Unlock()
}

// Describe describes all the metrics ever exported by the Rocketmq exporter. It
// implements prometheus.Collector.
func (e *RocketmqExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- rocketmqGroupDiff
	ch <- rocketmqGroupRetryDiff
	ch <- rocketmqGroupDlqDiff
	ch <- rocketmqGroupCount
	ch <- rocketmqProducerOffset
	ch <- rocketmqTopicRetryOffset
	ch <- rocketmqTopicDlqOffset
	ch <- rocketmqClientConsumeFailMsgCount
	ch <- rocketmqClientConsumeFailMsgTps
	ch <- rocketmqClientConsumeOkMsgTps
	ch <- rocketmqClientConsumeRt
	ch <- rocketmqClientConsumerPullRt
	ch <- rocketmqClientConsumerPullTps
	ch <- rocketmqBrokerTps
	ch <- rocketmqBrokerQps
	ch <- rocketmqBrokeRuntimePmdt0ms
	ch <- rocketmqBrokeRuntimePmdt0to10ms
	ch <- rocketmqBrokeRuntimePmdt10to50ms
	ch <- rocketmqBrokeRuntimePmdt50to100ms
	ch <- rocketmqBrokeRuntimePmdt100to200ms
	ch <- rocketmqBrokeRuntimePmdt200to500ms
	ch <- rocketmqBrokeRuntimePmdt500to1s
	ch <- rocketmqBrokeRuntimePmdt1to2s
	ch <- rocketmqBrokeRuntimePmdt2to3s
	ch <- rocketmqBrokeRuntimePmdt3to4s
	ch <- rocketmqBrokeRuntimePmdt4to5s
	ch <- rocketmqBrokeRuntimePmdt5to10s
	ch <- rocketmqBrokeRuntimePmdt10stomore
	ch <- rocketmqConsumerTps
	ch <- rocketmqGroupConsumeTps
	ch <- rocketmqConsumerOffset
	ch <- rocketmqGroupConsumeTotalOffset
	ch <- rocketmqConsumerMessageSize
	ch <- rocketmqSendBackNums
	ch <- rocketmqGroupGetLatencyByStoreTime
	ch <- rocketmqProducerTps
	ch <- rocketmqProducerMessageSize
	ch <- rocketmqBrokeRuntimeMsgPutTotalTodayNow
	ch <- rocketmqBrokeRuntimeMsgGetTotalTodayNow
	ch <- rocketmqBrokeRuntimeDispatchBehindBytes
	ch <- rocketmqBrokeRuntimePutMessageSizeTotal
	ch <- rocketmqBrokeRuntimePutMessageAverageSize
	ch <- rocketmqBrokeRuntimeQueryThreadPoolQueueCapacity
	ch <- rocketmqBrokeRuntimeRemainTransientStoreBufferNumbs
	ch <- rocketmqBrokeRuntimeEarliestMessageTimestamp
	ch <- rocketmqBrokeRuntimePutMessageEntireTimeMax
	ch <- rocketmqBrokeRuntimeStartAcceptSendrequestTime
	ch <- rocketmqBrokeRuntimeSendThreadPoolQueueSize
	ch <- rocketmqBrokeRuntimePutMessageTimesTotal
	ch <- rocketmqBrokeRuntimeGetMessageEntireTimeMax
	ch <- rocketmqBrokeRuntimePageCacheLockTimeMills
	ch <- rocketmqBrokeRuntimeCommitLogDiskRatio
	ch <- rocketmqBrokeRuntimeConsumeQueueDiskRatio
	ch <- rocketmqBrokeRuntimeGetFoundTps600
	ch <- rocketmqBrokeRuntimeGetFoundTps60
	ch <- rocketmqBrokeRuntimeGetFoundTps10
	ch <- rocketmqBrokeRuntimeGetTotalTps600
	ch <- rocketmqBrokeRuntimeGetTotalTps60
	ch <- rocketmqBrokeRuntimeGetTotalTps10
	ch <- rocketmqBrokeRuntimeGetTransferedTps600
	ch <- rocketmqBrokeRuntimeGetTransferedTps60
	ch <- rocketmqBrokeRuntimeGetTransferedTps10
	ch <- rocketmqBrokeRuntimeGetMissTps600
	ch <- rocketmqBrokeRuntimeGetMissTps60
	ch <- rocketmqBrokeRuntimeGetMissTps10
	ch <- rocketmqBrokeRuntimePutTps600
	ch <- rocketmqBrokeRuntimePutTps60
	ch <- rocketmqBrokeRuntimePutTps10
	ch <- rocketmqBrokeRuntimeDispatchMaxBuffer
	ch <- rocketmqBrokeRuntimePullThreadPoolQueueCapacity
	ch <- rocketmqBrokeRuntimeSendThreadPoolQueueCapacity
	ch <- rocketmqBrokeRuntimePullThreadPoolQueueSize
	ch <- rocketmqBrokeRuntimeQueryThreadPoolQueueSize
	ch <- rocketmqBrokeRuntimePullThreadPoolQueueHeadWaitTimeMills
	ch <- rocketmqBrokeRuntimeQueryThreadPoolQueueHeadWaitTimeMills
	ch <- rocketmqBrokeRuntimeSendThreadPoolQueueHeadWaitTimeMills
	ch <- rocketmqBrokeRuntimeMsgGetTotalYesterdayMorning
	ch <- rocketmqBrokeRuntimeMsgPutTotalYesterdayMorning
	ch <- rocketmqBrokeRuntimeMsgGetTotalTodayMorning
	ch <- rocketmqBrokeRuntimeMsgPutTotalTodayMorning
	ch <- rocketmqBrokeRuntimeCommitLogDirCapacityFree
	ch <- rocketmqBrokeRuntimeCommitLogDirCapacityTotal
	ch <- rocketmqBrokeRuntimeCommitLogMaxOffset
	ch <- rocketmqBrokeRuntimeCommitLogMinOffset
	ch <- rocketmqBrokeRuntimeRemainHowManyDataToFlush
}
