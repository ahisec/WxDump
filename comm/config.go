package comm

//V1.0.0
//By cmluZw 2024-08-12

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"golang.org/x/sys/windows"
)

//搬运:https://github.com/LC044/WeChatMsg/blob/master/app/decrypt/version_list.json

//go:embed WX_OFFS.json
var versionListData []byte

var PROCESS_ALL_ACCESS = uint32(
	windows.PROCESS_QUERY_INFORMATION |
		windows.PROCESS_VM_READ |
		windows.PROCESS_VM_WRITE |
		windows.PROCESS_VM_OPERATION |
		windows.PROCESS_CREATE_THREAD |
		windows.PROCESS_DUP_HANDLE |
		windows.PROCESS_TERMINATE |
		windows.PROCESS_SUSPEND_RESUME |
		windows.PROCESS_SET_QUOTA |
		windows.PROCESS_SET_INFORMATION |
		windows.PROCESS_QUERY_LIMITED_INFORMATION)

// 支持自动获取数据的版本号
var SupportAutoGetDataVersionList = []string{
	//"3.9.10.16",
	"3.9.12.17",
}

type versionList struct {
	Versions map[string][]int
}

//微信名
//0
//手机号

func Get_version_list() map[string][]int {
	// 定义一个map变量用于存储解析后的JSON数据
	var versionList map[string][]int

	// 解析嵌入的JSON数据到map
	err := json.Unmarshal(versionListData, &versionList)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return nil
	}
	return versionList
	// 打印versionList查看结果
	//fmt.Printf("%+v\n", versionList["2.4.5.1"])
}

////从version_list.json中读取版本和偏移量
//
//func Get_version_list() map[string][]int {
//	// 设置JSON文件路径
//	VERSION_LIST_PATH := "./comm/version_list.json"
//
//	// 读取JSON文件内容
//	fileContent, err := ioutil.ReadFile(VERSION_LIST_PATH)
//	if err != nil {
//		fmt.Printf("Error reading file '%s': %v\n", VERSION_LIST_PATH, err)
//		return nil
//	}
//
//	// 定义一个map变量用于存储解析后的JSON数据
//	var versionList map[string][]int
//
//	// 解析JSON数据到map
//	err = json.Unmarshal(fileContent, &versionList)
//	if err != nil {
//		fmt.Printf("Error decoding JSON: %v\n", err)
//		return nil
//	}
//	return versionList
//	// 打印versionList查看结果
//	//fmt.Printf("%+v\n", versionList["2.4.5.1"])
//}
