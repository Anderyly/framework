/*
 * @author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc
 * @copyright Copyright (c) 2022
 *
 */

package lib

import "os"

var _ File = (*file)(nil)

type File interface {
	IsFile(filename string) bool
	Upload()
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

func (con *file) Upload() {

}
