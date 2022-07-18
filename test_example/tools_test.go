package test_example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/tools"
	"github.com/loebfly/tools/sftp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"testing"
	"time"
)

func TestMath(t *testing.T) {
	random := tools.Math.Random

	rI := random.IntFromRange(0, 10)
	rF := random.Float64FromRange(0, 10)
	fmt.Println("随机范围整数:", rI)
	fmt.Println("随机范围浮点数:", rF)

	b := random.Bool()
	fmt.Println("随机布尔值", b)

	src := random.SrcEnNumber()
	//src = random.SrcEn()
	//src = random.SrcEnUpper()
	//src = random.SrcEnLower()
	length := 10
	res := random.Generate(src, length)
	fmt.Println("随机字符串:", res)

	condition := tools.Math.Condition

	condition.IfThenElse(true, func() {
		fmt.Println("true")
	}, func() {
		fmt.Println("false")
	})

	condition.IfThen(true, func() {
		fmt.Println("true")
	})

	ifNil := condition.DefaultIfNil(nil, 1)
	fmt.Println("ifNil:", ifNil)

	ifNil = condition.DefaultIfNil(2, 1)
	fmt.Println("ifNotNil:", ifNil)

	nonNil := condition.FirstNonNil(nil, 3)
	fmt.Println("nonNil:", nonNil)
}

func TestString(t *testing.T) {
	verify := tools.String.Verify
	res := verify.IsEmail("abc@example.com")
	fmt.Println("IsEmail:", res)

	res = verify.IsCNMobile("13800138000")
	fmt.Println("IsCNMobile:", res)

	res = verify.IsChinese("中文")
	fmt.Println("IsChinese:", res)

	res = verify.IsIDCard("420102199001020301")
	fmt.Println("IsIDCard:", res)

	res = verify.IsEnglish("abc")
	fmt.Println("IsEnglish:", res)

	res = verify.IsNumber("123")
	fmt.Println("IsNumber:", res)

	res = verify.IsLowerCase("abc")
	fmt.Println("IsLowerCase:", res)

	res = verify.IsUpperCase("ABC")
	fmt.Println("IsUpperCase:", res)
}

func TestCrypto(t *testing.T) {
	crypto := tools.Crypto

	src := "loebfly"
	encrypt := crypto.MD5.EncodeString(src)
	fmt.Println("MD5 Encrypt:", encrypt)

	encrypt = crypto.MD5.EncodeBytes([]byte(src))
	fmt.Println("MD5 Encrypt:", encrypt)

	encrypt = crypto.Base64.EncodeString(src)
	fmt.Println("Base64 Encrypt:", encrypt)

	encrypt = string(crypto.Base64.EncodeByte([]byte(src)))
	fmt.Println("Base64 Encrypt:", encrypt)

	decrypt := crypto.Base64.DecodeString(encrypt)
	fmt.Println("Base64 Decrypt:", decrypt)

	decrypt = string(crypto.Base64.DecodeByte([]byte(encrypt)))
	fmt.Println("Base64 Decrypt:", decrypt)
}

func TestMap(t *testing.T) {
	extend := tools.Map.Extend

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	srcIMap := map[string]interface{}{
		"name": "loebfly",
		"age":  18,
	}

	isExist := extend.IsExistI(srcIMap, "name")
	fmt.Println("interface map isExist:", isExist)

	isEmpty := extend.IsEmptyI(srcIMap, "name")
	fmt.Println("interface map isEmpty:", isEmpty)

	keys := extend.GetKeysI(srcIMap)
	fmt.Println("interface map keys:", keys)

	iValues := extend.GetValuesI(srcIMap)
	fmt.Println("interface map values:", iValues)

	iJson := extend.ToJsonFromMapI(srcIMap)
	fmt.Println("interface map json:", iJson)

	srcSMap := map[string]string{
		"name": "loebfly",
		"age":  "18",
	}

	isExist = extend.IsExistS(srcSMap, "name")
	fmt.Println("string map isExist:", isExist)

	isEmpty = extend.IsEmptyS(srcSMap, "name")
	fmt.Println("string map isEmpty:", isEmpty)

	keys = extend.GetKeysS(srcSMap)
	fmt.Println("string map keys:", keys)

	sValues := extend.GetValuesS(srcSMap)
	fmt.Println("string map values:", sValues)

	sJson := extend.ToJsonFromMapS(srcSMap)
	fmt.Println("string map json:", sJson)

	user := User{
		Name: "loebfly",
		Age:  18,
	}

	obj := extend.ToMapIFromObj(user)
	fmt.Println("object map:", obj)

	merge := tools.Map.Merge
	stringRes := merge.Strings(map[string]string{"a": "1", "b": "2"}, map[string]string{"c": "3", "d": "4"})
	fmt.Println("Merge Strings:", stringRes)

	interfaceRes := merge.Interfaces(map[string]interface{}{"a": 1.1, "b": 2.2}, map[string]interface{}{"c": 3.3, "d": 4.4})
	fmt.Println("Merge Interfaces:", interfaceRes)

	intRes := merge.Ints(map[string]int{"a": 1, "b": 2}, map[string]int{"c": 3, "d": 4})
	fmt.Println("Merge Ints:", intRes)
}

func TestGin(t *testing.T) {
	ctx := gin.Context{}
	params := tools.Gin.GetParams(&ctx)
	fmt.Println("params:", params)
}

func TestJson(t *testing.T) {
	json := tools.Json

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{
		Name: "loebfly",
		Age:  18,
	}

	jsonStr := json.ToJson(user)
	fmt.Println("ToJson:", jsonStr)

	userMap := map[string]interface{}{
		"name": "loebfly",
		"age":  18,
	}

	jsonStr = json.ToJson(userMap)
	fmt.Println("ToJson:", jsonStr)

	user2 := User{}
	json.ToObjS(jsonStr, &user2)
	fmt.Println("ToObjS:", user2)

	json.ToObjB([]byte(jsonStr), &user2)
	fmt.Println("ToObjB:", user2)

	json.ToObjI(userMap, &user2)
	fmt.Println("ToObjI:", user2)
}

func TestFile(t *testing.T) {

	selfPath := tools.File.SelfPath()
	fmt.Println("selfPath:", selfPath)

	selfDir := tools.File.SelfDir()
	fmt.Println("selfDir:", selfDir)

	src := "/Users/luchunqing/Documents/github.com/tools/test_example/tools_test.go"

	basename := tools.File.GetBasename(src)
	fmt.Println("Basename:", basename)

	dir := tools.File.GetDir(src)
	fmt.Println("Dir:", dir)

	ext := tools.File.GetExt(src)
	fmt.Println("Ext:", ext)

	mTime, err := tools.File.GetMTime(src)
	if err != nil {
		fmt.Println("GetMTime err:", err)
	}
	fmt.Println("MTime:", mTime)

	size, err := tools.File.GetSize(src)
	if err != nil {
		fmt.Println("GetSize err:", err)
	}
	fmt.Println("Size:", size)

	isExist := tools.File.IsExist(src)
	fmt.Println("IsExist:", isExist)

	isFile := tools.File.IsFile(src)
	fmt.Println("IsFile:", isFile)

	isDir := tools.File.IsDir(src)
	fmt.Println("IsDir:", isDir)

	//contentBytes, err := tools.File.ReadBytes(src)
	//if err != nil {
	//	fmt.Println("ReadBytes err:", err)
	//}
	//fmt.Println("Content:", string(contentBytes))
	//
	//contentString, err := tools.File.ReadString(src)
	//if err != nil {
	//	fmt.Println("ReadString err:", err)
	//}
	//fmt.Println("Content:", contentString)

	contentLines, err := tools.File.ReadLines(src)
	if err != nil {
		fmt.Println("ReadLines err:", err)
	}
	fmt.Println("Content:", contentLines)

	_, err = tools.File.WriteString(dir+"/hello1.txt", "hello world")
	if err != nil {
		fmt.Println("WriteString err:", err)
	}

	_, err = tools.File.WriteBytes(dir+"/hello2.txt", []byte("hello world"))
	if err != nil {
		fmt.Println("WriteBytes err:", err)
	}

	err = tools.File.Remove(dir + "/hello1.txt")
	if err != nil {
		fmt.Println("Remove err:", err)
	}

	err = tools.File.Remove(dir + "/hello2.txt")
	if err != nil {
		fmt.Println("Remove err:", err)
	}
}

func TestSQL(t *testing.T) {
	sql := tools.SQL

	verify, s := sql.Verify("select * from user where id = 1")
	fmt.Println("Verify:", verify, s)
}

func TestIP(t *testing.T) {
	ip := tools.IP

	isIP := ip.IsCorrect("127.0.0.1")
	fmt.Println("IsCorrect:", isIP)

	ipStr := ip.GetLocal()
	fmt.Println("Local:", ipStr)

	isIntranet := ip.IsIntranet("127.0.0.1")
	fmt.Println("IsIntranet:", isIntranet)

	IPv4, err := ip.GetV4s()
	if err != nil {
		fmt.Println("GetV4s err:", err)
	}
	fmt.Println("IPv4:", IPv4)

	isUsed := ip.IsPortUsed(10000)
	fmt.Println("IsPortUsed:", isUsed)
}

func TestShell(t *testing.T) {
	shell := tools.Shell

	err := shell.Firewall.AddPort("8080")
	if err != nil {
		fmt.Println("AddPort err:", err)
	}
	fmt.Println("AddPort:", "8080")

	err = shell.Firewall.RemovePort("8080")
	if err != nil {
		fmt.Println("RemovePort err:", err)
	}
	fmt.Println("RemovePort:", "8080")

	err = shell.Nginx.Restart()
	if err != nil {
		return
	}
}

func TestSFTP(t *testing.T) {
	sftpClient, sshClient, err := tools.SFTP.Connect("user", "pwd", "127.0.0.1", 22)
	if err != nil {
		return
	}
	defer func(SFTP *sftpT.Enter, sftpClient *sftp.Client, sshClient *ssh.Client) {
		err := SFTP.Close(sftpClient, sshClient)
		if err != nil {
			fmt.Println("Close err:", err)
		}
	}(tools.SFTP, sftpClient, sshClient)

	err = tools.SFTP.UploadFile(sftpClient, "./test.txt", "/tmp/test.txt")
	if err != nil {
		fmt.Println("UploadFile err:", err)
	}
}

func TestTime(t *testing.T) {
	timeStr := tools.Time.Fmt.GetNowDateTime()
	fmt.Println("GetNowDateTime:", timeStr)

	timeStr = tools.Time.Fmt.GetNowDate()
	fmt.Println("GetNowDate:", timeStr)

	timeStr = tools.Time.Fmt.GetNowTime()
	fmt.Println("GetNowTime:", timeStr)

	timeStr = tools.Time.Fmt.GetNowWeekday()
	fmt.Println("GetNowWeekday:", timeStr)

	timeStr = tools.Time.Fmt.Convert("2018-01-01 00:00:00", "yyyy-MM-dd")
	fmt.Println("Convert:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "2006-01-02 15:04:05")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "yyyy-MM-dd HH:mm:ss")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "yyyy-MM-dd HH:mm")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "yyyy-MM-dd HH")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "yyyy-MM-dd")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "yyyy-MM")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "yyyy")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "MMMM")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "MMM")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "MM")
	fmt.Println("Get:", timeStr)

	timeStr = tools.Time.Fmt.Get(time.Now(), "M")
	fmt.Println("Get:", timeStr)
}
