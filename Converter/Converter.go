package Converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

type Converter struct {
	Paths []string
	F     *os.File
	Img   *image.Image
}

// ファイルパスを読み込み
func GetFilePath(name string) []string {
	files, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	var s []string

	for _, file := range files {
		//.DS_Store を削除
		s = append(s, file.Name())
	}
	return s
}

// インスタンスを作成
func NewConverter(name string) Converter{
	s := GetFilePath("jpeg")
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return Converter{
		F:     f,
		Paths: s,
	}
}

// 画像を読み込み
func (c *Converter) DecodeJpeg() {
	img, err := jpeg.Decode(c.F)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	c.Img = &img
}

// 画像を出力
func (c *Converter) EncodePng(name string) {
	f, _ := os.Create(name)
	if err := png.Encode(f, *c.Img); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

}
