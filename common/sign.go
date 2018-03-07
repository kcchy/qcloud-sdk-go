package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func sign(requestMethod string, endpoint string, params url.Values, secretKey string) (sign string, err error) {

	var srcStr string
	var flag_sha256 int
	flag_sha256 = 0

	srcStr += strings.ToUpper(requestMethod)
	srcStr += endpoint
	srcStr += RequestPath
	srcStr += "?"

	if params == nil {
		return "", nil
	}

	// 排序
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "SignatureMethod" && params[k][0] == "HmacSHA256" {
			flag_sha256 = 1
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var srcParms string
	for i := range keys {
		k := keys[i]
		srcParms += "&" + fmt.Sprintf("%v", k) + "=" + fmt.Sprintf("%v", params[k][0])
	}
	srcStr += srcParms[1:]
	fmt.Println(srcStr)

	if flag_sha256 == 1 {
		hmacObj := hmac.New(sha256.New, []byte(secretKey))
		hmacObj.Write([]byte(srcStr))
		sign = base64.StdEncoding.EncodeToString(hmacObj.Sum(nil))

	} else {
		hmacObj := hmac.New(sha1.New, []byte(secretKey))
		hmacObj.Write([]byte(srcStr))
		sign = base64.StdEncoding.EncodeToString(hmacObj.Sum(nil))
	}
	return sign, nil
}
