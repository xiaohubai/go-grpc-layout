package model

type CaptchaResp struct {
	CaptchaID     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}
