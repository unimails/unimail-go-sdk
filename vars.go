// Package unimail is a golang sdk for unimail
package unimail

import "time"

const UNIMAIL_VERSION = "0.3.0"

var timeout = 60 * time.Second

var supportLang = []string{"en", "zh", "vi", "id", "gu", "th"}

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
