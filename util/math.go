package util

import (
	"math/rand"
)

//求一个二维数组 的笛卡尔积
func Cartesian(arrs [][]interface{}) {
	//lens := func(i int) int { return len(arrs[i]) }
	//var n []string
	//for ix := make([]int, len(arrs)); ix[0] < lens(0); NextIndex(ix, lens) {
	//	var r []*dbGoods.GenSku
	//	var x string
	//	for j, k := range ix {
	//		r = append(r, arrs[j][k])
	//		x += arrs[j][k].SaleAttrValueName
	//	}
	//	n = append(n, x)
	//	fmt.Println(r)
	//	fmt.Println(n)
	//}
}

func NextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

//如 80 显示 8折 
func IntToSale(i int) float64 {
	return float64(i) / float64(10)
}

//百分比
func IntToPercent(i int) float64 {
	return float64(i) / float64(100)
}

//const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterBytes = "0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits,as many as letterIdxBits
)

func GenRand(n int) string {
	b := make([]byte,n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}