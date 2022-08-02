package util

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

//声明新切片类型
type units []uint32

// Len 返回切片长度
func (x units) Len() int {
	return len(x)
}

// Less 比对两个数大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

// Swap 切片中两个值的交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// 当hash环上没有数据时，提示错误
var errEmpty = errors.New("hash环没有数据")

// Consistent 创建结构体保存一致性hash信息
type Consistent struct {
	circle       map[uint32]string // hash环，key为哈希值，值存放节点的信息
	sortedHashes units             // 已经排序的节点hash切片
	VirtualNode  int               // 虚拟节点个数，用来增加hash的平衡性
	sync.RWMutex                   // map 读写锁
}

// NewConsistent 创建一致性hash算法结构体，设置默认节点数量
func NewConsistent() *Consistent {
	return &Consistent{
		circle:      make(map[uint32]string), //初始化变量
		VirtualNode: 20,                      //设置虚拟节点个数
	}
}

// 自动生成key值
func (c *Consistent) generateKey(element string, index int) string {
	return element + strconv.Itoa(index) // 副本key生成逻辑
}

// 获取hash位置
func (c *Consistent) hashkey(key string) uint32 {
	if len(key) < 64 {
		var srcatch [64]byte                          // 声明一个数组长度为64
		copy(srcatch[:], key)                         // 拷贝数据到数组中
		return crc32.ChecksumIEEE(srcatch[:len(key)]) // 使用IEEE 多项式返回数据的CRC-32校验和
	}
	return crc32.ChecksumIEEE([]byte(key))
}

// 更新排序，方便查找
func (c *Consistent) updateSortedHashes() {
	hashes := c.sortedHashes[:0]
	// 判断切片容量，是否过大，如果过大则重置
	if cap(c.sortedHashes)/(c.VirtualNode*4) > len(c.circle) {
		hashes = nil
	}

	// 添加hashes
	for k := range c.circle {
		hashes = append(hashes, k)
	}

	// 对所有节点hash值进行排序，
	// 方便之后进行二分查找
	sort.Sort(hashes)
	// 重新赋值
	c.sortedHashes = hashes
}

// Add 向hash环中添加节点
func (c *Consistent) Add(element string) {
	c.Lock()         // 加锁
	defer c.Unlock() // 解锁
	c.add(element)
}

// 添加节点
func (c *Consistent) add(element string) {
	// 循环虚拟节点，设置副本
	for i := 0; i < c.VirtualNode; i++ {
		c.circle[c.hashkey(c.generateKey(element, i))] = element // 根据生成的节点添加到hash环中
	}
	// 更新排序
	c.updateSortedHashes()
}

// Remove 删除一个节点
func (c *Consistent) Remove(element string) {
	c.Lock()         // 加锁
	defer c.Unlock() // 解锁
	c.remove(element)
}

// 删除节点
func (c *Consistent) remove(element string) {
	for i := 0; i < c.VirtualNode; i++ {
		delete(c.circle, c.hashkey(c.generateKey(element, i)))
	}
	c.updateSortedHashes()
}

// Get 根据数据标示获取最近的服务器节点信息
func (c *Consistent) Get(name string) (string, error) {
	c.RLock()         // 加锁
	defer c.RUnlock() // 解锁
	// 如果为零则返回错误
	if len(c.circle) == 0 {
		return "", errEmpty
	}
	// 计算hash值
	key := c.hashkey(name)
	i := c.search(key)
	return c.circle[c.sortedHashes[i]], nil
}

// 顺时针查找最近的节点
func (c *Consistent) search(key uint32) int {
	// 查找算法
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	// 使用"二分查找"算法来搜索指定切片满足条件的最小值
	i := sort.Search(len(c.sortedHashes), f)
	// 如果超出范围则设置i=0
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return i
}
