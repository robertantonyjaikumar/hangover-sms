package config

import "github.com/robertantonyjaikumar/hangover-common/config"

type Email struct {
	Host     string
	Port     string
	Sender   string
	Password string
}

func NewEmail() Email {
	return Email{
		Host:     config.CFG.V.GetString("email.host"),
		Port:     config.CFG.V.GetString("email.port"),
		Sender:   config.CFG.V.GetString("email.sender"),
		Password: config.CFG.V.GetString("email.password"),
	}
}
