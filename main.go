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
		items := scanDir(config.Dir)
		for _, item := range items {
			if isRightExt(item) && !record.HasRead(item) {
				record.AddRecord(item)
				newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(item))
				relativePath, err := filepath.Rel(config.Dir, filepath.Join(filepath.Dir(item), newFileName))
				if err != nil {
					log.Println(err)
					continue
				}
				upload.AddTask(item, relativePath)
				log.Printf("add record %s", item)
			}
		}
	}
}
func isRightExt(name string) bool {
	for _, v := range config.Exts {
		if strings.HasSuffix(name, v) {
			return true
		}
	}
	return false
}
func scanDir(baseDir string) []string {
	newItems := []string{}
	items, err := os.ReadDir(baseDir)
	if err != nil {
		log.Println(err)
		return []string{}
	}
	for _, v := range items {
		newPath := filepath.Join(baseDir, v.Name())
		if v.IsDir() {
			newItems = append(newItems, scanDir(newPath)...)
		} else {
			newItems = append(newItems, newPath)
		}
	}
	return newItems
}
