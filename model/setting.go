package model

type Setting struct {
	Key         *string // Key
	Secret      *string // 密钥
	ShopId      *string // 店铺ID
	RetryCount  *int    // 重试次数
	AccessToken *string // 刷新令牌
}

func (c *Setting) SetAccessToken(data string) *Setting {
	if len(data) > 0 {
		c.AccessToken = &data
	}
	return c
}

func (c *Setting) SetShopId(data string) *Setting {
	if len(data) > 0 {
		c.ShopId = &data
	}
	return c
}

func (c *Setting) SetKey(data string) *Setting {
	if len(data) > 0 {
		c.Key = &data
	}
	return c
}

func (c *Setting) SetSecret(data string) *Setting {
	if len(data) > 0 {
		c.Secret = &data
	}
	return c
}

func (c *Setting) SetRetryCount(data int) *Setting {
	if data > 0 {
		c.RetryCount = &data
	}
	return c
}
