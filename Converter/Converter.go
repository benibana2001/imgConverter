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
	FileInfo *FileInfo
	Files    []*os.File
	Imgs     []*image.Image
}


// インスタンスを作成
func NewConverter() Converter {

	c := Converter{
		FileInfo: createFileInfo(),
	}

	// 全ファイルのio.Writerを作成する
	var files []*os.File
	for _, path := range c.FileInfo.Base.FilePaths {
		f, err := os.Open(path)
		files = append(files, f)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to open file : ", path)
			os.Exit(6)
		}
	}
	c.Files = files

	return c
}

// 画像を読み込み
func (c *Converter) Decode() {
	var imgs []*image.Image

	// jpeg を png へ変換
	if c.FileInfo.Dist.Extension == "png" {
		for _, file := range c.Files {
			img, _ := jpeg.Decode(file)
			imgs = append(imgs, &img)
		}
	}

	// png を jpeg へ変換
	if c.FileInfo.Dist.Extension == "jpg" {
		for _, file := range c.Files {
			img, _ := png.Decode(file)
			imgs = append(imgs, &img)
		}
	}
	c.Imgs = imgs
}

// 画像を出力
func (c *Converter) Encode() {
	dist := c.FileInfo.Dist.DirName
	if err := os.Mkdir(dist, 0777); err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	extPng := ".png"
	extJpg := ".jpg"

	// jpeg to png
	if c.FileInfo.Dist.Extension == "png" {
		for i, img := range c.Imgs {
			f, _ := os.Create(dist + "/" + getFileNameWithoutExt(c.FileInfo.Base.FilePaths[i]) + extPng)
			fmt.Println(dist + "/" + getFileNameWithoutExt(c.FileInfo.Base.FilePaths[i]) + extPng)

			err := png.Encode(f, *img)

			if err != nil {
				fmt.Printf("type is %T\n", img)
				fmt.Println(err)
				os.Exit(3)
			}
		}
	}

	// png to jpeg
	if c.FileInfo.Dist.Extension == "jpg" {
		fmt.Println("c.Imgs is ", c.Imgs)
		for i, img := range c.Imgs {
			f, _ := os.Create(dist + "/" + getFileNameWithoutExt(c.FileInfo.Base.FilePaths[i]) + extJpg)
			fmt.Println(dist + "/" + getFileNameWithoutExt(c.FileInfo.Base.FilePaths[i]) + extJpg)


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
