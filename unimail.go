// Package unimail is a golang sdk for unimail
package unimail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
)

// UnimailClient unimail client to operate unimail
type UnimailClient interface {
	//  CheckConnect 检查连接是否正常
	CheckConnect() bool

	// SetLanguage 设置语言 [zh, en, vi, th, gu, id]
	SetLanguage(lang string) error

	// SendEmail 发送邮件
	//
	// receiver 接收人, 多个人使用;进行分割
	// subject 邮件标题
	// content 邮件内容
	SendEmail(receiver string, subject string, content string) Result

	// BatchSendEmail 批量发送邮件
	//
	// receivers 接收人, 多个人使用
	// subject 邮件标题
	// content 邮件内容
	BatchSendEmail(receivers []string, subject string, content string) Result
	// todo
	SendEmailAsync(receiver string, subject string, content string) Result
	// todo
	BatchSendEmailAsync(receivers []string, subject string, content string) Result
	// todo
	CheckResult(key string) Result
}

type keyStruct interface {
	GetKey() string
}

// New 创建一个UnimailClient
//
// key Unimail 获取的已授权 key
func New(key string) UnimailClient {
	return &unimail{
		Host:   "https://unimail-back.allcloud.top",
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
		Host:   "https://unimail-back.allcloud.top",
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
	urlPath := "/api/email/checkConnection"
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

func (c *unimail) SendEmail(receiver string, subject string, content string) (result Result) {
	urlPath := "/api/email/sendEmail"
	var data = map[string]interface{}{
		"authorization": c.Key,
		"receiver":      receiver,
		"title":         subject,
		"content":       content,
	}
	bdata, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", c.Host+urlPath, bytes.NewBuffer(bdata))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", c.lang)

	resp, err := c.client.Do(req)
	if err != nil { // 请求出错
		return result.HttpError()
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	return
}

func (c *unimail) BatchSendEmail(receivers []string, subject string, content string) (result Result) {
	urlPath := "/api/email/batchSendEmail"
	var data = map[string]interface{}{
		"authorization": c.Key,
		"receivers":     receivers,
		"title":         subject,
		"content":       content,
	}
	bdata, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", c.Host+urlPath, bytes.NewBuffer(bdata))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", c.lang)

	resp, err := c.client.Do(req)
	if err != nil { // 请求出错
		return result.HttpError()
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	return
}

// todo
func (c *unimail) SendEmailAsync(receiver string, subject string, content string) (result Result) {
	return c.SendEmail(receiver, subject, content)
}

// todo
func (c *unimail) BatchSendEmailAsync(receivers []string, subject string, content string) (result Result) {
	return c.BatchSendEmail(receivers, subject, content)

}

// todo
func (c *unimail) CheckResult(key string) (result Result) {
	return result.HttpError()
}
