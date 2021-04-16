package apiusage

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Before_Request(c *gin.Context, r *ReplayData) {
	r.Time = time.Now().UnixNano()
	r.Host = c.Request.Host
	r.Method = c.Request.Method
	r.URL = r.Host + r.Path
	r.Path = c.Request.URL.Path

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	r.Request.Body = string(buf[0:num])

	r.Request.Content_length = c.Request.ContentLength
	r.Request.Headers = make(map[string]string)
	for k, v := range c.Request.Header {
		r.Request.Headers[k] = strings.Join(v, ",")
	}
	r.Request.Mimetype = r.Request.Headers["Content-Type"]
	r.Request.Args = make(HeadersProfiling)
	for k, v := range c.Request.URL.Query() {
		r.Request.Args[k] = v
	}

	r.Client.Referer = c.Request.Referer()
	r.Client.Authorization = r.Request.Headers["Authorization"]
	r.Client.Origin = r.Request.Headers["Origin"]
	r.Client.Remote_address = c.ClientIP()
	r.Client.User_agent = c.Request.UserAgent()
}

func After_Request(c *gin.Context, rw *RespWriter, tearDown *func(*ReplayData)) {
	rw.r.Status_code = c.Writer.Status()

	rw.r.Request.Process_time = time.Now().UnixNano() - rw.r.Time

	rw.r.Response.Body = rw.body.String()
	if rw.r.Status_code != 200 && rw.r.Status_code != 301 {
		var m ResponseError

		rw.r.Error = true
		err := json.Unmarshal(rw.body.Bytes(), &m)
		if err == nil {
			rw.r.Error_msg = m.Message
		} else {
			rw.r.Error_msg = "replay: invalid response formating"
		}
	}

	rw.r.Response.Content_length = rw.body.Len()
	rw.r.Response.Mimetype = c.Writer.Header().Get("Content-Type")
	rw.r.Response.Headers = make(HeadersProfiling)
	for k, v := range c.Writer.Header() {
		rw.r.Response.Headers[k] = v
	}

	go (*tearDown)(rw.r)
}
