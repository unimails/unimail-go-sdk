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

	result := c.SendEmail("i-curve@qq.com", "go test", "this a go client test email")
	if !result.IsSucess() {
		t.Errorf("send email error: %+v", result)
	}
}

func TestBatchSendEmail(t *testing.T) {
	c := New(key)

	result := c.BatchSendEmail([]string{"i-curve@qq.com", "i_curve@qq.com"}, "go test", "this a go client batch test email")
	if !result.IsSucess() {
		t.Errorf("send email error: %+v", result)
	}
}
