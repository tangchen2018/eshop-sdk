package tiktok

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	http2 "net/http"
	"sort"
	"strings"
	"time"
)

type Client struct {
	model.Client
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{Setting: setting}}
}

func (p *Client) Execute() {

	if p.Request.Method == nil {
		p.SetMethod("GET")
	}
	if p.Request.Params == nil {
		p.SetParams(make(model.BodyMap))
	}
	if p.Setting.Key == nil {
		p.Err = utils.Err("Key is null..")
		return
	}
	if p.Request.Path == nil {
		p.Err = utils.Err("Path is null..")
		return
	}

	p.Request.Params.Set("app_key", *p.Setting.Key).
		Set("timestamp", fmt.Sprintf("%d", time.Now().Unix()))

	if p.Setting.AccessToken != nil {
		p.Request.Params.Set("access_token", *p.Setting.AccessToken)
	}
	if p.Setting.ShopId != nil {
		p.Request.Params.Set("shop_id", *p.Setting.ShopId)
	}

	if *p.Request.Path != GETACCESS {
		p.Request.Params.Set("sign", p.sign())
	}

	p.Client.Request.Req = &http.HttpRequest{
		Method: *p.Request.Method,
		Url:    p.urlParse(),
	}

	for key, value := range p.Request.Params {
		p.Client.Request.Req.SetParams(key, value.(string))
	}

	p.Client.Request.Req.SetHeader("Content-Type", "application/json")

	if strings.ToUpper(*p.Request.Method) == http2.MethodPost ||
		strings.ToUpper(*p.Request.Method) == http2.MethodPut {
		p.Client.Request.Req.SetBody([]byte(p.Request.Body.JsonBody()))
	}

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	result := new(Response)
	if p.Err = p.Client.Request.Req.To(result); p.Err != nil {
		return
	}

	p.Response.Response.Code = fmt.Sprintf("%d", result.Code)
	p.Response.Response.Message = result.Message
	p.Response.Response.RequestId = result.RequestId
	p.Response.Response.Data = result.Response

	if result.Code == 0 {
		p.Response.Success = true
	}
}

func (p *Client) sign() string {

	keys := []string{}
	union := map[string]string{}

	for key, val := range p.Client.Request.Params {
		if key == "access_token" || key == "sign" {
			continue
		}
		union[key] = val.(string)
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var message bytes.Buffer
	message.WriteString(fmt.Sprintf("%s%s", *p.Client.Setting.Secret, *p.Client.Request.Path))
	for _, key := range keys {
		message.WriteString(fmt.Sprintf("%s%s", key, union[key]))
	}
	message.WriteString(*p.Client.Setting.Secret)
	msg := message.String()
	hash := hmac.New(sha256.New, []byte(*p.Client.Setting.Secret))
	if _, err := hash.Write([]byte(msg)); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func (p *Client) urlParse() string {
	if *p.Client.Request.Path == GETACCESS || *p.Client.Request.Path == REFRESHTOKEN {
		return fmt.Sprintf("%s%s", AUTHSITE, *p.Client.Request.Path)
	} else {
		return fmt.Sprintf("%s%s", SERVER_URl, *p.Client.Request.Path)
	}
}
