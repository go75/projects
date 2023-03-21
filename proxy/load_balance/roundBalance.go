package loadbalance

import "errors"

type RoundBalance struct {
	curIndex int
	targetAddrs []string
	// 观察主体
	// conf LoadBalanceConf
}

func (b *RoundBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	b.targetAddrs = append(b.targetAddrs, params...)
	return nil
}

func (b *RoundBalance) Next() string {
	if len(b.targetAddrs) == 0 {
		return ""
	}

	b.curIndex = (b.curIndex+1) % len(b.targetAddrs)

	return b.targetAddrs[b.curIndex]
}

func (b *RoundBalance) Get(string) (string, error) {
	return b.Next(), nil
}

func (b *RoundBalance) SetConf() {

}