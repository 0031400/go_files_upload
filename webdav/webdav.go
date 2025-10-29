package webdav

import (
	"fmt"
	"go_files_upload/config"
	"io"
	"log"
	"net/http"
	"os"
)

func Upload(name string, path string) {
	file, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
	url := fmt.Sprintf("%s%s%s", config.WebdavURL, config.WebdavPath, path)
	request, err := http.NewRequest(http.MethodPut, url, file)
	if err != nil {
		log.Println(err)
		return
	}
	request.SetBasicAuth(config.WebdavUsername, config.WebdavPassword)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	if res.StatusCode != 201 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%s %d %s\n", url, res.StatusCode, string(b))
	}
}
