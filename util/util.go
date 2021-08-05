package util

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)
//检查参数map中 key 是否存在
func CheckMInt(para map[int64]interface{}, field int64) bool {
	if para == nil {
		return false
	}
	if _, ok := para[field]; !ok {
		return false
	}
	return true
}

//检查参数map中 key 是否存在
func CheckM(para map[string]interface{}, field ...string) bool {
	if para == nil {
		return false
	}
	for _, e := range field {
		if _, ok := para[e]; !ok {
			return false
		}
	}
	return true
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func RemoveRepeatInt64(uids []int64) []int64 {
	tempUids := []int64{}
	for _, i := range uids {
		if len(tempUids) == 0 {
			tempUids = append(tempUids, i)
		} else {
			for k, v := range tempUids {
				if i == v {
					break
				}
				if k == len(tempUids)-1 {
					tempUids = append(tempUids, i)
				}
			}
		}
	}
	return tempUids
}

//去掉url参数
func GetImgUri(uri string) string {
	parse, _ := url.Parse(uri)
	return "https://" + parse.Host + parse.Path
}

//查看d中元素arr中是否存在 如果存在去除 返回d数去除后
func InArrayCutHas(arr []int64, d []int64) (res []int64) {
	if len(d) == 0 {
		return
	}
	if len(arr) == 0 {
		return d
	}
	for _, v := range d {
		repeat := false
		for _, vv := range arr {
			if v == vv {
				repeat = true
				break
			}
		}
		if !repeat {
			res = append(res, v)
		}
	}
	return
}

func InArrayInt64(arr []int64, s int64) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}

func InArrayInt(arr []int, s int) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}

func InArrayString(arr []string, s string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}

//float 保留2位小数
func FloatRound2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//float 保留4位小数
func FloatRound4(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}

func FloatToString(num float64) (convert string) {
	convert = strconv.FormatFloat(num, 'f', -1, 64)
	return
}

//科学计数法表示
func FloatToStringExp(num float64) (convert string) {
	convert = strconv.FormatFloat(num, 'E', -1, 64)
	return
}

//科学计数法表示
func StringToFloat64(num string) (convert float64) {
	float, err := strconv.ParseFloat(num, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return float
}

func GetRealIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

//判断类型是否是decimal的字段 返回数组 用于sql .Omit(arr...)
func CheckDecimalArr(o interface{}) (res []string) {
	typ := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o).Elem()
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type.String() == "decimal.Decimal" {
			f := v.Field(i)
			d := f.Interface().(decimal.Decimal)
			if d.IsZero() {
				name := ToUnderlineLower(typ.Field(i).Name)
				res = append(res, name)
			}
		}
	}
	return
}

//将struct name 转为 sql 字段类型  如 MyName 转为 my_name
func ToUnderlineLower(s string) string {
	arr := strings.Split(s, "")
	for k, v := range s {
		if unicode.IsUpper(v) {
			if k == 0 {
				arr[0] = strings.ToLower(arr[0])
			} else {
				arr[k] = "_" + strings.ToLower(arr[k])
			}
		}
	}
	return strings.Join(arr, "")
}

//手机号中间4位换星星
func MobileToStar(uid int64) string {
	uidStr := Int64ToString(uid)
	return uidStr[:3] + "****" + uidStr[7:]
}

func Json(o interface{}) string {
	marshal, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	return string(marshal)
}

/**
  Markdown自动换行
*/
func MarkdownAutoNewline(str string) string {
	re, _ := regexp.Compile("\\ *\\n")
	str = re.ReplaceAllLiteralString(str, "  \n")
	//m.Content=strings.Replace(m.Content, "\n", "  \n", -1)
	reg := regexp.MustCompile("```([\\s\\S]*)```")
	//返回str中第一个匹配reg的字符串
	data := reg.Find([]byte(str))
	strs := strings.Replace(string(data), "  \n", "\n", -1)
	re, _ = regexp.Compile("```([\\s\\S]*)```")
	return re.ReplaceAllLiteralString(str, strs)
}