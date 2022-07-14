package gin

import gonicgin "github.com/gin-gonic/gin"

type contextExt struct{}

// GetParams 获取请求参数
func (ce contextExt) GetParams(c *gonicgin.Context) map[string]string {
	params := make(map[string]string)
	if c.Request.Method == "GET" {
		for k, v := range c.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	} else {
		err := c.Request.ParseForm()
		if err != nil {
			return params
		}
		for k, v := range c.Request.PostForm {
			params[k] = v[0]
		}
		for k, v := range c.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	}
}
