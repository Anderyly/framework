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
	"path/filepath"
)

var _ Dir = (*dir)(nil)

type Dir interface {
	IsFile(filename string) bool
	IsPath(path string) (bool, error)
	Create(filePath string) error
	Write(filename string, data []byte, perm os.FileMode) error
	Get(localDir string) int
}

type dir struct {
}

func NewDir() Dir {
	return &dir{}
}

func (con *dir) IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true

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
	if !con.IsFile(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (con *dir) Write(filename string, data []byte, perm os.FileMode) error {
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
