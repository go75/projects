package loadbalance

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32

type UInt32Slice []uint32

func (s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s UInt32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ConsistentHashBalance struct {
	lock *sync.RWMutex
	// 哈希函数
	hash Hash
	// 虚拟节点个数
	replicas int
	// 已排序的节点hash切片
	keys UInt32Slice
	// key是hash值, value是addr
	hashMap map[uint32]string
	// 观察主体
	//conf LoadBalanceConf
}

func NewConsistentHashBalance(replicas int, fn Hash) *ConsistentHashBalance {
	if fn == nil {
		return &ConsistentHashBalance{
			lock: new(sync.RWMutex),
			hash: crc32.ChecksumIEEE,
			replicas: replicas,
			hashMap: make(map[uint32]string, 0),
		}
	}

	return &ConsistentHashBalance{
		lock: new(sync.RWMutex),
		hash: fn,
		replicas: replicas,
		hashMap: make(map[uint32]string, 0),
	}
}

// 验证是否为空
func (b *ConsistentHashBalance) IsEmpty() bool {
	return len(b.keys) == 0
}

func (b *ConsistentHashBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	b.lock.Lock()
	defer b.lock.Unlock()
	// 多个虚拟节点映射同一个addr
	for i := 0; i < b.replicas; i++ {
		hash := b.hash([]byte(strconv.Itoa(i) + addr))
		b.keys = append(b.keys, hash)
		b.hashMap[hash] = addr
	}

	// 对所有虚拟节点的hash值进行排序, 方便之后的二分查找
	sort.Sort(b.keys)
	return nil
}

func (b *ConsistentHashBalance) Get(key string) (string, error) {
	if b.IsEmpty() {
		return "", errors.New("node is empty")
	}
	hash := b.hash([]byte(key))

	b.lock.RLock()
	defer b.lock.RUnlock()
	// 通过二分查找获取最优节点, 第一个"服务器hash值"大于"数据hash值"的就是选择的服务器节点
	index := sort.Search(len(b.keys), func(i int) bool {
		return b.keys[i] >= hash
	})

	// 如果查找结果 大于 服务器节点哈希数组的最大索引，表示此时该对象哈希值位于最后一个节点之后，那么放入第一个节点中
	if index == len(b.keys) {
		index = 0
	}
	return b.hashMap[b.keys[index]], nil
}