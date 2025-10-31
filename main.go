package main

import (
	"fmt"
	"go_files_upload/config"
	"go_files_upload/durable"
	"go_files_upload/logger"
	"go_files_upload/record"
	"go_files_upload/webdav"
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
	for {
		time.Sleep(time.Second)
		items := scanDir(config.Dir)
		for _, item := range items {
			if isRightExt(item) {
				work(item)
			}
		}
	}
}
func work(itemPath string) {
	if record.HasRead(itemPath) {
		return
	}
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(itemPath))
	relativeDir, err := filepath.Rel(config.Dir, filepath.Dir(itemPath))
	if err != nil {
		log.Println(err)
		return
	}
	relativePath := filepath.Join(relativeDir, newFileName)
	suc := webdav.Upload(itemPath, relativePath)
	if !suc {
		return
	}
	record.AddRecord(itemPath)
	log.Printf("add record %s", itemPath)
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
