package loadbalance

type LbType uint8

const (
	LbRandom LbType = iota
	LbRound
	LbWeightRound
	LbConsistentHash
)

type LoadBalance interface {
	Add(...string) error
	Get(string) (string, error)
	// 服务发现
	// Update()
}

func LoadBalanceFactory(lbType LbType) LoadBalance {
	switch lbType {
	case LbRandom:
		return &RandomBalance{}
	case LbRound:
		return &RoundBalance{}
	case LbWeightRound:
		return &WeightRoundBalance{}
	default:
		return NewConsistentHashBalance(10, nil)
	}
}