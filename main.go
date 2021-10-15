package main

import (
	"github.com/pwh19920920/butterfly"
	"github.com/pwh19920920/butterfly-admin/src/app/starter"
)

func main() {
	_ = starter.InitButterflyAdmin()
	butterfly.Run()
}
