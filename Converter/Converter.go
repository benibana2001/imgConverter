package Converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Converter struct {
	Paths []string
	F     []*os.File
	Imgs   []*image.Image
}

// ファイルパスを読み込み
func WalkFilePath(name string) ([]string, error) {
	var s []string

	err := filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".jpg" {
			s = append(s, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return s, nil
}

// インスタンスを作成
func NewConverter(name string) Converter{

	// TEST
	s, err := WalkFilePath(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	//

	//f, err := os.Open(name)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	var fs []*os.File

	for _, path := range s {
		f, _ := os.Open(path)
		fs = append(fs, f)
	}

	return Converter{
		F:     fs,
		Paths: s,
	}
}

// 画像を読み込み
func (c *Converter) DecodeJpeg() {
	var imgs []*image.Image
	for _, file := range c.F {
		img, _ := jpeg.Decode(file)
		imgs = append(imgs, &img)
	}
	c.Imgs = imgs
}

// 画像を出力
func (c *Converter) EncodePng(dirname string) {
	if err := os.Mkdir(dirname, 0777); err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	for i, img := range c.Imgs {
		f, _ := os.Create(dirname + "/" + getFileNameWithoutExt(c.Paths[i]) + ".png")
		fmt.Println(dirname + "/" + getFileNameWithoutExt(c.Paths[i]) + ".png")
		if err := png.Encode(f, *img); err != nil {
			fmt.Printf("type is %T\n", img)
			fmt.Println(err)
			os.Exit(3)
		}
	}
}
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
