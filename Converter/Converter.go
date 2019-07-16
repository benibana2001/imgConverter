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

	return c
}

// 変換処理
func (c *Converter) Convert() {
	c.OpenFiles()
	c.Decode()
	c.Encode()
}

// 全てのファイルのio.Writerを作成する
func (c *Converter) OpenFiles() {
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
}

// 画像を読み込み
func (c *Converter) Decode() {
	var imgs []*image.Image

	for _, file := range c.Files {
		var img image.Image

		switch c.FileInfo.Base.Extension {
		case "jpg" :
			img, _ = jpeg.Decode(file)
		case "png" :
			img, _ = png.Decode(file)
		}

		imgs = append(imgs, &img)
	}

	c.Imgs = imgs
}

// 画像を出力
func (c *Converter) Encode() {
	if err := os.Mkdir(c.FileInfo.Dist.DirName, 0777); err != nil {
		fmt.Println(err)
		fmt.Println(c.FileInfo.Dist.DirName)
		os.Exit(5)
	}

	for i, img := range c.Imgs {
		newFile := c.FileInfo.Dist.DirName + "/" + getFileName(c.FileInfo.Base.FilePaths[i]) + "." + c.FileInfo.Dist.Extension
		f, _ := os.Create(newFile)
		fmt.Println(newFile)


		switch c.FileInfo.Dist.Extension {
		case "png" :
			err := png.Encode(f, *img)
			if err != nil {
				fmt.Println(err)
				os.Exit(3)
			}

		case "jpg" :
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

func getFileName(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
