package main

import (
	"github.com/ShinyShield/AbyShopServ/config"
	"github.com/ShinyShield/AbyShopServ/pkg/server"
	"github.com/ShinyShield/AbyShopServ/pkg/store"
)

func main() {
	// 初始化設定檔案
	config.Init()
	// 初始化資料庫
	server.Init()
	store.Init()
	// 初始化伺服器
}
