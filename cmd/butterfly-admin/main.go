package main

import (
	"github.com/pwh19920920/butterfly-admin/internal/starter"
	"github.com/pwh19920920/butterfly/pkg/server"
)

func main() {
	_, _ = starter.InitButterflyAdmin()
	server.StartHttpServer()
}
