package crypto

import (
	"crypto/md5"
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"io"
	"os"
	"strings"
)

type MD5 struct{}

func (m *MD5) EncodeString(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	md := fmt.Sprintf("%x", h.Sum(nil))
	return md
}

func (m *MD5) EncodeBytes(bytes []byte) string {
	h := md5.New()
	_, _ = io.WriteString(h, string(bytes))
	md := fmt.Sprintf("%x", h.Sum(nil))
	return md
}

func (m *MD5) EncodeReader(reader io.Reader) string {
	h := md5.New()
	_, _ = io.Copy(h, reader)
	md := fmt.Sprintf("%x", h.Sum(nil))
	return md
}

func (m *MD5) EncodeFile(filename string) string {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return ""
	}
	defer f.Close()

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

func (m *MD5) EncodeStringMap(sm map[string]string) string {
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
	return m.EncodeString(sign)
}
