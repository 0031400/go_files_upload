package main

import (
	"fmt"
	"go_files_upload/config"
	"go_files_upload/durable"
	"go_files_upload/logger"
	"go_files_upload/record"
	"go_files_upload/upload"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	logger.Init()
	config.Init()
	durable.Init()
	record.Init()
	upload.Init()
	for {
		time.Sleep(time.Second)
		files, err := os.ReadDir(config.Dir)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, file := range files {
			name := file.Name()
			if strings.HasSuffix(name, config.Ext) && !record.HasRead(name) {
				record.AddRecord(name)
				go func() {
					newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(name))
					upload.Upload(filepath.Join(config.Dir, name), newFileName)
				}()
				log.Printf("add record %s", name)
			}
		}
	}
}
