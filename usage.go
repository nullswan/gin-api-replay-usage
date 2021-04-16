package apiusage

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

func Register(tearDown func(e *ReplayData)) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := new(ReplayData)
		Before_Request(c, r)

		rw := &RespWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
			r:              r,
		}
		c.Writer = rw

		c.Next()

		After_Request(c, rw, &tearDown)
	}
}
