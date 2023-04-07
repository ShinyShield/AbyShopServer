package main

import (
	"github.com/ShinyShield/AbyShopServ/config"
	"github.com/ShinyShield/AbyShopServ/pkg/server"
	"github.com/ShinyShield/AbyShopServ/pkg/store"
)

func main() {

	config.Init()
	server.Init()
	store.Init()
}
