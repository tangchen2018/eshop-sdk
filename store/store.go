package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tangchen2018/eshop-sdk/api"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/utils"
	"log"
	"reflect"
	"sync"
	"time"
)

var TokenStore *Store

type Event struct {
	Success      bool //是否成功
	Msg          string
	Token        *model.Token
	ResponseData []byte
}

type Store struct {
	TokenMap *sync.Map
	count    int
	//TokenList            Tokens
	CallBack             func(e *Event) // accessToken刷新回调
	SecondsBeforeRefresh int64          // 到期提前多少秒刷新accessToken
}

type Request struct {
	Key                string // Key
	Secret             string // 密钥
	ShopId             string `json:"shopId"`
	PlatformCode       string `json:"platformCode"`
	AccessToken        string `json:"accessToken"`          //访问令牌
	RefreshToken       string `json:"refreshToken"`         //刷新令牌
	AccessTokenExpire  int64  `json:"accessTokenExpireIn"`  //访问令牌有效期(时间戳)
	RefreshTokenExpire int64  `json:"refreshTokenExpireIn"` //刷新令牌有效期(时间戳)
}

func (p *Store) Id(r *Request) string {
	return fmt.Sprintf("%s-%s", r.PlatformCode, r.ShopId)
}

func (p *Store) Save(r *Request) (*model.Token, error) {

	var t *model.Token

	id := p.Id(r)
	if row, ok := p.TokenMap.Load(id); ok {
		t = row.(*model.Token)
		if t.TaskStatus == "2" {
			return nil, errors.New("正在执行中,请稍等...")
		}
	} else {
		p.count++
	}

	t = &model.Token{
		Id:                 p.Id(r),
		AccessToken:        r.AccessToken,
		RefreshToken:       r.RefreshToken,
		AccessTokenExpire:  r.AccessTokenExpire,
		RefreshTokenExpire: r.RefreshTokenExpire,
		PlatformCode:       r.PlatformCode,
		TaskStatus:         "0",
		Setting: model.Setting{
			Key:         &r.Key,
			Secret:      &r.Secret,
			ShopId:      &r.ShopId,
			AccessToken: &r.AccessToken,
		},
	}

	p.TokenMap.Store(t.Id, t)
	p.SyncTask(t)
	return t, nil
}

func (p *Store) Del(id string) {
	p.TokenMap.Delete(id)
	p.count--
}

func New() *Store {
	TokenStore = &Store{TokenMap: new(sync.Map), SecondsBeforeRefresh: 30 * 60}
	return TokenStore
}

func (p *Store) Listen() {
	//if len(p.TokenList) > 0 && p.TokenList[0].AccessTokenExpire-p.SecondsBeforeRefresh >= utils.TimestampSecond() {
	//
	//}

	for {
		log.Println("token数量->", p.count)

		id := p.Id(&Request{PlatformCode: model.PFC_TIKTOK, ShopId: "5"})
		if row, ok := p.TokenMap.Load(id); ok {
			t := row.(*model.Token)
			tmp, _ := json.Marshal(t)
			log.Println(string(tmp))
		}
		time.Sleep(10 * time.Second)
	}

	//<-make(chan bool)
}

func (p *Store) SyncTask(t *model.Token) {
	if t.TaskStatus == "0" {
		t.TaskStatus = "1"
		diff := t.AccessTokenExpire - p.SecondsBeforeRefresh - utils.TimestampSecond()

		log.Printf("延迟时间[%d]", diff)

		time.AfterFunc(time.Duration(diff)*time.Second, p.callback(t))
	}
}

func (p *Store) callback(t *model.Token) func() {

	return func() {
		p.callbackRun(p.callbackBefore(t))
		t.TaskStatus = "0"
	}
}

func (p *Store) callbackBefore(t *model.Token) *Event {

	e := &Event{}

	defer func() {
		err := recover()
		if err != nil {
			e1 := reflect.ValueOf(err)
			e.Success = false
			e.Msg = e1.String()
		}
	}()

	t.TaskStatus = "2"
	if p.CallBack == nil {
		e.Msg = "error: callback is void!"
		return e
	}

	c := api.New(t.PlatformCode, &t.Setting).RefreshToken(model.BodyMap{"refresh_token": t.RefreshToken})
	if c.Err != nil {
		e.Msg = c.Err.Error()
	} else if !c.Response.Success {
		e.Msg = c.Response.Response.Message
	} else if c.Response.Success {

		tokenResponse := c.Response.Response.DataTo.(model.Token)

		e.ResponseData = c.Response.Response.Data

		e.Success = true
		t.TaskStatus = "0"
		if t1, err := p.Save(&Request{
			Key:                *t.Setting.Key,
			Secret:             *t.Setting.Secret,
			ShopId:             *t.Setting.ShopId,
			PlatformCode:       t.PlatformCode,
			RefreshToken:       tokenResponse.RefreshToken,
			AccessToken:        tokenResponse.AccessToken,
			AccessTokenExpire:  tokenResponse.AccessTokenExpire,
			RefreshTokenExpire: tokenResponse.RefreshTokenExpire,
		}); err != nil {
			e.Success = false
			e.Msg = err.Error()
		} else {
			e.Token = t1
		}
	}
	return e
}

func (p *Store) callbackRun(e *Event) {

	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		e1 := reflect.ValueOf(err)
	//		log.Fatal("callbackRun->", e1.String())
	//	}
	//}()
	p.CallBack(e)
}
