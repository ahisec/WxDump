package main

import (
	"WxDump/comm"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var key = ""
	wxid := comm.Get_result()
	dbPath := comm.GetFilePath(wxid)
	addrLen := 8
	is_normal, nickName, account, key_tmp := comm.Get_info()
	if is_normal != true {
		key_tmp, err := comm.GetKey(dbPath, addrLen)
		if (err != nil) || (key_tmp == "None") {
			fmt.Println("Error:", err)
			return
		} else {
			fmt.Println("Key:", key_tmp)
		}
		key = key_tmp
	} else {
		fmt.Println("nickName:", nickName)
		fmt.Println("account:", account)
		fmt.Println("Key:", key_tmp)
		key = key_tmp
	}
	comm.CopyMsgDb(dbPath)
	comm.DecryptDb(key)

	// 将解密后的文件打包成 ZIP
	CurrentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	// decrypted 目录路径
	decryptedPath := filepath.Join(CurrentPath, "decrypted")

	// 读取 decrypted 目录下的所有文件
	files, err := os.ReadDir(decryptedPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	// 遍历文件
	for _, file := range files {
		if !file.IsDir() { // 只处理文件
			fmt.Println("Found file:", file.Name())
			comm.ZipFile(decryptedPath+"/"+file.Name(), decryptedPath+"/"+file.Name()+".zip")
		}
	}
}
