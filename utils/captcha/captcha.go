package captcha

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

// 生成验证码
func Captcha(ctx *gin.Context) {
	l := captcha.DefaultLen
	w, _ := strconv.Atoi(viper.GetString("captcha.w"))
	h, _ := strconv.Atoi(viper.GetString("captcha.h"))
	l, _ = strconv.Atoi(viper.GetString("captcha.length"))
	captchaId := captcha.NewLen(l)
	session := sessions.Default(ctx)
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = Serve(ctx.Writer, ctx.Request, captchaId, ".png", viper.GetString("captcha.lang"), false, w, h)
}

// 验证码的header头设置，生成页面显示
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

// 验证
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
