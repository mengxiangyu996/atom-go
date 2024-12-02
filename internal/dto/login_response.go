package dto

// 验证码响应体
type CaptchaResponse struct {
	CaptchaId    string `json:"captchaId"`
	CaptchaImage string `json:"captchaImage"`
}
