package loadbalance

import (
	"errors"
	"strconv"
)

type WeightNode struct {
	addr string
	// 权重值
	weight int
	// 临时权重
	currentWeight int
}

type WeightRoundBalance struct {
	targetNodes []*WeightNode
}

func (b *WeightRoundBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	weight, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{addr: params[0], weight: int(weight)}
	b.targetNodes = append(b.targetNodes, node)
	return nil
}

func (b *WeightRoundBalance) Next() string {
	// 累计权重
	total := 0

	// 最佳节点
	var best *WeightNode
	
	// 找到最佳节点
	for _, n := range b.targetNodes {
		// 统计所有有效权重之和
		total += n.weight

		// 变更节点临时权重为的节点临时权重+节点有效权重
		n.currentWeight += n.weight

		// 选择最大临时权重点节点
		if best == nil || n.currentWeight > best.currentWeight {
			best = n
		}
	}

	if best == nil {
		return ""
	}

	// 此时最佳节点已确定

	// 变更临时权重为 临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr
}

func (b *WeightRoundBalance) Get(string) (string, error) {
	return b.Next(), nil
}
