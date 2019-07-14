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
	Ext   string
	Paths []string
	Files []*os.File
	Imgs  []*image.Image
}

// ファイルパスを読み込み
func walkFilePath(dirname string) ([]string, error) {
	var s []string

	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".png" {
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
		Ext:  os.Args[3],
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
func (c *Converter) Decode() {
	var imgs []*image.Image

	// jpeg を png へ変換
	if c.Ext == "png" {
		for _, file := range c.Files {
			img, _ := jpeg.Decode(file)
			imgs = append(imgs, &img)
		}
	}

	// png を jpeg へ変換
	if c.Ext == "jpg" {
		for _, file := range c.Files {
			img, _ := png.Decode(file)
			imgs = append(imgs, &img)
		}
	}
	c.Imgs = imgs
}

// 画像を出力
func (c *Converter) Encode() {
	dist := c.Dist
	if err := os.Mkdir(dist, 0777); err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	extPng := ".png"
	extJpg := ".jpg"

	// jpeg to png
	if c.Ext == "png" {
		for i, img := range c.Imgs {
			f, _ := os.Create(dist + "/" + getFileNameWithoutExt(c.Paths[i]) + extPng)
			fmt.Println(dist + "/" + getFileNameWithoutExt(c.Paths[i]) + extPng)

			err := png.Encode(f, *img)

			if err != nil {
				fmt.Printf("type is %T\n", img)
				fmt.Println(err)
				os.Exit(3)
			}
		}
	}

	// png to jpeg
	if c.Ext == "jpg" {
		fmt.Println("c.Imgs is ", c.Imgs)
		for i, img := range c.Imgs {
			f, _ := os.Create(dist + "/" + getFileNameWithoutExt(c.Paths[i]) + extJpg)
			fmt.Println(dist + "/" + getFileNameWithoutExt(c.Paths[i]) + extJpg)


			options := jpeg.Options{Quality: 100}
			err := jpeg.Encode(f, *img, &options)

			if err != nil {
				fmt.Printf("type is %T\n", img)
				fmt.Println(err)
				os.Exit(3)
			}
		}
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
