package test

import (
	"encoding/base64"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
	"strings"
	"testing"
	"time"
)

func TestBase64(t *testing.T) {
	text := "你好，大乔乔"
	cipher := base64.StdEncoding.EncodeToString([]byte(text)) //base64是一种公开透明的转码方式，不是加密算法
	fmt.Println(cipher)
	bs, _ := base64.StdEncoding.DecodeString(cipher)
	if string(bs) != text {
		t.Fail()
	}
}

func TestJWT(t *testing.T) {
	secret := "123456"
	header := jwt.DefautHeader
	payload := jwt.JwtPayload{
		ID:          "rj4t49tu49",
		Issue:       "微信",
		Audience:    "王者荣耀",
		Subject:     "购买道具",
		IssueAt:     time.Now().Unix(),
		Expiration:  time.Now().Add(2 * time.Hour).Unix(),
		UserDefined: map[string]any{"name": strings.Repeat("Ageless", 100)}, //When the amount of information is large, the JWT length may exceed 4K.
	}

	if token, err := jwt.GenJWT(header, payload, secret); err != nil {
		fmt.Printf("generating json web token failed: %v", err)
	} else {
		fmt.Println(token)
		if _, p, err := jwt.VerifyJwt(token, secret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("JWT auth passed, welcome %s !\n", p.UserDefined["name"])
		}
	}
}
