package test_example

import (
	"fmt"
	"github.com/loebfly/tools"
	"testing"
)

/*
	Math   = new(math.Enter)    // 计算工具
	String = new(stringT.Enter) // 字符串工具
	Map    = new(mapT.Enter)    // map工具
	Crypto = new(crypto.Enter)  // 加密工具
	Gin    = new(ginT.Enter)    // gin工具
	Json   = new(jsonT.Enter)   // json工具
	File   = new(fileT.Enter)   // 文件工具
	SQL    = new(sqlT.Enter)    // sql工具
	IP     = new(ipT.Enter)     // ip工具
	Shell  = new(shell.Enter)   // shell工具
	SFTP   = new(sftpT.Enter)   // sftp工具
*/

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

	res = verify.IsIP("127.0.0.1")
	fmt.Println("IsIP:", res)

	res = verify.IsEnglish("abc")
	fmt.Println("IsEnglish:", res)

	res = verify.IsNumber("123")
	fmt.Println("IsNumber:", res)

	res = verify.IsLowerCase("abc")
	fmt.Println("IsLowerCase:", res)

	res = verify.IsUpperCase("ABC")
	fmt.Println("IsUpperCase:", res)
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
