// Package unimail is a golang sdk for unimail
package unimail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"strings"
)

// UnimailClient unimail client to operate unimail
type UnimailClient interface {
	//  CheckConnect 检查连接是否正常
	CheckConnect() bool

	// SetLanguage 设置语言 [zh, en, vi, th, gu, id]
	SetLanguage(lang string) error

	// SendEmail 发送邮件
	//
	// req UnimailReq 发送邮件请求体
	SendEmail(req UnimailReq) Result
}

type keyStruct interface {
	GetKey() string
}

// New 创建一个UnimailClient
//
// key Unimail 获取的已授权 key
func New(key string) UnimailClient {
	return &unimail{
		Host:   DOMAIN,
		Key:    key,
		lang:   "zh",
		client: &http.Client{Timeout: timeout},
	}
}

// NewByStruct 创建一个UnimailClient
//
// keyStruct 实现了 GetKey() string 的结构体, 用于获取已授权的 unimail key
func NewByStruct(key keyStruct) UnimailClient {
	if key == nil {
		panic("unimail key is nil")
	}
	return &unimail{
		Host:   DOMAIN,
		Key:    key.GetKey(),
		lang:   "zh",
		client: &http.Client{Timeout: timeout},
	}
}

type newUnimailClient interface {
	GetHost() string
	GetKey() string
}

// NewUnimailClient 创建一个UnimailClient
//
// newUnimailClient 实现了 GetHost() string, GetKey() string 的结构体, 用于获取已授权的unimail host, key
func NewUnimailClient(client newUnimailClient) UnimailClient {
	if client == nil {
		panic("unimail client is nil")
	}
	return &unimail{
		Host:   client.GetHost(),
		Key:    client.GetKey(),
		lang:   "zh",
		client: &http.Client{Timeout: timeout},
	}
}

type unimail struct {
	Host   string
	Key    string
	lang   string
	client *http.Client
}

func (c *unimail) SetLanguage(lang string) error {
	if !slices.Contains(supportLang, lang) {
		return fmt.Errorf("unimail: language not support, you support is %s, support lang is: %v", lang, supportLang)
	}
	c.lang = lang
	return nil
}

func (c *unimail) CheckConnect() bool {
	urlPath := "/v2/checkConnection"
	var data = map[string]interface{}{
		"authorization": c.Key,
	}
	bdata, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", c.Host+urlPath, bytes.NewBuffer(bdata))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", c.lang)

	resp, err := c.client.Do(req)
	if err != nil { // 请求出错
		return false
	}
	var result Result
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	return result.IsSucess()
}

func (c *unimail) SendEmail(req UnimailReq) (result Result) {
	urlPath := "/v2/sendEmail"

	if len(req.Receivers) == 0 {
		return Result{Code: 400, Msg: "receivers is required"}
	}

	if req.Subject == "" {
		return Result{Code: 400, Msg: "subject is required"}
	}

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	bodyWriter.WriteField("authorization", c.Key)
	bodyWriter.WriteField("receiver", strings.Join(req.Receivers, ";"))
	bodyWriter.WriteField("from", req.From)
	if req.Cc != "" {
		bodyWriter.WriteField("cc", req.Cc)
	}
	if req.Bcc != "" {
		bodyWriter.WriteField("bcc", req.Bcc)
	}
	bodyWriter.WriteField("subject", req.Subject)
	if req.TxtContent != "" {
		bodyWriter.WriteField("txtContent", req.TxtContent)
	}
	if req.HtmlContent != "" {
		bodyWriter.WriteField("htmlContent", req.HtmlContent)
	}
	for id, attachment := range req.Attachments {
		bodyWriter.WriteField(fmt.Sprintf("attachments[%d].name", id), attachment.Name)
		if attachment.FileAttachment != nil {
			part, _ := bodyWriter.CreateFormFile(fmt.Sprintf("attachments[%d].fileAttachment", id), attachment.Name)
			io.Copy(part, attachment.FileAttachment)
		} else if attachment.UrlAttachment != "" {
			bodyWriter.WriteField(fmt.Sprintf("attachments[%d].urlAttachment", id), attachment.UrlAttachment)
		}
	}
	bodyWriter.Close()

	httpReq, _ := http.NewRequest("POST", c.Host+urlPath, bodyBuffer)
	httpReq.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	httpReq.Header.Set("Accept-Language", c.lang)
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return result.HttpError()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	return
}
