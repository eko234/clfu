package clfu

import (
	"container/list"
	"sync"
)

type ValueType interface{}
type KeyType interface{}

// FrequencyNode represents a node in the frequency linked list
type FrequencyNode struct {
	// frequency count - never decreases
	count uint
	// valuesList contains pointer to the head of values linked list
	valuesList *list.List
	// actual content of the next element
	inner *list.Element
}

// creates a new frequency list node with the given count
func NewFrequencyNode(count uint) *FrequencyNode {
	return &FrequencyNode{
		count:      count,
		valuesList: list.New(),
		inner:      nil,
	}
}

// KeyRefNode represents the value held on the LRU cache frequency list node
type KeyRefNode struct {
	// contains the actual value wrapped by a list element
	inner *list.Element
	// contains reference to the frequency node element
	parentFreqNode *list.Element
	// contains pointer to the key
	keyRef *KeyType
}

func NewKeyRefNode(keyRef *KeyType, parent *list.Element) *KeyRefNode {
	return &KeyRefNode{
		inner:          nil,
		parentFreqNode: parent,
		keyRef:         keyRef,
	}
}

// LFUCache implements all the methods and data-structures required for LFU cache
type LFUCache struct {
	// rwLock is a read-write mutex which provides concurrent reads but exclusive writes
	rwLock sync.RWMutex
	// a hash table of <KeyType, *ValueType> for quick reference of values based on keys
	lookupTable map[KeyType]*ValueType
	// internal linked list that contains frequency mapping
	frequencies *list.List
	// maxSize represents the maximum number of elements that can be in the cache before eviction
	maxSize uint
}

// MaxSize returns the maximum size of the cache at that point in time
func (lfu *LFUCache) MaxSize() uint {
	lfu.rwLock.RLock()
	defer lfu.rwLock.RUnlock()

	return lfu.maxSize
}

// CurrentSize returns the number of elements in that cache
func (lfu *LFUCache) CurrentSize() uint {
	lfu.rwLock.RLock()
	defer lfu.rwLock.RUnlock()

	return uint(len(lfu.lookupTable))
}

// customize the max size of the cache
func (lfu *LFUCache) SetMaxSize(size uint) {
	lfu.rwLock.Lock()
	defer lfu.rwLock.Unlock()

	lfu.maxSize = size
}

// create a new instance of LFU cache
func NewLFUCache(maxSize uint) *LFUCache {
	return &LFUCache{
		rwLock:      sync.RWMutex{},
		lookupTable: make(map[KeyType]*ValueType),
		maxSize:     maxSize,
		frequencies: list.New(),
	}
}
