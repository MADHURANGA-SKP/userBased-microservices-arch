package main

import (
	"gateway/gateway"
	"net/http"
)

type handler struct {
	gateway gateway.OMSGateway
}

func NewHandler(gateway gateway.OMSGateway) *handler {
	return &handler{gateway}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	
}