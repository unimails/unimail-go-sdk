// Package unimail is a golang sdk for unimail
package unimail

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

const UNIMAIL_VERSION = "1.0.0"

const DOMAIN = "https://uniapi.allcloud.top"

// DEBUG
// const DOMAIN = "http://127.0.0.1:8080/api/email"

var timeout = 120 * time.Second

var supportLang = []string{"en", "zh", "vi", "id", "gu", "th"}

// UnimailReq 发送邮件请求体
type UnimailReq struct {
	From        string            `json:"from,omitempty"`        // 昵称(可空)
	Receivers   []string          `json:"receivers,omitempty"`   // 收件人(必填)
	Cc          string            `json:"cc,omitempty"`          // 抄送(可空)
	Bcc         string            `json:"bcc,omitempty"`         // 密送(可空)
	Subject     string            `json:"subject,omitempty"`     // 主题(必填)
	TxtContent  string            `json:"txtContent,omitempty"`  // 文本内容(和htmlContent至少填一个)
	HtmlContent string            `json:"htmlContent,omitempty"` // HTML内容(和txtContent至少填一个)
	Attachments []EmailAttachment `json:"attachments,omitempty"` // 附件(可空)
}

// EmailAttachment 邮件附件
type EmailAttachment struct {
	Name           string
	FileAttachment io.Reader
	UrlAttachment  string
}

// 追加附件
func (a *UnimailReq) AppendAttachment(emailAttachment *EmailAttachment) error {
	if emailAttachment.FileAttachment == nil && emailAttachment.UrlAttachment == "" {
		return fmt.Errorf("unimail: attachment is empty")
	}
	a.Attachments = append(a.Attachments, *emailAttachment)
	return nil
}

// 追加文件附件
func (a *UnimailReq) AppendFile(name string, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("unimail: failed to open file: %v", err)
	}
	defer file.Close()

	// 读取文件内容到内存，避免 defer close 导致的问题
	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("unimail: failed to read file: %v", err)
	}

	return a.AppendAttachment(&EmailAttachment{
		Name:           name,
		FileAttachment: io.NopCloser(bytes.NewReader(content)),
	})
}

// 追加链接附件
func (a *UnimailReq) AppendUri(name string, url string) error {
	return a.AppendAttachment(&EmailAttachment{
		Name:          name,
		UrlAttachment: url,
	})
}

// Result 统一返回结果结构体
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *Result) IsSucess() bool {
	return c.Code == 0
}

func (c *Result) GetMsg() string {
	return c.Msg
}

func (c *Result) HttpError() Result {
	return Result{Code: 500, Msg: "http error"}
}
