package captcha

import "github.com/mojocn/base64Captcha"

type Captcha struct {
	captcha *base64Captcha.Captcha
}

// 初始化验证码
func NewCaptcha() *Captcha {

	store := base64Captcha.DefaultMemStore

	driver := base64Captcha.NewDriverDigit(40, 120, 4, 0.7, 1)

	return &Captcha{
		captcha: base64Captcha.NewCaptcha(driver, store),
	}
}

// 生成验证码
// uuid, base64, answer
func (c *Captcha) Generate() (string, string) {

	id, b64s, _, err := c.captcha.Generate()
	if err != nil {
		return "", ""
	}
	
	return id, b64s
}

// 验证验证码
func (c *Captcha) Verify(id, answer string) bool {
	return c.captcha.Verify(id, answer, true)
}