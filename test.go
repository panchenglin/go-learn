package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

func show_substr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}
		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func main() {
	// 定义几个变量，用于接收命令行的参数值
	var dir = "D:/test"
	err := os.Chdir(dir) //存在的目录
	if err == nil {
		lateDir, _ := os.Getwd()
		//获取文件或目录相关信息
		fileInfoList, err := ioutil.ReadDir(lateDir)
		if err != nil {
			log.Fatal(err)
		}
		for i := range fileInfoList {
			//打印当前文件或目录下的文件或目录名
			var fileName = fileInfoList[i].Name()
			nameStr := strings.Fields(fileName)
			preName := nameStr[0]
			//preName := show_substr(fileName, 9)
			firDir := lateDir + "/" + "丁二烯装置-随机资料-" + preName + "-上海日机装"
			fileDirName := lateDir + "/" + fileName
			if !IsFile(fileDirName) || strings.LastIndex(fileName, ".pdf") == -1 {
				continue
			}
			dirBool, _ := PathExists(firDir)
			if dirBool != true {
				os.Mkdir(firDir, os.ModePerm)
			}
			// 移动文件
			os.Rename(fileDirName, firDir+"/"+fileName)
		}
	} else {
		fmt.Println("error:", err)
	}
}
