package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var allowExt = []string{".jpg", ".png", ".gif", ".jpeg"}

type FilePara struct {
	Width    int
	Height   int
	FileName string
	Path     string
	Url      string
}

func GetFileName(fileName string) (filename string, err error) {
	ext := path.Ext(fileName)
	ext = strings.ToLower(ext)
	if ext != ".jpg" && ext != ".png" && ext != ".gif" && ext != ".jpeg" {
		//return "", errors.New("文件类型错误")
	}
	if len(ext) == 0 {
		ext = ".jpg"
	}
	filename = time.Now().Format("20060102") + strconv.Itoa(int(time.Now().UnixNano())) + ext
	return filename, nil
}

func GetFileNameFix(fileName string) (filename string, err error) {
	ext := path.Ext(fileName)
	ext = strings.ToLower(ext)
	if ext != ".jpg" && ext != ".png" && ext != ".gif" && ext != ".jpeg" {
		//return "", errors.New("文件类型错误")
	}
	if len(ext) == 0 {
		ext = ".jpg"
	}
	filename = fileName + ext
	return filename, nil
}

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

func IsDir(path string) error {
	b, err := PathExists(path)
	if err != nil {
		return err
	}
	if !b {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

//保存多个文件 如果key为空 刚默认为 files
func SaveMultipartFile(c *gin.Context, key string, imgPath string) (fileNameArr []string, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	files := form.File[key]
	if len(files) > 0 {
		for _, img := range files {
			fname, err := GetFileName(img.Filename)
			if err != nil {
				fmt.Println(err)
			}
			//检查目录是否存在  不存在则创建
			if err = IsDir(imgPath); err != nil {
				fmt.Println(errors.New("创建文件夹失败"+imgPath), err)
				return nil, err
			}
			if err = c.SaveUploadedFile(img, imgPath+fname); err != nil {
				fmt.Println(errors.New("文件保存失败"), err)
				return nil, err
			}
			fileNameArr = append(fileNameArr, fname)
		}
	}
	return
}

func SaveOneFileReplace(c *gin.Context, key string, imgPath string, uid int64) (fileName string, err error) {
	img, err := c.FormFile(key)
	if err != nil {
		fmt.Println("SaveOneFile-", err)
		return
	}
	if img == nil || img.Size < 10 {
		return
	}
	fileName = Int64ToString(uid) + ".jpg"
	//根据电话获取现有的图片名称
	exist, s := GetFileByShellExist(imgPath, uid)
	if exist {
		fileName = s
	}
	//检查目录是否存在  不存在则创建
	if err = IsDir(imgPath); err != nil {
		fmt.Println(errors.New("创建文件夹失败"+imgPath), err)
		return
	}
	if err = c.SaveUploadedFile(img, imgPath+fileName); err != nil {
		fmt.Println(errors.New("文件保存失败"), err)
		return
	}
	return
}

func UploadOneFileReplaceNew(c *gin.Context, imgPath string, uid int64) (fileName string, err error) {
	img, err := c.FormFile("file")
	if err != nil {
		fmt.Println("SaveOneFile-", err)
		return
	}
	if img == nil || img.Size < 10 {
		return
	}
	fileName = Int64ToString(uid) + ".jpg"
	//根据电话获取现有的图片名称
	exist, s := GetFileByShellExist(imgPath, fileName)
	if exist {
		fileName = s
	}
	//检查目录是否存在  不存在则创建
	if err = IsDir(imgPath); err != nil {
		fmt.Println(errors.New("创建文件夹失败"+imgPath), err)
		return
	}
	if err = c.SaveUploadedFile(img, imgPath+fileName); err != nil {
		fmt.Println(errors.New("文件保存失败"), err)
		return
	}
	return
}

func SaveOneFile(c *gin.Context, key string, imgPath string) (fileName string, err error) {
	img, err := c.FormFile(key)
	if err != nil {
		fmt.Println("SaveOneFile-", err)
		return
	}
	if img == nil || img.Size < 10 {
		return
	}
	fileName, err = GetFileName(img.Filename)
	if err != nil {
		fmt.Println(err)
	}
	//检查目录是否存在  不存在则创建
	if err = IsDir(imgPath); err != nil {
		fmt.Println(errors.New("创建文件夹失败"+imgPath), err)
		return
	}
	if err = c.SaveUploadedFile(img, imgPath+fileName); err != nil {
		fmt.Println(errors.New("文件保存失败"), err)
		return
	}
	return
}

func SaveOneFileUser(c *gin.Context, key string, imgPath string, filename string) (fileName string, err error) {
	img, err := c.FormFile(key)
	if err != nil {
		fmt.Println("SaveOneFile-", err)
		return
	}
	if img == nil || img.Size < 10 {
		return
	}
	fileName, err = GetFileNameFix(filename)
	if err != nil {
		fmt.Println(err)
	}
	//检查目录是否存在  不存在则创建
	if err = IsDir(imgPath); err != nil {
		fmt.Println(errors.New("创建文件夹失败"+imgPath), err)
		return
	}
	if err = c.SaveUploadedFile(img, imgPath+fileName); err != nil {
		fmt.Println(errors.New("文件保存失败"), err)
		return
	}
	return
}

//
//func SaveOneFileToThumb(c *gin.Context, key string, para *FilePara) (fileName string, err error) {
//	img, err := c.FormFile(key)
//	if err != nil {
//		fmt.Println("SaveOneFile-", err)
//		return
//	}
//	if img == nil || img.Size < 10 {
//		return
//	}
//	name, err := GetFileName(img.Filename)
//	if err != nil {
//		fmt.Println(err)
//	}
//	para.FileName = "sm_" + name
//	//检查目录是否存在  不存在则创建
//	if err = IsDir(para.Path); err != nil {
//		fmt.Println(errors.New("创建文件夹失败"+para.Path), err)
//		return
//	}
//	if err := SaveThumbnailByFile(img, para); err != nil {
//		fmt.Println(err)
//		return "", err
//	}
//	return para.FileName, err
//}

func SaveFileByByte(bytes []byte, imgPath string, filename string) (err error) {
	//检查目录是否存在  不存在则创建
	if err = IsDir(imgPath); err != nil {
		fmt.Println(errors.New("创建文件夹失败"+imgPath), err)
		return err
	}
	outputPath := imgPath + filename
	err = ioutil.WriteFile(outputPath, bytes, 0400)
	if err != nil {
		fmt.Printf("error writing out resized image, %s\n", err)
		fmt.Println(err)
		return err
	}
	return
}

type RspImage struct {
	Id   int64  `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type ImageSize struct {
	Width  int
	Height int
}

func UploadOneReplace(c *gin.Context, path string, uid int64) (name string) {
	fileName, err := SaveOneFileReplace(c, "file", path, uid)
	if err != nil {
		fmt.Println(err)
		RspError(c, err.Error())
		return
	}
	return fileName
}

//
//func UploadOne(c *gin.Context, root string, middle string) (res *RspImage) {
//	fileName, err := util.SaveOneFile(c, "file", root+middle)
//	if err != nil {
//		fmt.Println(err)
//		util.RspError(c, err.Error())
//		return
//	}
//	path := middle + fileName
//	return &RspImage{Path: path, Name: fileName}
//}

func UploadOne(c *gin.Context, path string) (res *RspImage) {
	fileName, err := SaveOneFile(c, "file", path)
	if err != nil {
		fmt.Println(err)
		RspError(c, err.Error())
		return
	}
	path = path + fileName
	return &RspImage{Name: fileName}
}

func UploadOneUser(c *gin.Context, uid int64, path string) (res *RspImage) {
	filename := Int64ToString(uid)
	fileName, err := SaveOneFileUser(c, "file", path, filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := path + fileName
	return &RspImage{Url: url, Name: fileName}
}
