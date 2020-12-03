package gowheel

import (
	"fmt"
	"io"
	"net/url"
	"log"
	"regexp"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
获取程序运行路径
*/
func CurrentPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1) + "/"
}

/*
检测文件或者文件夹是否已经存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
//CopyFile 复制文件
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
/**
下载文件
 */
func DownloadFile(fileUrl string, path string, name string)(filename string,err error) {
	//校验文件保存路径是否正确
	var (
		match bool
		data []byte
		response *http.Response
	)
	match,err = regexp.MatchString(`^.+/$`,path)
	if err != nil{
		return
	}
	if !match{
		err = fmt.Errorf("Path must end with /")
		return
	}
	//保存目录不存在新建
	exists, _ := PathExists(path)
	if !exists {
		err =os.MkdirAll(path, os.ModePerm)   //创建多级目录
		if err != nil{
			return
		}
	}
	//保存文件存在直接返回
	filename = path + name
	exists, _ = PathExists(filename)
	if exists {
		return
	}
	//拆解链接
	_, err = url.Parse(fileUrl)
	if err != nil {
		log.Println("parse url failed:", fileUrl, err)
		return
	}
	//保存文件
	response, err = http.Get(fileUrl)
	defer func() {
		_ = response.Body.Close()
	}()
	if err != nil {
		log.Println("get file_url failed:", err)
		return
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", fileUrl, err)
		return
	}
	err= FilePutContents(filename, data)
	return
}

/*
写文件
 */
func FilePutContents(filename string, contents []byte)(err error) {
	var file *os.File
	file, err = os.Create(filename)
	defer func() {
		_ = file.Close()
	}()

	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}
	_,err = file.Write(contents)
	return
}
