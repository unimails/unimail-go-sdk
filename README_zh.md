# unimail-go-sdk

> 当前sdk的版本是v2, 如果你需要用以前的v1版本, 请切换v1分支

unimail 的 go 语言 sdk, 快速集成到你的项目

[english docs](README.md)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [unimail-go-sdk](#unimail-go-sdk)
  - [使用](#使用)
  - [api docs](#api-docs)
  - [支持的语言](#支持的语言)

<!-- /code_chunk_output -->

## 使用

- 安装

```shell
go get github.com/i-curve/unimail-go-sdk
```

- 初始化客户端

你需要申请一个 key

main.go

```go
package main

import (
	"fmt"

	unimailGo "github.com/i-curve/unimail-go-sdk"
)

// 输入你的key
var key = ""

func main() {
	client := unimailGo.New(key)

    // 检查客户端连接
	status := client.CheckConnect()
	fmt.Println(client, status)
}
```

- 发邮件

```go
    // 构造请求体
	req := UnimailReq{
		// From:        "通知",
		Receivers:   []string{"收件人,数组, 可以多个收件人"},
    	// Cc:          "",
		// Bcc:         "",
		Subject:     "test email", // 主题
		TxtContent:  "common attachment email test", // htmlContext和txtContext二选一即可
		HtmlContent: "<h1>common attachment email test this is a test <span style=\"font-size: 20px\">email 2</span></h1>",
	}
    // 添加文件附件
	req.AppendFile("attach1.txt", "./textAttachment.txt")
	// 添加http资源文件
    req.AppendUri("attach2.txt", "https://text.com/attach2.txt")
    // 添加 io.Reader 资源文件
	// req.AppendAttachment(&EmailAttachment{
	// 	Name: "filename.txt",
	// 	FileAttachment: io.Reader,
	// })

    // 发送邮件
    result := client.SendEmail(req)

    if result.IsSucess() {
        fmt.Println("send email success")
    } else {
        fmt.Printf("send email fail, error: %s\n", result.GetMsg())
    }
```

## api docs

1. New(key string) UnimailClient

init a client by key

2. NewByStruct(key keyStruct) UnimailClient

keyStruct is a struct that implements GetKey function

3. NewUnimailClient(client newUnimailClient) UnimailClient

newUnimailClient is a struct that implements GetKey, GetHost function

4. client.SetLanguage(language string) error

set language for the client,default is zh

5. client.CheckConnect() bool

check the host and key is ok

6. client.SendEmail(UnimailReq) Result

please see usage

## 支持的语言

sdk 默认返回的 msg 为中文

- [x] english (en)
- [x] simple chinese (zh)
- [x] vietnamese (vi)
- [x] idonesian (id)
- [x] thai (th)
- [x] gujarati (gu)

如果你需要添加了更多语言，欢迎提 issue

- 提示

如果想要使用 unimail sdk, 请通过邮件联系作者 i-curve@qq.com
