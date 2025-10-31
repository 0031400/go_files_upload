package config

import (
	"flag"
	"strings"
)

var JsonFile = ""
var WebdavPath = ""
var WebdavUsername = ""
var WebdavPassword = ""
var Dir = ""
var Exts = []string{}

func Init() {
	jsonFileFlag := flag.String("f", "data.json", "the durable json file")
	webdavPathFlag := flag.String("wp", "http://127.0.0.1:6065/one/", "the webdav server path")
	webdavUsernameFlag := flag.String("wuser", "admin", "the webdav server username")
	WebdavPasswordFlag := flag.String("wpasswd", "admin", "the webdav server password")
	DirFlag := flag.String("d", "data", "the data dir")
	ExtsFlag := flag.String("e", ".png", "the files exts")
	flag.Parse()
	JsonFile = *jsonFileFlag
	WebdavPath = *webdavPathFlag
	WebdavUsername = *webdavUsernameFlag
	WebdavPassword = *WebdavPasswordFlag
	Dir = *DirFlag
	Exts = strings.Split((*ExtsFlag), ",")
}
