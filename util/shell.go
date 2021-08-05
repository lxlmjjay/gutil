package util

import (
	"bytes"
	"errors"
	"os/exec"
	"reflect"
	"strings"
)

const ShellToUse = "bash"

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func GetFileByShellExist(path string, key interface{}) (bool, string) {
	username := ""
	of := reflect.TypeOf(key)
	if of.String() == "int64" {
		username = Int64ToString(key.(int64))
	} else if of.String() == "string" {
		username = key.(string)
	}
	//获取头像
	//fmt.Println(`ls ` + dao.Conf.System.ImgHeader + `/` + username + `* |sed 's#.*/##'`)
	err, s, s2 := Shellout(`ls ` + path + `/` + username + `* |sed 's#.*/##'`)
	if err != nil || len(s2) > 0 {
		//log.Debug.Println(username, " 头像获取错误,", err, s2)
		return false, "default.jpg"
	} else {
		filename := strings.Replace(s, "\n", "", -1)
		return true, filename
	}
}

func GetFileByShell(path string, key interface{}) string {
	username := ""
	of := reflect.TypeOf(key)
	if of.String() == "int64" {
		username = Int64ToString(key.(int64))
	} else if of.String() == "string" {
		username = key.(string)
	}
	//获取头像
	//fmt.Println(`ls ` + dao.Conf.System.ImgHeader + `/` + username + `* |sed 's#.*/##'`)
	err, s, s2 := Shellout(`ls ` + path + `/` + username + `* |sed 's#.*/##'`)
	if err != nil || len(s2) > 0 {
		//log.Debug.Println(username, " 头像获取错误,", err, s2)
		return "default.jpg"
	} else {
		filename := strings.Replace(s, "\n", "", -1)
		return filename
	}
}

func GetFileByShellArr(path string, key interface{}) (arr []string, err error) {
	username := ""
	of := reflect.TypeOf(key)
	if of.String() == "int64" {
		username = Int64ToString(key.(int64))
	} else if of.String() == "string" {
		username = key.(string)
	}
	//获取头像
	//fmt.Println(`ls ` + path + `/` + username + `* |sed 's#.*/##'`)
	err, s, s2 := Shellout(`ls ` + path + `/` + username + `.jpg* |sed 's#.*/##'`)
	if err != nil || len(s2) > 0 {
		err, s, s2 = Shellout(`ls ` + path + `/` + username + `* |sed 's#.*/##'`)
		if err != nil || len(s2) > 0 {
			//log.Debug.Println(username, " 头像获取错误,", err, s2)
			err = errors.New("图片获取错误")
			return
		} else {
			arr = strings.Split(s, "\n")
			return
		}
		//log.Debug.Println(username, " 头像获取错误,", err, s2)
	} else {
		arr = strings.Split(s, "\n")
		return
	}
	return
}
