/*
 * @author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc
 * @copyright Copyright (c) 2022
 *
 */

package lib

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

var _ Upload = (*upload)(nil)

type Upload interface {
	Get(file *multipart.FileHeader) (bool, string)
	Set(path string, ext []string, size int64) Upload
}

type upload struct {
	Suffix  []string
	Address string
	Size    int64
}

func NewUpload() Upload {
	return &upload{}
}

func (con *upload) Set(path string, ext []string, size int64) Upload {
	return &upload{
		Suffix:  ext,
		Address: path,
		Size:    size,
	}
}

func (con *upload) Get(file *multipart.FileHeader) (bool, string) {

	fileExt := strings.ToLower(path.Ext(file.Filename))

	suffix := ""
	isExt := false
	for _, v := range con.Suffix {
		if "."+v == fileExt {
			isExt = true
		}
		suffix += v + ","
	}

	if !isExt {
		return false, "上传失败!只允许" + suffix + "文件"
	}

	if file.Size > con.Size {
		return false, "上传文件过大"
	}

	fileName := NewStr().Md5(fmt.Sprintf("%s%s", file.Filename, time.Now().String()))
	fileDir := fmt.Sprintf("static/upload/%d-%d/", time.Now().Year(), time.Now().Month())
	if con.Address != "" {
		fileDir = fmt.Sprintf(con.Address, time.Now().Year(), time.Now().Month())
	}

	err := NewDir().Create(fileDir)
	if err != nil {
		return false, err.Error()
	}

	filePath, err := os.Create(fileDir + fileName + fileExt)
	if err != nil {
		return false, err.Error()
	}

	data, err := file.Open()
	if err != nil {
		return false, err.Error()
	}

	defer data.Close()

	var context []byte = make([]byte, 1024)

	for {
		n, err := data.Read(context)
		filePath.Write(context[:n])
		if err != nil {
			if err == io.EOF {
				return false, fileDir + fileName + fileExt
			} else {
				return false, err.Error()
			}
		}
	}

}
