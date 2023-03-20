/*
 * @author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc
 * @copyright Copyright (c) 2022
 *
 */

package lib

import (
	"os"
	"path/filepath"
)

var _ Dir = (*dir)(nil)

type Dir interface {
	IsPath(path string) (bool, error) // 判断路径是否存在
	Create(filePath string) error     // 创建文件夹 注意 linux系统/为根目录
	Get(localDir string) int          // 获取文件夹下文件数量
}

type dir struct {
}

func NewDir() Dir {
	return &dir{}
}

func (con *dir) IsPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (con *dir) Create(filePath string) error {
	if ok, _ := con.IsPath(filePath); !ok {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (con *dir) Get(localDir string) int {
	var num int
	filepath.Walk(localDir, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		num++
		return nil
	})
	return num
}
