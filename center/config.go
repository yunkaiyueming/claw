package center

import (
	"encoding/xml"
	"flag"
	"io/ioutil"
)

var (
	configFile = flag.String("config", "./config.xml", "config file name")

	BaseConfig BaseConfigPack
)

func init() {
	flag.Parse()
}

func InitConfig() {
	GetConfig(&BaseConfig)
}

func GetConfig(v interface{}) {
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(content, v)
	if err != nil {
		panic(err)
	}
}


// Base config
type BaseConfigPack struct {
	XMLName xml.Name `xml:"clawconfig"`
	Master MasterConfig `xml:"master"`
	Harbor HarborConfig `xml:"harbor"`
	Gate GateConfig `xml:"gate"`
}

type MasterConfig struct {
	ListenAddr string `xml:"listenAddr,attr"`
	IsMaster bool `xml:"isMaster,attr"`
}

type HarborConfig struct {
	Id string `xml:"id,attr"`
}

type GateConfig struct {
	ListenAddr string `xml:"listenAddr,attr"`
}
