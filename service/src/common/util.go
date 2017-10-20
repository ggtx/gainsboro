package common

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

func CheckInitRequest(signature, ts, nonce string) bool {
	array := []string{kToken, ts, nonce}
	sort.Strings(array)
	Log.Debug("[CheckInitRequest]array:%v", array)

	hm := sha1.New()
	for _, a := range array {
		hm.Write([]byte(a))
	}
	sig := fmt.Sprintf("%x", hm.Sum(nil))

	Log.Debug("[CheckInitRequest]sig:%s", sig)

	return string(sig) == signature
}

type ResponseAccessTokenJson struct {
	At        string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

func GetNewAccessToken(tsnow int64) string {
	at := ""
	ApiAccessToken.RLock()
	timeout := ApiAccessToken.ExpiresIn-300 < tsnow-ApiAccessToken.AtStartTime
	ApiAccessToken.RUnlock()
	if !timeout {
		ApiAccessToken.RLock()
		at = ApiAccessToken.AccessToken
		ApiAccessToken.RUnlock()
		return at
	}

	atApi := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + GainServerConfig.AppId + "&secret=" + GainServerConfig.AppSecret
	resp, err := http.Get(atApi)
	if err != nil {
		Log.Warn("Get access_token err:%v", err)
		return ""
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Warn("Read Body err:%v", err)
		return ""
	}

	//Log.Debug("access_token resp:%v", string(b))

	rat := &ResponseAccessTokenJson{}
	if err := json.Unmarshal(b, rat); err != nil {
		Log.Warn("Unmarshal err:%v", err)
		return ""
	}

	ApiAccessToken.Lock()
	ApiAccessToken.AccessToken = rat.At
	ApiAccessToken.ExpiresIn = rat.ExpiresIn
	ApiAccessToken.AtStartTime = time.Now().Unix()
	ApiAccessToken.Unlock()

	return rat.At
}
