package model

import "github.com/tangchen2018/go-utils/http"

type Request struct {
	Path   *string
	Method *string
	Params BodyMap
	Body   BodyMap
	Req    *http.HttpRequest
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
