package durable

import (
	"encoding/json"
	"go_files_upload/config"
	"log"
	"os"
	"sync"
)

var mu sync.RWMutex

func Write(data []string) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	err = os.WriteFile(config.JsonFile, b, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
}
func Init() {
	_, err := os.Open(config.JsonFile)
	if os.IsNotExist(err) {
		log.Println("init durable file")
		Write([]string{})
	} else if err != nil {
		log.Println(err)
	}
}
func Read() []string {
	mu.RLock()
	defer mu.RUnlock()
	b, err := os.ReadFile(config.JsonFile)
	if err != nil {
		log.Println(err)
		return []string{}
	}
	d := []string{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}
