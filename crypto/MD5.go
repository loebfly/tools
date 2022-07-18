package crypto

import (
	"crypto/md5"
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"io"
	"os"
	"strings"
)

type md5Crypto struct{}

// EncodeString 对字符串进行 md5 加密
func (mc md5Crypto) EncodeString(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	md := fmt.Sprintf("%x", h.Sum(nil))
	return md
}

// EncodeBytes 对字节数组进行 md5 加密
func (mc md5Crypto) EncodeBytes(bytes []byte) string {
	h := md5.New()
	_, _ = io.WriteString(h, string(bytes))
	md := fmt.Sprintf("%x", h.Sum(nil))
	return md
}

// EncodeReader 对 io.Reader 进行 md5 加密
func (mc md5Crypto) EncodeReader(reader io.Reader) string {
	h := md5.New()
	_, _ = io.Copy(h, reader)
	md := fmt.Sprintf("%x", h.Sum(nil))
	return md
}

// EncodeFile 对文件进行 md5 加密
func (mc md5Crypto) EncodeFile(filename string) string {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return ""
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	h := md5.New()         //创建 md5 句柄
	_, err = io.Copy(h, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		fmt.Println(err)
		return ""
	}
	md := h.Sum(nil)                //计算 MD5 值，返回 []byte
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str
}

// EncodeStringMap 对字符串 map 进行 md5 加密
func (mc md5Crypto) EncodeStringMap(sm map[string]string) string {
	sortedmap := treemap.NewWithStringComparator()
	for k, v := range sm {
		sortedmap.Put(k, v)
	}
	sign := ""
	sortedmap.Each(func(key interface{}, value interface{}) {
		if key.(string) != "sign" && sm[key.(string)] != "" {
			sign = sign + key.(string) + "=" + value.(string) + "&"
		}
	})
	sign = strings.TrimRight(sign, "&")
	return mc.EncodeString(sign)
}
