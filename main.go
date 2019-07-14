package main

import (
	"github.com/benibana2001/imgConverter/Converter"
)

func main() {
	// 渡された引数を取得
	c := Converter.NewConverter()
	c.Decode()
	c.Encode()
}
