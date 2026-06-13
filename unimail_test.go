package unimail

import (
	"testing"
)

var key = ""

func TestCheckConnection(t *testing.T) {
	c := New(key)

	if !c.CheckConnect() {
		t.Error("connection error")
	}
}

func TestSendEmail(t *testing.T) {
	c := New(key)
	req := UnimailReq{
		From:        "Demo",
		Cc:          "",
		Bcc:         "",
		Receivers:   []string{"email1@gmail.com", "email2@gmail.com"},
		Subject:     "test email",
		TxtContent:  "common attachment email test",
		HtmlContent: "<h1>common attachment email test this is a test <span style=\"font-size: 20px\">email 2</span></h1>",
	}
	req.AppendFile("attach1.txt", "./textAttachment.txt")
	// req.AppendAttachment(&EmailAttachment{
	// 	Name:           "filename.txt",
	// 	FileAttachment: xxx,
	// })
	result := c.SendEmail(req)
	if !result.IsSucess() {
		t.Errorf("send email error: %+v", result)
	}
}
