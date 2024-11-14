package main

import (
	"crypto/tls"
	"time"
)

type Http2 struct {
	Enable           bool          `json:"Enable" yaml:"Enable"`
	DisableKeepalive bool          `json:"DisableKeepalive" yaml:"DisableKeepalive"`
	ReadTimeout      time.Duration `json:"ReadTimeout" yaml:"ReadTimeout"`
}
type HertzCfg struct {
	UseNetpoll bool  `json:"UseNetpoll" yaml:"UseNetpoll"`
	Http2      Http2 `json:"Http2" yaml:"Http2"`
	Tls        Tls   `json:"Tls" yaml:"Tls"`
}

type Tls struct {
	Enable bool `json:"Enable" yaml:"Enable"`
	Cfg    tls.Config
	Cert   string `json:"CertFile" yaml:"CertFile"`
	Key    string `json:"KeyFile" yaml:"KeyFile"`
	ALPN   bool   `json:"ALPN" yaml:"ALPN"`
}
