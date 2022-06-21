package exporter

import "github.com/prometheus/client_golang/prometheus"

var (
	rocketmqGroupDiff                                         *prometheus.Desc
	rocketmqGroupRetryDiff                                    *prometheus.Desc
	rocketmqGroupDlqDiff                                      *prometheus.Desc
	rocketmqGroupCount                                        *prometheus.Desc
	rocketmqProducerOffset                                    *prometheus.Desc
	rocketmqTopicRetryOffset                                  *prometheus.Desc
	rocketmqTopicDlqOffset                                    *prometheus.Desc
	rocketmqClientConsumeFailMsgCount                         *prometheus.Desc
	rocketmqClientConsumeFailMsgTps                           *prometheus.Desc
	rocketmqClientConsumeOkMsgTps                             *prometheus.Desc
	rocketmqClientConsumeRt                                   *prometheus.Desc
	rocketmqClientConsumerPullRt                              *prometheus.Desc
	rocketmqClientConsumerPullTps                             *prometheus.Desc
	rocketmqBrokerTps                                         *prometheus.Desc
	rocketmqBrokerQps                                         *prometheus.Desc
	rocketmqBrokeRuntimePmdt0ms                               *prometheus.Desc
	rocketmqBrokeRuntimePmdt0to10ms                           *prometheus.Desc
	rocketmqBrokeRuntimePmdt10to50ms                          *prometheus.Desc
	rocketmqBrokeRuntimePmdt50to100ms                         *prometheus.Desc
	rocketmqBrokeRuntimePmdt100to200ms                        *prometheus.Desc
	rocketmqBrokeRuntimePmdt200to500ms                        *prometheus.Desc
	rocketmqBrokeRuntimePmdt500to1s                           *prometheus.Desc
	rocketmqBrokeRuntimePmdt1to2s                             *prometheus.Desc
	rocketmqBrokeRuntimePmdt2to3s                             *prometheus.Desc
	rocketmqBrokeRuntimePmdt3to4s                             *prometheus.Desc
	rocketmqBrokeRuntimePmdt4to5s                             *prometheus.Desc
	rocketmqBrokeRuntimePmdt5to10s                            *prometheus.Desc
	rocketmqBrokeRuntimePmdt10stomore                         *prometheus.Desc
	rocketmqConsumerTps                                       *prometheus.Desc
	rocketmqGroupConsumeTps                                   *prometheus.Desc
	rocketmqConsumerOffset                                    *prometheus.Desc
	rocketmqGroupConsumeTotalOffset                           *prometheus.Desc
	rocketmqConsumerMessageSize                               *prometheus.Desc
	rocketmqSendBackNums                                      *prometheus.Desc
	rocketmqGroupGetLatencyByStoreTime                        *prometheus.Desc
	rocketmqProducerTps                                       *prometheus.Desc
	rocketmqProducerMessageSize                               *prometheus.Desc
	rocketmqBrokeRuntimeMsgPutTotalTodayNow                   *prometheus.Desc
	rocketmqBrokeRuntimeMsgGetTotalTodayNow                   *prometheus.Desc
	rocketmqBrokeRuntimeDispatchBehindBytes                   *prometheus.Desc
	rocketmqBrokeRuntimePutMessageSizeTotal                   *prometheus.Desc
	rocketmqBrokeRuntimePutMessageAverageSize                 *prometheus.Desc
	rocketmqBrokeRuntimeQueryThreadPoolQueueCapacity          *prometheus.Desc
	rocketmqBrokeRuntimeRemainTransientStoreBufferNumbs       *prometheus.Desc
	rocketmqBrokeRuntimeEarliestMessageTimestamp              *prometheus.Desc
	rocketmqBrokeRuntimePutMessageEntireTimeMax               *prometheus.Desc
	rocketmqBrokeRuntimeStartAcceptSendrequestTime            *prometheus.Desc
	rocketmqBrokeRuntimeSendThreadPoolQueueSize               *prometheus.Desc
	rocketmqBrokeRuntimePutMessageTimesTotal                  *prometheus.Desc
	rocketmqBrokeRuntimeGetMessageEntireTimeMax               *prometheus.Desc
	rocketmqBrokeRuntimePageCacheLockTimeMills                *prometheus.Desc
	rocketmqBrokeRuntimeCommitLogDiskRatio                    *prometheus.Desc
	rocketmqBrokeRuntimeConsumeQueueDiskRatio                 *prometheus.Desc
	rocketmqBrokeRuntimeGetFoundTps600                        *prometheus.Desc
	rocketmqBrokeRuntimeGetFoundTps60                         *prometheus.Desc
	rocketmqBrokeRuntimeGetFoundTps10                         *prometheus.Desc
	rocketmqBrokeRuntimeGetTotalTps600                        *prometheus.Desc
	rocketmqBrokeRuntimeGetTotalTps60                         *prometheus.Desc
	rocketmqBrokeRuntimeGetTotalTps10                         *prometheus.Desc
	rocketmqBrokeRuntimeGetTransferedTps600                   *prometheus.Desc
	rocketmqBrokeRuntimeGetTransferedTps60                    *prometheus.Desc
	rocketmqBrokeRuntimeGetTransferedTps10                    *prometheus.Desc
	rocketmqBrokeRuntimeGetMissTps600                         *prometheus.Desc
	rocketmqBrokeRuntimeGetMissTps60                          *prometheus.Desc
	rocketmqBrokeRuntimeGetMissTps10                          *prometheus.Desc
	rocketmqBrokeRuntimePutTps600                             *prometheus.Desc
	rocketmqBrokeRuntimePutTps60                              *prometheus.Desc
	rocketmqBrokeRuntimePutTps10                              *prometheus.Desc
	rocketmqBrokeRuntimeDispatchMaxBuffer                     *prometheus.Desc
	rocketmqBrokeRuntimePullThreadPoolQueueCapacity           *prometheus.Desc
	rocketmqBrokeRuntimeSendThreadPoolQueueCapacity           *prometheus.Desc
	rocketmqBrokeRuntimePullThreadPoolQueueSize               *prometheus.Desc
	rocketmqBrokeRuntimeQueryThreadPoolQueueSize              *prometheus.Desc
	rocketmqBrokeRuntimePullThreadPoolQueueHeadWaitTimeMills  *prometheus.Desc
	rocketmqBrokeRuntimeQueryThreadPoolQueueHeadWaitTimeMills *prometheus.Desc
	rocketmqBrokeRuntimeSendThreadPoolQueueHeadWaitTimeMills  *prometheus.Desc
	rocketmqBrokeRuntimeMsgGetTotalYesterdayMorning           *prometheus.Desc
	rocketmqBrokeRuntimeMsgPutTotalYesterdayMorning           *prometheus.Desc
	rocketmqBrokeRuntimeMsgGetTotalTodayMorning               *prometheus.Desc
	rocketmqBrokeRuntimeMsgPutTotalTodayMorning               *prometheus.Desc
	rocketmqBrokeRuntimeCommitLogDirCapacityFree              *prometheus.Desc
	rocketmqBrokeRuntimeCommitLogDirCapacityTotal             *prometheus.Desc
	rocketmqBrokeRuntimeCommitLogMaxOffset                    *prometheus.Desc
	rocketmqBrokeRuntimeCommitLogMinOffset                    *prometheus.Desc
	rocketmqBrokeRuntimeRemainHowManyDataToFlush              *prometheus.Desc
)

func InitMetricDesc() {

	var groupDiffLabelNames = []string{"group", "topic", "countOfOnlineConsumers", "msgModel"}
	var groupCountLabelNames = []string{"caddr", "localaddr", "topic", "group"}
	var topicOffsetLabelNames = []string{"cluster", "broker", "topic", "lastUpdateTimestamp"}
	var dlqTopicOffsetLabelNames = []string{"cluster", "broker", "group", "lastUpdateTimestamp"}
	var groupClientMetricLabelNames = []string{"group", "topic", "clientAddr", "clientId"}
	var groupLatencyByStoretimeLabelNames = []string{"cluster", "broker", "topic", "group"}
	var brokerNumsLabelNames = []string{"cluster", "broker", "brokerIP"}
	var groupNumsLabelNames = []string{"cluster", "broker", "topic", "group"}
	var topicNumsLabelNames = []string{"cluster", "broker", "topic"}
	var brokerRuntimeMetricLabelNames = []string{"cluster", "brokerIP", "brokerVersion", "brokerVersionDes", "bootTime"}

	rocketmqGroupDiff = prometheus.NewDesc(
		"rocketmq_group_diff",
		"GroupDiff",
		groupDiffLabelNames,
		nil,
	)
	rocketmqGroupRetryDiff = prometheus.NewDesc(
		"rocketmq_group_retrydiff",
		"GroupRetryDiff",
		groupDiffLabelNames,
		nil,
	)
	rocketmqGroupDlqDiff = prometheus.NewDesc(
		"rocketmq_group_dlqdiff",
		"GroupDLQDiff",
		groupDiffLabelNames,
		nil,
	)
	rocketmqGroupCount = prometheus.NewDesc(
		"rocketmq_group_count",
		"GroupCount",
		groupCountLabelNames,
		nil,
	)
	rocketmqProducerOffset = prometheus.NewDesc(
		"rocketmq_producer_offset",
		"TopicOffset",
		topicOffsetLabelNames,
		nil,
	)
	rocketmqTopicRetryOffset = prometheus.NewDesc(
		"rocketmq_topic_retry_offset",
		"TopicRetryOffset",
		topicOffsetLabelNames,
		nil,
	)
	rocketmqTopicDlqOffset = prometheus.NewDesc(
		"rocketmq_topic_dlq_offset",
		"TopicRetryOffset",
		dlqTopicOffsetLabelNames,
		nil,
	)
	rocketmqClientConsumeFailMsgCount = prometheus.NewDesc(
		"rocketmq_client_consume_fail_msg_count",
		"consumerClientFailedMsgCounts",
		groupClientMetricLabelNames,
		nil,
	)
	rocketmqClientConsumeFailMsgTps = prometheus.NewDesc(
		"rocketmq_client_consume_fail_msg_tps",
		"consumerClientFailedTPS",
		groupClientMetricLabelNames,
		nil,
	)
	rocketmqClientConsumeOkMsgTps = prometheus.NewDesc(
		"rocketmq_client_consume_ok_msg_tps",
		"consumerClientOKTPS",
		groupClientMetricLabelNames,
		nil,
	)
	rocketmqClientConsumeRt = prometheus.NewDesc(
		"rocketmq_client_consume_rt",
		"consumerClientRT",
		groupClientMetricLabelNames,
		nil,
	)
	rocketmqClientConsumerPullRt = prometheus.NewDesc(
		"rocketmq_client_consumer_pull_rt",
		"consumerClientPullRT",
		groupClientMetricLabelNames,
		nil,
	)
	rocketmqClientConsumerPullTps = prometheus.NewDesc(
		"rocketmq_client_consumer_pull_tps",
		"consumerClientPullTPS",
		groupClientMetricLabelNames,
		nil,
	)
	rocketmqBrokerTps = prometheus.NewDesc(
		"rocketmq_broker_tps",
		"BrokerPutNums",
		brokerNumsLabelNames,
		nil,
	)
	rocketmqBrokerQps = prometheus.NewDesc(
		"rocketmq_broker_qps",
		"BrokerGetNums",
		brokerNumsLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt0ms = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_0ms",
		"PutMessageDistributeTimeMap0ms",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt0to10ms = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_0to10ms",
		"PutMessageDistributeTimeMap0to10ms",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt10to50ms = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_10to50ms",
		"PutMessageDistributeTimeMap10to50ms",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt50to100ms = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_50to100ms",
		"PutMessageDistributeTimeMap50to100ms",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt100to200ms = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_100to200ms",
		"PutMessageDistributeTimeMap100to200ms",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt200to500ms = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_200to500ms",
		"PutMessageDistributeTimeMap200to500ms",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt500to1s = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_500to1s",
		"PutMessageDistributeTimeMap500to1s",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt1to2s = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_1to2s",
		"PutMessageDistributeTimeMap1to2s",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt2to3s = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_2to3s",
		"PutMessageDistributeTimeMap2to3s",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt3to4s = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_3to4s",
		"PutMessageDistributeTimeMap3to4s",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt4to5s = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_4to5s",
		"PutMessageDistributeTimeMap4to5s",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt5to10s = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_5to10s",
		"PutMessageDistributeTimeMap5to10s",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePmdt10stomore = prometheus.NewDesc(
		"rocketmq_brokeruntime_pmdt_10stomore",
		"PutMessageDistributeTimeMap10toMore",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqConsumerTps = prometheus.NewDesc(
		"rocketmq_consumer_tps",
		"GroupGetNums",
		groupNumsLabelNames,
		nil,
	)
	rocketmqGroupConsumeTps = prometheus.NewDesc(
		"rocketmq_group_consume_tps",
		"GroupConsumeTPS",
		groupNumsLabelNames,
		nil,
	)
	rocketmqConsumerOffset = prometheus.NewDesc(
		"rocketmq_consumer_offset",
		"GroupBrokerTotalOffset",
		groupNumsLabelNames,
		nil,
	)
	rocketmqGroupConsumeTotalOffset = prometheus.NewDesc(
		"rocketmq_group_consume_total_offset",
		"GroupConsumeTotalOffset",
		groupNumsLabelNames,
		nil,
	)
	rocketmqConsumerMessageSize = prometheus.NewDesc(
		"rocketmq_consumer_message_size",
		"GroupGetMessageSize",
		groupNumsLabelNames,
		nil,
	)
	rocketmqSendBackNums = prometheus.NewDesc(
		"rocketmq_send_back_nums",
		"SendBackNums",
		groupNumsLabelNames,
		nil,
	)
	rocketmqGroupGetLatencyByStoreTime = prometheus.NewDesc(
		"rocketmq_group_get_latency_by_storetime",
		"GroupGetLatencyByStoreTime",
		groupLatencyByStoretimeLabelNames,
		nil,
	)
	rocketmqProducerTps = prometheus.NewDesc(
		"rocketmq_producer_tps",
		"TopicPutNums",
		topicNumsLabelNames,
		nil,
	)
	rocketmqProducerMessageSize = prometheus.NewDesc(
		"rocketmq_producer_message_size",
		"TopicPutMessageSize",
		topicNumsLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeMsgPutTotalTodayNow = prometheus.NewDesc(
		"rocketmq_brokeruntime_msg_put_total_today_now",
		"brokerRuntimeMsgPutTotalTodayNow",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeMsgGetTotalTodayNow = prometheus.NewDesc(
		"rocketmq_brokeruntime_msg_gettotal_today_now",
		"brokerRuntimeMsgGetTotalTodayNow",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeDispatchBehindBytes = prometheus.NewDesc(
		"rocketmq_brokeruntime_dispatch_behind_bytes",
		"brokerRuntimeDispatchBehindBytes",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePutMessageSizeTotal = prometheus.NewDesc(
		"rocketmq_brokeruntime_put_message_size_total",
		"brokerRuntimePutMessageSizeTotal",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePutMessageAverageSize = prometheus.NewDesc(
		"rocketmq_brokeruntime_put_message_average_size",
		"brokerRuntimePutMessageAverageSize",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeQueryThreadPoolQueueCapacity = prometheus.NewDesc(
		"rocketmq_brokeruntime_query_threadpool_queue_capacity",
		"brokerRuntimeQueryThreadPoolQueueCapacity",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeRemainTransientStoreBufferNumbs = prometheus.NewDesc(
		"rocketmq_brokeruntime_remain_transientstore_buffer_numbs",
		"brokerRuntimeRemainTransientStoreBufferNumbs",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeEarliestMessageTimestamp = prometheus.NewDesc(
		"rocketmq_brokeruntime_earliest_message_timestamp",
		"brokerRuntimeEarliestMessageTimeStamp",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePutMessageEntireTimeMax = prometheus.NewDesc(
		"rocketmq_brokeruntime_putmessage_entire_time_max",
		"brokerRuntimePutMessageEntireTimeMax",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeStartAcceptSendrequestTime = prometheus.NewDesc(
		"rocketmq_brokeruntime_start_accept_sendrequest_time",
		"brokerRuntimeStartAcceptSendRequestTimeStamp",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeSendThreadPoolQueueSize = prometheus.NewDesc(
		"rocketmq_brokeruntime_send_threadpool_queue_size",
		"brokerRuntimeSendThreadPoolQueueSize",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePutMessageTimesTotal = prometheus.NewDesc(
		"rocketmq_brokeruntime_putmessage_times_total",
		"brokerRuntimePutMessageTimesTotal",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetMessageEntireTimeMax = prometheus.NewDesc(
		"rocketmq_brokeruntime_getmessage_entire_time_max",
		"brokerRuntimeGetMessageEntireTimeMax",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePageCacheLockTimeMills = prometheus.NewDesc(
		"rocketmq_brokeruntime_pagecache_lock_time_mills",
		"brokerRuntimePageCacheLockTimeMills",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeCommitLogDiskRatio = prometheus.NewDesc(
		"rocketmq_brokeruntime_commitlog_disk_ratio",
		"brokerRuntimeCommitLogDiskRatio",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeConsumeQueueDiskRatio = prometheus.NewDesc(
		"rocketmq_brokeruntime_consumequeue_disk_ratio",
		"brokerRuntimeConsumeQueueDiskRatio",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetFoundTps600 = prometheus.NewDesc(
		"rocketmq_brokeruntime_getfound_tps600",
		"brokerRuntimeGetFoundTps600",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetFoundTps60 = prometheus.NewDesc(
		"rocketmq_brokeruntime_getfound_tps60",
		"brokerRuntimeGetFoundTps60",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetFoundTps10 = prometheus.NewDesc(
		"rocketmq_brokeruntime_getfound_tps10",
		"brokerRuntimeGetFoundTps10",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetTotalTps600 = prometheus.NewDesc(
		"rocketmq_brokeruntime_gettotal_tps600",
		"brokerRuntimeGetTotalTps600",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetTotalTps60 = prometheus.NewDesc(
		"rocketmq_brokeruntime_gettotal_tps60",
		"brokerRuntimeGetTotalTps60",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetTotalTps10 = prometheus.NewDesc(
		"rocketmq_brokeruntime_gettotal_tps10",
		"brokerRuntimeGetTotalTps10",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetTransferedTps600 = prometheus.NewDesc(
		"rocketmq_brokeruntime_gettransfered_tps600",
		"brokerRuntimeGetTransferedTps600",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetTransferedTps60 = prometheus.NewDesc(
		"rocketmq_brokeruntime_gettransfered_tps60",
		"brokerRuntimeGetTransferedTps60",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetTransferedTps10 = prometheus.NewDesc(
		"rocketmq_brokeruntime_gettransfered_tps10",
		"brokerRuntimeGetTransferedTps10",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetMissTps600 = prometheus.NewDesc(
		"rocketmq_brokeruntime_getmiss_tps600",
		"brokerRuntimeGetMissTps600",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetMissTps60 = prometheus.NewDesc(
		"rocketmq_brokeruntime_getmiss_tps60",
		"brokerRuntimeGetMissTps60",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeGetMissTps10 = prometheus.NewDesc(
		"rocketmq_brokeruntime_getmiss_tps10",
		"brokerRuntimeGetMissTps10",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePutTps600 = prometheus.NewDesc(
		"rocketmq_brokeruntime_put_tps600",
		"brokerRuntimePutTps600",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePutTps60 = prometheus.NewDesc(
		"rocketmq_brokeruntime_put_tps60",
		"brokerRuntimePutTps60",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePutTps10 = prometheus.NewDesc(
		"rocketmq_brokeruntime_put_tps10",
		"brokerRuntimePutTps10",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeDispatchMaxBuffer = prometheus.NewDesc(
		"rocketmq_brokeruntime_dispatch_maxbuffer",
		"brokerRuntimeDispatchMaxBuffer",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePullThreadPoolQueueCapacity = prometheus.NewDesc(
		"rocketmq_brokeruntime_pull_threadpoolqueue_capacity",
		"brokerRuntimePullThreadPoolQueueCapacity",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeSendThreadPoolQueueCapacity = prometheus.NewDesc(
		"rocketmq_brokeruntime_send_threadpoolqueue_capacity",
		"brokerRuntimeSendThreadPoolQueueCapacity",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePullThreadPoolQueueSize = prometheus.NewDesc(
		"rocketmq_brokeruntime_pull_threadpoolqueue_size",
		"brokerRuntimePullThreadPoolQueueSizeF",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeQueryThreadPoolQueueSize = prometheus.NewDesc(
		"rocketmq_brokeruntime_query_threadpoolqueue_size",
		"brokerRuntimeQueryThreadPoolQueueSize",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimePullThreadPoolQueueHeadWaitTimeMills = prometheus.NewDesc(
		"rocketmq_brokeruntime_pull_threadpoolqueue_headwait_timemills",
		"brokerRuntimePullThreadPoolQueueHeadWaitTimeMills",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeQueryThreadPoolQueueHeadWaitTimeMills = prometheus.NewDesc(
		"rocketmq_brokeruntime_query_threadpoolqueue_headwait_timemills",
		"brokerRuntimeQueryThreadPoolQueueHeadWaitTimeMills",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeSendThreadPoolQueueHeadWaitTimeMills = prometheus.NewDesc(
		"rocketmq_brokeruntime_send_threadpoolqueue_headwait_timemills",
		"brokerRuntimeSendThreadPoolQueueHeadWaitTimeMills",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeMsgGetTotalYesterdayMorning = prometheus.NewDesc(
		"rocketmq_brokeruntime_msg_gettotal_yesterdaymorning",
		"brokerRuntimeMsgGetTotalYesterdayMorning",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeMsgPutTotalYesterdayMorning = prometheus.NewDesc(
		"rocketmq_brokeruntime_msg_puttotal_yesterdaymorning",
		"brokerRuntimeMsgPutTotalYesterdayMorning",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeMsgGetTotalTodayMorning = prometheus.NewDesc(
		"rocketmq_brokeruntime_msg_gettotal_todaymorning",
		"brokerRuntimeMsgGetTotalTodayMorning",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeMsgPutTotalTodayMorning = prometheus.NewDesc(
		"rocketmq_brokeruntime_msg_puttotal_todaymorning",
		"brokerRuntimeMsgPutTotalTodayMorning",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeCommitLogDirCapacityFree = prometheus.NewDesc(
		"rocketmq_brokeruntime_commitlogdir_capacity_free",
		"brokerRuntimeCommitLogDirCapacityFree",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeCommitLogDirCapacityTotal = prometheus.NewDesc(
		"rocketmq_brokeruntime_commitlogdir_capacity_total",
		"brokerRuntimeCommitLogDirCapacityTotal",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeCommitLogMaxOffset = prometheus.NewDesc(
		"rocketmq_brokeruntime_commitlog_maxoffset",
		"brokerRuntimeCommitLogMaxOffset",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeCommitLogMinOffset = prometheus.NewDesc(
		"rocketmq_brokeruntime_commitlog_minoffset",
		"brokerRuntimeCommitLogMinOffset",
		brokerRuntimeMetricLabelNames,
		nil,
	)
	rocketmqBrokeRuntimeRemainHowManyDataToFlush = prometheus.NewDesc(
		"rocketmq_brokeruntime_remain_howmanydata_toflush",
		"brokerRuntimeRemainHowManyDataToFlush",
		brokerRuntimeMetricLabelNames,
		nil,
	)
}
