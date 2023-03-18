/*
 * @author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc
 * @copyright Copyright (c) 2022
 *
 */

package lib

import (
	"io"
	"os"
)

var _ File = (*file)(nil)

type File interface {
	IsFile(filename string) bool                                // 判断文件是否存在
	Write(filename string, data []byte, perm os.FileMode) error // 文件写入
}

type file struct {
}

func NewFile() File {
	return &file{}
}

func (con *file) IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (con *file) Write(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
