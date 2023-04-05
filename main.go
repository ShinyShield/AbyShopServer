package main

import (
	"github.com/ShinyShield/AbyShopServer/config"
	"github.com/ShinyShield/AbyShopServer/pkg/server"
	"github.com/ShinyShield/AbyShopServer/pkg/store"
)

func main() {
	// 初始化設定檔案
	config.Init()
	// 初始化資料庫
	store.Init()
	// 初始化伺服器
	server.Init()
}
