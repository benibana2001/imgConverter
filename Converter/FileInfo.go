package Converter

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Base DirType
	Dist DirType
}

type DirType struct {
	DirName   string
	Extension string
	FilePaths []string
}


func createFileInfo() *FileInfo {
	a := FileInfo{
		DirType{},
		DirType{},
	}
	a.setArgs()

	// 変換元ディレクトリ内の全てのファイルのパスをセット
	ss, err := walkFilePath(a.Base.DirName)
	if err != nil {
		fmt.Println(err)
	}
	a.Base.FilePaths = ss

	return &a
}

func (a *FileInfo) setArgs() {
	//　フラグをセット
	var (
		baseExtension = flag.String("base", "jpg", "base extension")
		distExtension = flag.String("dist", "png", "dist extension")
	)

	// フラグをパース
	flag.Parse()

	// 変換元ファイルの情報をセットする
	a.Base.DirName = flag.Arg(0)
	a.Base.Extension = *baseExtension

	// 変換後ファイルの情報をセットする
	a.Dist.DirName = flag.Arg(1)
	a.Dist.Extension = *distExtension

}

// ディレクトリ内にある全てのファイルのパスを取得する
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
	fmt.Println(s)
	return s, nil
}
