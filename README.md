# unimail-go-sdk

> The current SDK version is v2. If you need to use the previous v1 version, please switch to the v1 branch.

This is a Go SDK for Unimail. Quickly integrate into your project.

[中文文档](README_zh.md)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [unimail-go-sdk](#unimail-go-sdk)
  - [usage](#usage)
  - [api docs](#api-docs)
  - [support language](#support-language)

<!-- /code_chunk_output -->

## usage

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

```go
    // Gen req
	req := UnimailReq{
		// From:        "Notice",
		Receivers:   []string{"receiver"},
    	// Cc:          "",
		// Bcc:         "",
		Subject:     "test email", //
		TxtContent:  "common attachment email test", // htmlContext和txtContext二选一即可
		HtmlContent: "<h1>common attachment email test this is a test <span style=\"font-size: 20px\">email 2</span></h1>",
	}
    // add file attachment
	req.AppendFile("attach1.txt", "./textAttachment.txt")
	// add uri resources
    req.AppendUri("attach2.txt", "https://text.com/attach2.txt")
    // add io.Reader
	// req.AppendAttachment(&EmailAttachment{
	// 	Name: "filename.txt",
	// 	FileAttachment: io.Reader,
	// })

    // send email
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

## support language

chinese is the default language for the sdk.

- [x] english (en)
- [x] simple chinese (zh)
- [x] vietnamese (vi)
- [x] idonesian (id)
- [x] thai (th)
- [x] gujarati (gu)

if you want to support other language, please open a issue.

- tips

> If you want to use this SDK, please contact the author via i-curve@qq.com.
