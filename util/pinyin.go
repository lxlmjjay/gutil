package util

import (
	"github.com/mozillazg/go-pinyin"
)

func GetFirstLetter(l string) string {
	i := pinyin.Pinyin(l, pinyin.NewArgs())
	return i[0][0][:1]
}
