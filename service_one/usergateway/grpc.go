package usergateway

import (
	"common/discovery"
	"common/util"
)

type gateway struct {
	registry discovery.Registry
	config util.Config
}

func NewUserGateway(registry discovery.Registry, config util.Config) *gateway {
	return &gateway{registry, config}
}

