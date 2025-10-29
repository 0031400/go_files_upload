package upload

import (
	"go_files_upload/config"
	"go_files_upload/webdav"
	"log"
	"sync"
	"time"
)

type task struct {
	name string
	path string
}

var mu sync.RWMutex
var tasks []task

func Init() {
	go func() {
		for {
			time.Sleep(time.Second)
			readUpload()
		}
	}()
}
func Upload(name, path string) {
	addTask(name, path)
}
func addTask(name, path string) {
	mu.Lock()
	tasks = append(tasks, task{name, path})
	mu.Unlock()
}
func readUpload() {
	if config.WebdavURL != "" {
		if len(tasks) != 0 {
			mu.RLock()
			t := tasks[0]
			mu.RUnlock()
			webdav.Upload(t.name, t.path)
			mu.Lock()
			tasks = tasks[1:]
			mu.Unlock()
		}
	} else {
		log.Println("lack uploader")
	}
}
