package internal

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/rocketmq-exporter-go/internal/remote"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/robertkrimen/otto"
)

type NamesrvsExt interface {
	Namesrvs

	QueryBrokerClusterInfoFromServer() (*ClusterInfo, error)

	ListTopic() ([]string, error)

	QueryTopicRouteInfoFromServer(topic string) (*TopicRouteData, error)
}

type namesrvsExt struct {
	*namesrvs
}

func NewNamesrvExt(resolver primitive.NsResolver) (*namesrvsExt, error) {
	addr := resolver.Resolve()
	if len(addr) == 0 {
		return nil, errors.New("no name server addr found with resolver: " + resolver.Description())
	}

	if err := primitive.NamesrvAddr(addr).Check(); err != nil {
		return nil, err
	}
	nameSrvClient := remote.NewRemotingClient()
	return &namesrvsExt{
		&namesrvs{
			srvs:             addr,
			lock:             new(sync.Mutex),
			nameSrvClient:    nameSrvClient,
			brokerVersionMap: make(map[string]map[string]int32, 0),
			brokerLock:       new(sync.RWMutex),
			resolver:         resolver,
		},
	}, nil
}

type ClusterInfo struct {
	// key -> brokerName
	BrokerDataTable map[string]*BrokerData `json:"brokerAddrTable"`

	// key -> clusterName, value -> list of the brokerName
	ClusterTable map[string][]string `json:"clusterAddrTable"`
}

func (ci *ClusterInfo) decode(data []byte) error {

	vm := otto.New()
	vm.Set("source", string(data))
	value, err := vm.Run(`
	    var code = 'JSON.stringify(' + source + ')';
		eval(code);
	`)
	if err != nil {
		return err
	}
	result, _ := value.ToString()

	return json.Unmarshal([]byte(result), ci)

}

func (namesrv *namesrvsExt) QueryBrokerClusterInfoFromServer() (*ClusterInfo, error) {

	var (
		response *remote.RemotingCommand
		err      error
	)

	cmd := remote.NewRemotingCommand(ReqGetBrokerClusterInfo, nil, nil)
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	response, err = namesrv.nameSrvClient.InvokeSync(ctx, namesrv.getNameServerAddress(), cmd)

	if err != nil {
		return nil, primitive.NewRemotingErr(err.Error())
	}

	switch response.Code {
	case ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		var clusterInfoResponse = &ClusterInfo{}
		err := clusterInfoResponse.decode(response.Body)
		if err == nil {
			return clusterInfoResponse, nil
		} else {
			return nil, err
		}
	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}
}

type ListTopicResponse struct {
	Topics []string `json:"topicList"`
}

func (r *ListTopicResponse) decode(data []byte) error {
	return json.Unmarshal(data, r)
}

func (namesrv *namesrvsExt) ListTopic() ([]string, error) {

	var (
		response *remote.RemotingCommand
		err      error
	)

	//if s.Size() == 0, response will be nil, lead to panic below.
	if namesrv.Size() == 0 {
		return nil, primitive.NewRemotingErr("namesrv list empty")
	}

	for i := 0; i < namesrv.Size(); i++ {

		cmd := remote.NewRemotingCommand(ReqGetAllTopicListFromNameServer, nil, nil)
		ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
		response, err = namesrv.nameSrvClient.InvokeSync(ctx, namesrv.getNameServerAddress(), cmd)

		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, primitive.NewRemotingErr(err.Error())
	}

	switch response.Code {
	case ResSuccess:

		if response.Body == nil {
			return nil, primitive.NewMQClientErr(response.Code, response.Remark)
		}

		var listTopicResponse = &ListTopicResponse{}
		err := listTopicResponse.decode(response.Body)
		if err == nil {
			return listTopicResponse.Topics, nil
		} else {
			return nil, err
		}
	default:
		return nil, primitive.NewMQClientErr(response.Code, response.Remark)
	}

}

func (namesrv *namesrvsExt) QueryTopicRouteInfoFromServer(topic string) (*TopicRouteData, error) {

	return namesrv.queryTopicRouteInfoFromServer(topic)

}
