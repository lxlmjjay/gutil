package util

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"html"
	"html/template"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

func Inter2Int64(id interface{}) (data int64) {
	of := reflect.TypeOf(id)
	if of.String() == "int64" {
		data = id.(int64)
		return
	} else if of.String() == "string" {
		data = StringToInt64(id.(string))
	} else {
		return 0
	}
	return
}

//每一个参数转string 并过滤html js 标签
func CheckString(para interface{}) (p string) {
	if para == nil || reflect.TypeOf(para).String() != "string" {
		return ""
	}
	p = fmt.Sprintf("%s", para)
	//p = html.EscapeString(p)
	//p = template.JSEscapeString(p)
	return strings.TrimSpace(p)
}

//每一个参数转string 并过滤html js 标签
func StringFilter(para interface{}) (p string) {
	if para == nil || reflect.TypeOf(para).String() != "string" {
		return ""
	}
	p = fmt.Sprintf("%s", para)
	p = html.EscapeString(p)
	p = template.JSEscapeString(p)
	return strings.TrimSpace(p)
}

func CheckInt(para string) int {
	i, err := strconv.Atoi(para)
	if err != nil {
		return 0
	}
	return i
}

func CheckInt64(para string) int64 {
	i, err := strconv.Atoi(para)
	if err != nil {
		return 0
	}
	return int64(i)
}

func StringUnescape(content string) string {
	content = UnescapeUnicode(content)
	i := 0
	s := html.UnescapeString(content)
	if strings.Contains(s, "&lt") {
		s = html.UnescapeString(s)
		i++
		if i >= 3 {
			return s
		}
	}
	return s
}

func UnescapeUnicode(raw string) string {
	//str, err := strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
	raw = strings.Replace(raw, `\\u`, `\u`, -1)
	s := strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1)
	str, err := strconv.Unquote(s)
	if err != nil {
		return raw
	}
	return str
}

func Int64ToString(i int64) string {
	return fmt.Sprintf("%v", i)
}

func StringToInt64(s string) int64 {
	if len(strings.TrimSpace(s)) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("StringToInt64-", s, err)
	}
	return i
}

//将int64 数组转 string数组
func StringArr2SqlStr(a []string) (s string) {
	if len(a) == 0 {
		return
	}
	for _, v := range a {
		s += `"` + v + `",`
	}
	s = strings.TrimRight(s, ",")
	return
}

//将int64 数组转 string数组
func Int64Arr2StringArr(a []int64) (b []string) {
	for _, v := range a {
		b = append(b, Int64ToString(v))
	}
	return
}

func Byte2Int64(i []byte) int64 {
	bytebuff := bytes.NewBuffer(i)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return data
}

func Byte2String(i []byte) string {
	return string(i[:])
}

func Int642Bytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

//去除字符串首尾逗号
func TrimComma(str string) string {
	return strings.TrimRight(strings.TrimLeft(str, ","), ",")
}

//bool 返回是old 追加 mod
func AppendStrComma(old string, mod string) string {
	modTrim := TrimComma(mod)
	oldTrim := TrimComma(old)
	if len(oldTrim) == 0 {
		return modTrim + ","
	}
	old = "," + oldTrim + ","
	return old + modTrim + ","
}

//判断加逗号的字符串mod是否存在old
func ExistStrComma(old string, mod string) bool {
	old = "," + TrimComma(old) + ","
	mod = "," + TrimComma(mod) + ","
	if strings.Index(old, mod) == -1 {
		return false
	}
	return true
}

func SubStrComma(old string, mod string) string {
	modTrim := TrimComma(mod)
	oldTrim := TrimComma(old)
	if len(oldTrim) == 0 {
		return ""
	}
	old = "," + TrimComma(old) + ","
	mod = "," + modTrim + ","
	return strings.Replace(old, mod, ",", -1)
}

func CheckLenString(para ...string) bool {
	for _, v := range para {
		if len(v) <= 0 {
			return false
		}
	}
	return true
}

func CheckLenInt64(para ...int64) bool {
	for _, v := range para {
		if v <= 0 {
			return false
		}
	}
	return true
}

func Hex2Int64(hex string) int64 {
	s := strings.TrimPrefix(hex, "0x")
	if len(s) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func Hex2Float64(hexs string) float64 {
	s := strings.TrimPrefix(hexs, "0x")
	if len(s) == 0 {
		return 0
	}
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	decodeString, err2 := hex.Decode(dst,src)
	fmt.Println(333,decodeString, err2)
	fmt.Println(444,s)
	i, err := strconv.ParseFloat(s, 16)
	fmt.Println(555,i,err)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func Int64ToHex(i int64) string {
	return "0x" + strconv.FormatInt(i, 16)
}

func Float64ToBigInt(f float64, basePriceFloat *big.Float) *big.Int {
	floatStr := strconv.FormatFloat(f, 'f', -1, 64)
	float, b := new(big.Float).SetString(floatStr)
	if !b {
		return nil
	}
	mul := new(big.Float).Mul(float, basePriceFloat)
	num := fmt.Sprintf("%.f", mul)
	setString, b := new(big.Int).SetString(num, 10)
	if !b {
		return nil
	}
	return setString
}
