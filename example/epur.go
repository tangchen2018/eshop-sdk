package main

//
//import (
//	"fmt"
//	"github.com/tangchen2018/eshop-sdk/model"
//	"github.com/tangchen2018/eshop-sdk/pfc/epur"
//)
//
//func main() {
//	api := epur.New(
//		new(model.Setting).
//			SetKey("appxxxxxxxxxxxxxxxx").
//			SetSecret("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").
//			SetServerUrl("http://dev.open.epur.cn").
//			SetShopId("appxxxxxxxxxxxxx__&&@@%%__110").
//			SetAuthCallbackUrl("https://www.baidu.com").
//			SetAccessToken(`W8picM3Zs0YMA1g9KvJPgvP7kHbehwjkSTIDMUWLuaok9L5Ua7y999YUoalZ`),
//	)
//	RefreshToken(api)
//}
////
////func GetAuthUrl(api *epur.Api) {
////	result := api.GetAuthUrl("123")
////	fmt.Println(result)
////}
////
////func GetToken(api *epur.Api) {
////	c := api.GetToken(model.BodyMap{"code": "PNYzImsQzisSBFzLY1ewE24p261IAcR6TqdKZDbGs0KrxC29ynVEydjAFmjp"})
////	if c.Err != nil {
////		panic(c.Err)
////	}
////	result := c.GetResponseTo().(epur.GetTokenResponse)
////	fmt.Println(result)
////}
////
////func RefreshToken(api *epur.Api) {
////	c := api.RefreshToken(model.BodyMap{"refreshToken": "LquqbTsBo0VN0SbE0ymwWvyd4SURTaZzK5cm5kAZvOYfMT9FPvRhexpwvaZJ"})
////	if c.Err != nil {
////		panic(c.Err)
////	}
////	result := c.GetResponseTo().(epur.GetTokenResponse)
////	fmt.Println(result)
////}
