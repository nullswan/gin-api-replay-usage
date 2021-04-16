package apiusage

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type RespWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
	r    *ReplayData
}

func (r RespWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HeadersProfiling map[string][]string

type ReplayData struct {
	Time                    int64
	Host, Path, Method, URL string
	Status_code             int
	Error                   bool
	Error_msg               string
	Client                  struct {
		Referer, Authorization, Origin, Remote_address, User_agent, Platform, Browser string
	}
	Request struct {
		Process_time, Content_length int64
		Mimetype, Body               string
		Headers                      map[string]string
		Args                         HeadersProfiling
	}
	Response struct {
		Content_length int
		Mimetype, Body string
		Headers        HeadersProfiling
	}
}
