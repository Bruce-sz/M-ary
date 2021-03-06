package M_ary

import (
	"math"
	"strings"
)

var (
	ThirtySixStrs = "0123456789ABCDEFGHIJKLNMOPQRSTUVWXYZ"
)

//Dec2ThirtySix is decimal convert to 36-ary
func Dec2ThirtySix(d10 int64) string {
	h36pos := make([]int64, 0) //每36进制数的10进制位置
	var tmp int64 = d10
	for {
		h36 := tmp % 36
		i10 := tmp / 36
		if h36 >= 0 {
			h36pos = append(h36pos, h36)
			tmp = i10
		}
		if i10 < 36 {
			h36pos = append(h36pos, i10)
			break
		}
	}
	hex36str := ""
	for _, x := range h36pos {
		hex36str = ThirtySixStrs[x:x+1] + hex36str
	}
	return hex36str
}

//ThirtySix2Dec is  36-ary convert to  decimal
func ThirtySix2Dec(StrHex36 string) int64 {

	var Base36To10 int //每位36进制基数对应的10进制数据大小
	var Dec10 int64  //被返回10进制数据
	var Base36 string  //36进制基数

	var StrHex36Len int = len(StrHex36)

	for Index := 1; Index <= StrHex36Len; Index = Index + 1 {
		Base36 = StrHex36[Index-1 : Index]
		Base36To10 = strings.Index(ThirtySixStrs, Base36)
		Dec10 = Dec10 + int64(float64(Base36To10)*math.Pow(float64(36), float64(StrHex36Len-Index) ))

	}

	return int64(Dec10)
}
