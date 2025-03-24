package config

import "github.com/robertantonyjaikumar/hangover-common/config"

type Jwt struct {
	Secret               string
	AccessTokenExpireIn  int
	RefreshTokenExpireIn int
}

func NewJwt() (jwt Jwt) {
	return Jwt{
		Secret:               config.CFG.V.GetString("jwt.secret"),
		AccessTokenExpireIn:  config.CFG.V.GetInt("jwt.access_token_expire_in"),
		RefreshTokenExpireIn: config.CFG.V.GetInt("jwt.refresh_token_expire_in"),
	}
}
