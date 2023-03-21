package loadbalance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	targetAddrs []string
	// 观察主体
	// conf LoadBalanceConf
}

func (b *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	b.targetAddrs = append(b.targetAddrs, params...)
	return nil
}

func (b *RandomBalance) Next() string {
	if len(b.targetAddrs) == 0 {
		return ""
	}

	return b.targetAddrs[rand.Intn(len(b.targetAddrs))]
}

func (b *RandomBalance) Get(string) (string, error) {
	return b.Next(), nil
}

func (b *RandomBalance) SetConf() {

}