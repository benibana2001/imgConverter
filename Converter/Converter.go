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
	Dir   string
	Dist  string
	Paths []string
	Files []*os.File
	Imgs  []*image.Image
}

// ファイルパスを読み込み
func walkFilePath(dirname string) ([]string, error) {
	var s []string

	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
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
func NewConverter() Converter {
	c := Converter{
		Dir:  os.Args[1],
		Dist: os.Args[2],
	}

	// ディレクトリ内の全ファイルのパスを取得する
	paths, err := walkFilePath(c.Dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	c.Paths = paths

	// 全ファイルのio.Writerを作成する
	var files []*os.File
	for _, path := range paths {
		f, _ := os.Open(path)
		files = append(files, f)
	}
	c.Files = files

	return c
}

// 画像を読み込み
func (c *Converter) DecodeJpeg() {
	var imgs []*image.Image
	for _, file := range c.Files {
		img, _ := jpeg.Decode(file)
		imgs = append(imgs, &img)
	}
	c.Imgs = imgs
}

// 画像を出力
func (c *Converter) EncodePng() {
	dist := c.Dist
	if err := os.Mkdir(dist, 0777); err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	for i, img := range c.Imgs {
		f, _ := os.Create(dist + "/" + getFileNameWithoutExt(c.Paths[i]) + ".png")
		fmt.Println(dist + "/" + getFileNameWithoutExt(c.Paths[i]) + ".png")
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
