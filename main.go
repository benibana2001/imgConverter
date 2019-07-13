package main

import (
	"fmt"
	"github.com/benibana2001/imgConverter/Converter"
)

func main() {
	c := Converter.NewConverter("jpeg/sample01.jpg")
	fmt.Println(c.Paths)
	c.DecodeJpeg()
	c.EncodePng("sample01.png")
}
