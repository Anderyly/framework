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
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

var _ Upload = (*upload)(nil)

type Upload interface {
	Ext(value []string) Upload
	Get(file *multipart.FileHeader) (int, string)
	Path(value string) Upload
}

type upload struct {
	Suffix  []string
	Address string
}

func NewUpload() Upload {
	return &upload{}
}

func (con *upload) Ext(value []string) Upload {
	return &upload{
		Suffix:  value,
		Address: con.Address,
	}
}

func (con *upload) Path(value string) Upload {
	return &upload{
		Suffix:  con.Suffix,
		Address: value,
	}
}

func (con *upload) Get(file *multipart.FileHeader) (int, string) {

	log.Println(con)
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
		return 400, "上传失败!只允许" + suffix + "文件"
	}

	fileName := NewStr().Md5(fmt.Sprintf("%s%s", file.Filename, time.Now().String()))
	fildDir := fmt.Sprintf("static/upload/%d-%d/", time.Now().Year(), time.Now().Month())
	if con.Address != "" {
		fildDir = fmt.Sprintf(con.Address, time.Now().Year(), time.Now().Month())
	}

	err := NewDir().Create(fildDir)
	if err != nil {
		return 400, err.Error()
	}

	filePath, err := os.Create(fildDir + fileName + fileExt)
	if err != nil {
		return 400, err.Error()
	}

	data, err := file.Open()
	if err != nil {
		return 400, err.Error()
	}

	defer data.Close()

	var context []byte = make([]byte, 1024)

	for {
		n, err := data.Read(context)
		filePath.Write(context[:n])
		if err != nil {
			if err == io.EOF {
				return 200, fildDir + fileName + fileExt
			} else {
				return 400, err.Error()
			}
		}
	}

}
