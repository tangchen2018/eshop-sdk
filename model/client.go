package model

import (
	"encoding/json"
	"github.com/tangchen2018/go-utils/http"
)

type Response struct {
	Success  bool `json:"success"`    // 是否成功
	Status   int  `json:"httpStatus"` // http状态码
	Response struct {
		RequestId string          `json:"requestId"` // 请求ID
		Code      string          `json:"code"`      // 平台code信息
		Message   string          `json:"message"`   // 平台msg信息
		Data      json.RawMessage `json:"data"`      // 数据
		DataTo    interface{}     `json:"dataTo"`    // 结构转化后的数据
	}
}

func (p *Response) ToMap() BodyMap {
	result := make(BodyMap)
	_ = json.Unmarshal(p.Response.Data, &result)
	return result
}

func (p *Response) To(row interface{}) error {
	return json.Unmarshal(p.Response.Data, row)
}

type Request struct {
	Path   *string
	Method *string
	Params BodyMap
	Body   BodyMap
	Req    *http.Request
}

type Client struct {
	Request  Request
	Response Response
	Setting  *Setting
	Err      error
}

func (p *Client) Execute() error {

	p.Response.Success = false

	if err := p.Request.Req.Do(); err != nil {
		return err
	} else {
		p.Response.Status = p.Request.Req.Response.StatusCode
	}

	return nil
}

func (c *Client) GetResponseTo() interface{} {
	return c.Response.Response.DataTo
}

func (c *Client) SetPath(data string) *Client {
	c.Request.Path = &data
	return c
}

func (c *Client) SetMethod(data string) *Client {
	c.Request.Method = &data
	return c
}

func (c *Client) SetParams(data BodyMap) *Client {
	c.Request.Params = data
	return c
}

func (c *Client) SetBody(data BodyMap) *Client {
	c.Request.Body = data
	return c
}
