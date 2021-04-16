# gin-api-replay-usage

See examples for passing func to tearDown, can be used to save document in (any) database or direct parsing through rpcx.

Errors Assertion:\
Error field is automatically set when response is different from 301 (Redirection) and 200 (Success)
Error_msg is set when you return error in your body.

```
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

{
	"code": 404,
	"error": "not found"
}
```


## Example returned typed Struct
```
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
		Args                         map[string][]string
	}
	Response struct {
		Content_length int
		Mimetype, Body string
		Headers        map[string][]string
	}
}
```

## Example Output Document
```
{
		"Time": 1618606390449228744,
		"Host": "localhost:3000",
		"Path": "/stars",
		"Method": "GET",
		"URL": "localhost:3000",
		"Status_code": 200,
		"Error": false,
		"Error_msg": "some error msg if any",
		"Client": {
				"Referer": "some ref if any",
				"Authorization": "Bearer Authorization",
				"Origin": "some origin if any",
				"Remote_address": "::1",
				"User_agent": "Mozilla/5.0 ... (Truncated Manually)",
				"Platform": "",
				"Browser": ""
		},
		"Request": {
				"Process_time": 9496,
				"Content_length": 38,
				"Mimetype": "application/json",
				"Body": "{"username": "usr", "password": "pwd"}",
				"Headers": {
					"Upgrade-Insecure-Requests" : "1",
					"Cache-Control" : "max-age=0",
					"Accept" : "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
					"Accept-Language" : "en-US,en;q=0.5",
					"Accept-Encoding" : "gzip, deflate",
					"Connection" : "keep-alive"
					"User-Agent": "Mozilla/5.0 ... (Truncated Manually)",
				},
				"Args": {
					"a": [
						"87"
					]
				}
		},
		"Response": {
				"Content_length": 6,
				"Mimetype": "application/json; charset=utf-8",
				"Body": "\"pong"",
				"Headers": {
						"Content-Type": [
								"application/json; charset=utf-8"
						]
				}
		}
}
```