# unimail-go-sdk

This is a Go SDK for Unimail. Quickly integrate into your project

[中文文档](README_zh.md)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [unimail-go-sdk](#unimail-go-sdk)
  - [simple usage](#simple-usage)
  - [api docs](#api-docs)
  - [support language](#support-language)

<!-- /code_chunk_output -->

## simple usage

- install

```shell
go get github.com/i-curve/unimail-go-sdk
```

- init a unimail client

you need a authorization key.

main.go

```go
package main

import (
	"fmt"

	unimailGo "github.com/i-curve/unimail-go-sdk"
)

// please input your key here
var key = ""

func main() {
	client := unimailGo.New(key)

    // check the client connection
	status := client.CheckConnect()
	fmt.Println(client, status)
}
```

- send email

example
receiver: aaa@gmail.com  
email subject: email subject  
email content: this is a email content

```go
    result := client.SendEmail("aaa@gmail.com", "email subject", "this is a email content")

    if result.IsSucess() {
        fmt.Println("send email success")
    } else {
        fmt.Printf("send email fail, error: %s\n", result.GetMsg())
    }
```

- batch send email

example
receivers: aaa@gmail.com,bbb@gmail.com  
email subject: email subject  
email content: this is a email content

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

## support language

chinese is the default language for the sdk.

- [x] english (en)
- [x] simple chinese (zh)
- [x] vietnamese (vi)
- [x] idonesian (id)
- [x] thai (th)
- [x] gujarati (gu)

if you want to support other language, please open a issue.
