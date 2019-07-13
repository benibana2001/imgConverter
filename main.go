package main

import "github.com/benibana2001/imgConverter/Converter"

func main() {
	c := Converter.NewConverter("jpeg")
	c.DecodeJpeg()
	c.EncodePng("output")
}
