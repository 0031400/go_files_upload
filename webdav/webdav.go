package webdav

import (
	"go_files_upload/config"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(name string, path string) bool {
	file, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return false
	}
	url := config.WebdavPath + path
	request, err := http.NewRequest(http.MethodPut, url, file)
	if err != nil {
		log.Println(err)
		return false
	}
	request.SetBasicAuth(config.WebdavUsername, config.WebdavPassword)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return false
	}
	if res.StatusCode == http.StatusCreated {
		return true
	}
	if res.StatusCode == http.StatusConflict {
		MkDir(filepath.Dir(path))
		return false
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	res.Body.Close()
	log.Printf("%s %d %s\n", url, res.StatusCode, string(b))
	return false
}
func MkDir(relativeDir string) bool {
	url := config.WebdavPath + relativeDir + "/"
	req, err := http.NewRequest("MKCOL", url, nil)
	if err != nil {
		log.Println(err)
		return false
	}
	req.SetBasicAuth(config.WebdavUsername, config.WebdavPassword)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	if res.StatusCode == http.StatusCreated {
		log.Printf("mkdir %s\n", relativeDir)
		return true
	}
	if res.StatusCode == http.StatusConflict {
		suc := MkDir(filepath.Dir(relativeDir))
		return suc
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	res.Body.Close()
	log.Printf("%s %d %s\n", url, res.StatusCode, string(b))
	return false
}
