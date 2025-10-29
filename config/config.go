package config

import "flag"

var JsonFile = ""
var WebdavURL = ""
var WebdavPath = ""
var WebdavUsername = ""
var WebdavPassword = ""
var Dir = ""
var Ext = ""

func Init() {
	jsonFileFlag := flag.String("f", "data.json", "the durable json file")
	webdavURLFlag := flag.String("wu", "http://127.0.0.1:6065", "the webdav server url")
	webdavPathFlag := flag.String("wp", "/one/", "the webdav server path")
	webdavUsernameFlag := flag.String("wuser", "admin", "the webdav server username")
	WebdavPasswordFlag := flag.String("wpasswd", "admin", "the webdav server password")
	DirFlag := flag.String("d", "data", "the data dir")
	ExtFlag := flag.String("e", ".png", "the files ext")
	flag.Parse()
	JsonFile = *jsonFileFlag
	WebdavURL = *webdavURLFlag
	WebdavPath = *webdavPathFlag
	WebdavUsername = *webdavUsernameFlag
	WebdavPassword = *WebdavPasswordFlag
	Dir = *DirFlag
	Ext = *ExtFlag
}
