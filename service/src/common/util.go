package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"sort"
)

func CheckInitRequest(signature, ts, nonce string) bool {
	array := []string{kToken, ts, nonce}
	sort.Strings(array)
	Log.Info("[CheckInitRequest]array:%v", array)
	s := array[0] + array[1] + array[2]

	hm := hmac.New(sha1.New, nil)
	hm.Write([]byte(s))
	sig := hm.Sum(nil)
	Log.Info("[CheckInitRequest]sig:%s", sig)

	return string(sig) == signature
}
