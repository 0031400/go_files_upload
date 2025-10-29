package record

import (
	"go_files_upload/durable"
	"slices"
	"sync"
)

var list []string
var mu sync.RWMutex

func Init() {
	mu.Lock()
	list = durable.Read()
	mu.Unlock()
}
func HasRead(name string) bool {
	mu.RLock()
	defer mu.RUnlock()
	return slices.Contains(list, name)
}
func AddRecord(name string) {
	mu.Lock()
	list = append(list, name)
	mu.Unlock()
	durable.Write(list)
}
