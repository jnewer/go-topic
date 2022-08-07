package main

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "golang-topic"
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
		Path:   "/",
	})

	return store
}

func Session(keyPairs string) gin.HandlerFunc {
	return sessions.Sessions(keyPairs, SessionConfig())
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "application/octet-stream")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audit/x-wav")
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

func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}

	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}

	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}

func captchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		}

		return false
	}

	return false
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./index.html")
	router.Use(Session("golang-topic"))
	router.GET("/captcha", func(c *gin.Context) {
		Captcha(c, 4)
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/captcha/verify", func(c *gin.Context) {
		value := c.PostForm("code")

		if captchaVerify(c, value) {
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})

	router.Run(":8080")
}
