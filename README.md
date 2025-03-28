# WxDump
## 前提
- 一个平平无奇的PC版微信聊天记录导出工具-红队命令行版 <br>
- 红队工具,微信聊天记录导出,微信聊天备份,数据解密，支持最新版微信<br>
- 无图形化界面，不需要用户交互，一键运行，导出无忧<br>
- 因为github上很多都不适用于最新的微信版本，于是产生了WxDump<br>
## 安装
```
go mod download
go build -o WxDump.exe
```

## 使用
```
WxDump.exe
```
在命令行运行后输出key和Decryption Success

![img.png](img.png)

并在当前目录下生成tmp,decrypted,key.txt文件，如下

![img_1.png](img_1.png)

<b>分别对应</b>
- key.txt 解密所需要的key，从内存中读取
- tmp 解密前的聊天记录
- decrypted 解密后的聊天记录

将decrypted下的*.db放入navicat，使用语句
```go
SELECT * FROM "MSG" WHERE StrContent  like'%密码%' 
```
快速搜索密码

## 图形化查看聊天记录
在右侧有写好的UI查看聊天记录的工具：WxDump-UI.exe
注意使用前提
- 确保WxDump-UI.exe与解密后的*.db在同一目录
- 确保当前目录下存在MicroMsg.db和MSG*.db
- 确保这些*.db是解密后的db

由于是fyne初学者，所以前端代码就不开源了，双击WxDump-UI.exe即可运行<br>
界面如下:
![img_2.png](img_2.png)

- 用法请自行探究，我觉得页面还是比较通俗易懂

## 参考
写工具的时候参考了以下几位大佬的项目
- https://github.com/LC044/WeChatMsg
- https://github.com/SpenserCai/GoWxDump
- https://github.com/AdminTest0/SharpWxDump?tab=readme-ov-file
- https://github.com/xaoyaoo/PyWxDump
## 免责声明
这很重要，仅供合法用途，请勿滥用

## 反馈和建议
- 提交issue
- 联系作者，QQ:2534395766