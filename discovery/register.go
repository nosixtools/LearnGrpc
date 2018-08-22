package discovery

import "time"

type RegisterInfo struct {
	Host           string
	Port           int
	Weight         int
	ServiceName    string
	BasePath       string
	Version        string
	UpdateInterval time.Duration
}

type Register interface {
	Register(info RegisterInfo) error
	DeRegister(info RegisterInfo) error
}
