# rocketmq-exporter-go
An application of Rocketmq exporter for Prometheus which purely developed by golang

Table of Contents
-----------------
- [rocketmq-exporter-go](#rocketmq-exporter-go)
  - [Table of Contents](#table-of-contents)
  - [Note](#note)
  - [Dependency](#dependency)
  - [Compile](#compile)
    - [Build](#build)
    - [Build For Linux on Windows](#build-for-linux-on-windows)
  - [Run](#run)
    - [Run Directly](#run-directly)
    - [Docker build](#docker-build)
    - [Docker Run](#docker-run)
  - [Flags](#flags)
  - [Metrics](#metrics)
    - [Producer](#producer)
    - [Consumer](#consumer)
    - [Consumer Group](#consumer-group)
    - [Broker](#broker)
    - [Broker Runtime](#broker-runtime)
  - [Grafana Dashboard](#grafana-dashboard)
  - [Contribute](#contribute)
  - [Contributors ✨](#contributors-)
  - [Donation](#donation)
  - [License](#license)

Note
----------
- We copy the rocketmq-client-go some module code into this project directly like internal, admin, consumer. The reason is that we can go quickly.
- If the suffix of the filename is ext, those file were extended by ourself in order to distinguish which is the source code.
- The admin has been extended for meeting to fetch some metric data. We should let this code come back to the resource project in the feature.
- However i will commit thees extension code to the source code as soon as i can

Dependency
----------

-  [Prometheus](https://prometheus.io)
-  [Rocketmq-Client](https://github.com/apache/rocketmq-client-go)
-  [Golang](https://golang.org)

Compile
-------

### Build
```shell
go build
```

### Build For Linux on Windows
```shell
# into cmd
set GOOS=linux
set GOARCH=amd64
go build
```

Run
---

### Run Directly
```shell
go run ./main.go --rocketmq.nameserver=[nameserverIp]
```

### Docker build
use docker buildx
```shell
mkdir [linux-GOARCH(had been replace actually value)]
# and the put the go file which has been built into the folder
docker buildx build --platform [platform] -t [tag] . [--push]
```
> Note: if you meet **error: multiple platforms feature is currently not supported for docker driver. Please switch to a different driver (eg. "docker buildx create --use")** , you should create a new [driver](https://docs.docker.com/engine/reference/commandline/buildx_create/)

### Docker Run
```shell
docker run -it --rm -p 9999:9999 [imageTag] --rocketmq.nameserver=[nameserverIp]
```


Flags
-----
This image is configurable using different flags

| Flag name | Default | Description|                 
| --------- | ------- | ---------- |
| rocketmq.nameserver.use.domain | false | Address of rocketmq nameserver whether use domian |
| rocketmq.nameserver | 127.0.0.1:9876 | Addresses (ip:port) of rocketmq nameserver |
| rocketmq.nameserver.domain | rocketmq:9876 | Addresses (domain:port) of rocketmq nameserver |
| workers | 100 | Number of workers |
| web.listen-address | :9999 | Address to listen on for web interface and telemetry |
| web.telemetry-path | /metrics | Path under which to expose metrics |
| log.file.enable | false | Log write to file enable |
| log.file.path | ./logs/go.log | Path of og file |


Metrics
-------

### Producer

| Name | Help |
| ---- | ---- |
| rocketmq_producer_tps | TopicPutNums |
| rocketmq_producer_offset | TopicOffset |
| rocketmq_producer_message_size | TopicPutMessageSize |

### Consumer

| Name | Help |
| ---- | ---- |
| rocketmq_topic_retry_offset | TopicRetryOffset |
| rocketmq_topic_dlq_offset | TopicRetryOffset |
| rocketmq_client_consume_rt | consumerClientRT |
| rocketmq_client_consumer_pull_rt | consumerClientPullRT |
| rocketmq_client_consumer_pull_tps | consumerClientPullTPS |
| rocketmq_client_consume_ok_msg_tps | consumerClientOKTPS |
| rocketmq_client_consume_fail_msg_count | consumerClientFailedMsgCounts |
| rocketmq_client_consume_fail_msg_tps | consumerClientFailedTPS |
| rocketmq_send_back_nums | SendBackNums |

### Consumer Group

| Name | Help |
| ---- | ---- |
| rocketmq_consumer_tps | GroupGetNums |
| rocketmq_consumer_offset | GroupBrokerTotalOffset |
| rocketmq_consumer_message_size | GroupGetMessageSize |
| rocketmq_group_consume_tps | GroupConsumeTPS |
| rocketmq_group_consume_total_offset | GroupConsumeTotalOffset |
| rocketmq_group_get_latency_by_storetime | GroupGetLatencyByStoreTime |
| rocketmq_group_count | GroupCount |
| rocketmq_group_diff | GroupDiff |
| rocketmq_group_retrydiff | GroupRetryDiff |
| rocketmq_group_dlqdiff | GroupDLQDiff |

### Broker

| Name | Help |
| ---- | ---- |
| rocketmq_broker_tps | BrokerPutNums |
| rocketmq_broker_qps | BrokerGetNums |

### Broker Runtime

| Name | Help |
| ---- | ---- |
| rocketmq_brokeruntime_pmdt_0ms | PutMessageDistributeTimeMap0ms |
| rocketmq_brokeruntime_pmdt_0to10ms | PutMessageDistributeTimeMap0to10ms |
| rocketmq_brokeruntime_pmdt_10to50ms | PutMessageDistributeTimeMap10to50ms |
| rocketmq_brokeruntime_pmdt_50to100ms | PutMessageDistributeTimeMap50to100ms |
| rocketmq_brokeruntime_pmdt_100to200ms | PutMessageDistributeTimeMap100to200ms |
| rocketmq_brokeruntime_pmdt_200to500ms | PutMessageDistributeTimeMap200to500ms |
| rocketmq_brokeruntime_pmdt_500to1s | PutMessageDistributeTimeMap500to1s |
| rocketmq_brokeruntime_pmdt_1to2s | PutMessageDistributeTimeMap1to2s |
| rocketmq_brokeruntime_pmdt_2to3s | PutMessageDistributeTimeMap2to3s |
| rocketmq_brokeruntime_pmdt_3to4s | PutMessageDistributeTimeMap3to4s |
| rocketmq_brokeruntime_pmdt_4to5s | PutMessageDistributeTimeMap4to5s |
| rocketmq_brokeruntime_pmdt_5to10s | PutMessageDistributeTimeMap5to10s |
| rocketmq_brokeruntime_pmdt_10stomore | PutMessageDistributeTimeMap10toMore |
| rocketmq_brokeruntime_consumequeue_disk_ratio | brokerRuntimeConsumeQueueDiskRatio |
| rocketmq_brokeruntime_getfound_tps600 | brokerRuntimeGetFoundTps600 |
| rocketmq_brokeruntime_getfound_tps60 | brokerRuntimeGetFoundTps60 |
| rocketmq_brokeruntime_getfound_tps10 | brokerRuntimeGetFoundTps10 |
| rocketmq_brokeruntime_gettransfered_tps600 | brokerRuntimeGetTransferedTps600 |
| rocketmq_brokeruntime_gettransfered_tps60 | brokerRuntimeGetTransferedTps60 |
| rocketmq_brokeruntime_gettransfered_tps10 | brokerRuntimeGetTransferedTps10 |
| rocketmq_brokeruntime_getmiss_tps600 | brokerRuntimeGetMissTps600 |
| rocketmq_brokeruntime_getmiss_tps60 | brokerRuntimeGetMissTps60 |
| rocketmq_brokeruntime_getmiss_tps10 | brokerRuntimeGetMissTps10 |
| rocketmq_brokeruntime_gettotal_tps600 | brokerRuntimeGetTotalTps600 |
| rocketmq_brokeruntime_gettotal_tps60 | brokerRuntimeGetTotalTps60 |
| rocketmq_brokeruntime_gettotal_tps10 | brokerRuntimeGetTotalTps10 |
| rocketmq_brokeruntime_put_tps600 | brokerRuntimePutTps600 |
| rocketmq_brokeruntime_put_tps60 | brokerRuntimePutTps60 |
| rocketmq_brokeruntime_put_tps10 | brokerRuntimePutTps10 |
| rocketmq_brokeruntime_put_latency_999 | brokerRuntimePutLatency99 |
| rocketmq_brokeruntime_put_latency_9999 | brokerRuntimePutLatency999 |
| rocketmq_brokeruntime_dispatch_maxbuffer | brokerRuntimeDispatchMaxBuffer |
| rocketmq_brokeruntime_dispatch_behind_bytes | brokerRuntimeDispatchBehindBytes |
| rocketmq_brokeruntime_put_message_size_total | brokerRuntimePutMessageSizeTotal |
| rocketmq_brokeruntime_put_message_average_size | brokerRuntimePutMessageAverageSize |
| rocketmq_brokeruntime_remain_transientstore_buffer_numbs | brokerRuntimeRemainTransientStoreBufferNumbs |
| rocketmq_brokeruntime_start_accept_sendrequest_time | brokerRuntimeStartAcceptSendRequestTimeStamp |
| rocketmq_brokeruntime_earliest_message_timestamp | brokerRuntimeEarliestMessageTimeStamp |
| rocketmq_brokeruntime_getmessage_entire_time_max | brokerRuntimeGetMessageEntireTimeMax |
| rocketmq_brokeruntime_putmessage_entire_time_max | brokerRuntimePutMessageEntireTimeMax |
| rocketmq_brokeruntime_putmessage_times_total | brokerRuntimePutMessageTimesTotal |
| rocketmq_brokeruntime_pagecache_lock_time_mills | brokerRuntimePageCacheLockTimeMills |
| rocketmq_brokeruntime_commitlog_minoffset | brokerRuntimeCommitLogMinOffset |
| rocketmq_brokeruntime_commitlog_maxoffset | brokerRuntimeCommitLogMaxOffset |
| rocketmq_brokeruntime_commitlog_disk_ratio | brokerRuntimeCommitLogDiskRatio |
| rocketmq_brokeruntime_pull_threadpoolqueue_capacity | brokerRuntimePullThreadPoolQueueCapacity |
| rocketmq_brokeruntime_pull_threadpoolqueue_size | brokerRuntimePullThreadPoolQueueSizeF |
| rocketmq_brokeruntime_send_threadpoolqueue_capacity | brokerRuntimeSendThreadPoolQueueCapacity |
| rocketmq_brokeruntime_send_threadpool_queue_size | brokerRuntimeSendThreadPoolQueueSize |
| rocketmq_brokeruntime_query_threadpool_queue_capacity | brokerRuntimeQueryThreadPoolQueueCapacity |
| rocketmq_brokeruntime_query_threadpoolqueue_size | brokerRuntimeQueryThreadPoolQueueSize |
| rocketmq_brokeruntime_pull_threadpoolqueue_headwait_timemills | brokerRuntimePullThreadPoolQueueHeadWaitTimeMills |
| rocketmq_brokeruntime_send_threadpoolqueue_headwait_timemills | brokerRuntimeSendThreadPoolQueueHeadWaitTimeMills |
| rocketmq_brokeruntime_query_threadpoolqueue_headwait_timemills | brokerRuntimeQueryThreadPoolQueueHeadWaitTimeMills |
| rocketmq_brokeruntime_msg_put_total_today_now | brokerRuntimeMsgPutTotalTodayNow |
| rocketmq_brokeruntime_msg_gettotal_today_now | brokerRuntimeMsgGetTotalTodayNow |
| rocketmq_brokeruntime_msg_gettotal_yesterdaymorning | brokerRuntimeMsgGetTotalYesterdayMorning |
| rocketmq_brokeruntime_msg_puttotal_yesterdaymorning | brokerRuntimeMsgPutTotalYesterdayMorning |
| rocketmq_brokeruntime_msg_gettotal_todaymorning | brokerRuntimeMsgGetTotalTodayMorning |
| rocketmq_brokeruntime_msg_puttotal_todaymorning | brokerRuntimeMsgPutTotalTodayMorning |
| rocketmq_brokeruntime_commitlogdir_capacity_free | brokerRuntimeCommitLogDirCapacityFree |
| rocketmq_brokeruntime_commitlogdir_capacity_total | brokerRuntimeCommitLogDirCapacityTotal |
| rocketmq_brokeruntime_remain_howmanydata_toflush | brokerRuntimeRemainHowManyDataToFlush |


Grafana Dashboard
-------
| Grafana Dashboard ID | Name | Detail |
|----------------------|------------------|------------------|
| 10477 | Rocketmq_dashboard | [Rocketmq_dashboard](https://grafana.com/grafana/dashboards/10477) |
| 14612 | RocketMQ监控 | [RocketMQ监控](https://grafana.com/grafana/dashboards/14612) |

Contribute
------------

If you like Rocketmq Exporter, please give me a star. This will help more people know Rocketmq Exporter.

Please feel free to send me [pull requests](https://github.com/liuymcn/rocketmq-exporter-go/pulls).

Task List

- improve the document such as metric descruption
- the guide of the using metric to judge whether the system is healthy
- the consumer is not completely, have to be updated
- fix some json serialization logic
- fix some strange logic
- some exception has not been handled
- add test unit
- add ACL feature
- add makefile
- add kubernate chart

Contributors ✨
----------

Thanks goes to these wonderful people:

<a href="https://github.com/liuymcn/rocketmq-exporter-go/graphs/contributors"> come on! </a>

Donation
--------

Your donation will encourage me to continue to improve Rocketmq Exporter. Support Alipay donation.

![](https://github.com/liuymcn/rocketmq-exporter-go/raw/main/alipay.png)

License
-------

Code is licensed under the [Apache License 2.0](https://github.com/liuymcn/rocketmq-exporter-go/blob/main/LICENSE).
