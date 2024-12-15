package storage

import "sync"

var (
	receiptStore = make(map[string]int)
	mu           sync.RWMutex
)

func StoreReceipt(id string, points int) {
	mu.Lock()
	defer mu.Unlock()
	receiptStore[id] = points
}

func GetPoints(id string) (int, bool) {
	mu.RLock()
	defer mu.RUnlock()
	points, found := receiptStore[id]
	return points, found
}
