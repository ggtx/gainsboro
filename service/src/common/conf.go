package common

import (
	"flag"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var (
	rootPath = ""
)

type SafeAccessToken struct {
	AccessToken string
	ExpiresIn   int64
	AtStartTime int64
	sync.RWMutex
}

var ApiAccessToken SafeAccessToken

type GainConfig struct {
	AppId     string
	AppSecret string
}

var GainServerConfig GainConfig

func init() {
	flag.StringVar(&rootPath, "rootpath", "/opt/gainsboro/", "rootPath")
	flag.Parse()

	f, err := os.Open(rootPath + "conf/gainsboro.yaml")
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(b, &GainServerConfig)
	if err != nil {
		panic(err)
	}

	//Log.Debug("conf:%s,%s", GainServerConfig.AppId, GainServerConfig.AppSecret)

	GetNewAccessToken(time.Now().Unix())

	//Log.Debug("access_token:%s, expires_in:%d, atstattime:%d", ApiAccessToken.AccessToken, ApiAccessToken.ExpiresIn, ApiAccessToken.AtStartTime)
}
