package main

import (
	"github.com/benibana2001/imgConverter/Converter"
)

func main() {
	// 変換用インスタンスを作成
	c := Converter.NewConverter()
	// 変換処理を実行
	c.Convert()
}
