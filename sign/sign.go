package sign

import (
	"bytes"
	"crypto/md5"
	"github.com/iancoleman/orderedmap"
	"net/url"
	"sort"
)

func encodeHex(data []byte) []byte {
	l := len(data)
	out := make([]byte, l<<1)
	i := 0
	toDigits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for var5 := 0; i < l; i++ {
		out[var5] = toDigits[(240&data[i])>>4]
		var5++
		out[var5] = toDigits[15&data[i]]
		var5++
	}
	return out
}

func GetSignature(params url.Values, appSecret string) (string, error) {
	// 先将参数以其参数名的字典序升序进行排序
	sortedParams := orderedmap.New()
	for k, v := range params {
		if v != nil && len(v) > 0 {
			sortedParams.Set(k, v[0])
		}
	}
	sortedParams.SortKeys(sort.Strings)
	// 遍历排序后的字典，将所有参数按"key=value"格式拼接在一起
	var baseString bytes.Buffer
	for _, k := range sortedParams.Keys() {
		v, _ := sortedParams.Get(k)
		if k != "" && k != "sign" && k != "key" && v != "" {
			baseString.WriteString(k)
			baseString.WriteString("=")
			baseString.WriteString(v.(string))
			baseString.WriteString("&")
		}
	}
	if baseString.Len() > 0 {
		baseString.Truncate(baseString.Len() - 1)
		baseString.WriteString(appSecret)
	}
	// 使用MD5对待签名串求签
	h := md5.New()
	h.Write(baseString.Bytes())
	return string(encodeHex(h.Sum(nil))), nil
}
