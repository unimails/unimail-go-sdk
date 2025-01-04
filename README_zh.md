# unimail-go-sdk

unimail 的 go 语言 sdk, 快速集成到你的项目

[english docs](README.md)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [unimail-go-sdk](#unimail-go-sdk)
  - [简单使用](#简单使用)
  - [api docs](#api-docs)
  - [支持的语言](#支持的语言)

<!-- /code_chunk_output -->

## 简单使用

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

例如
收件人: aaa@gmail.com  
邮件标题: email subject  
邮件正文: this is a email content

```go
    result := client.SendEmail("aaa@gmail.com", "email subject", "this is a email content")

    if result.IsSucess() {
        fmt.Println("send email success")
    } else {
        fmt.Printf("send email fail, error: %s\n", result.GetMsg())
    }
```

- 批量发送邮件

例如
收件人: aaa@gmail.com,bbb@gmail.com  
邮件标题: email subject  
邮件正文: this is a email content

```go
	bresult := client.BatchSendEmail([]string{"aaa@gmail.com", "bbb@gmail.com"}, "email subject", "this is a email content")

	if bresult.IsSucess() {
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

6. client.SendEmail(receiver string, subject string, content string) Result

send email to receiver. if you have many receiver, you can concat the receiver by ";" or use BatchSendEmail

7. client.BatchSendEmail(receivers []string, subject string, content string) Result

like SendEmail, but receivers is a slice

## 支持的语言

sdk 默认返回的 msg 为中文

- [x] english (en)
- [x] simple chinese (zh)
- [x] vietnamese (vi)
- [x] idonesian (id)
- [x] thai (th)
- [x] gujarati (gu)

如果你需要添加了更多语言，欢迎提 issue
