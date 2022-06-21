package internal

import (
	"math/rand"

	"github.com/rocketmq-exporter-go/internal/utils"
)

/*
Selects a (preferably master) broker address from the registered list.
If the master's address cannot be found,
a slave broker address is selected in a random manner.

@return Broker address.
*/
func (bd *BrokerData) SelectBrokerAddr() string {
	addr, exist := bd.BrokerAddresses[MasterId]
	if !exist && len(bd.BrokerAddresses) > 0 {
		i := utils.AbsInt(rand.Int())
		i = i % len(bd.BrokerAddresses)
		for _, v := range bd.BrokerAddresses {
			if i <= 0 {
				addr = v
				break
			}
			i--
		}
	}

	return addr
}
