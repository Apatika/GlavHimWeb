package config

import "time"

const (
	defaultHttpPort               = "8080"
	defaultHttpRWTimeout          = 15 * time.Second
	defaultHttpMaxHeaderMegabytes = 1
)

type (
	HTTPconfig struct {
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}
)
