package main

import (
	"github.com/webronin26/shopping-server/config"
	"github.com/webronin26/shopping-server/pkg/server"
	"github.com/webronin26/shopping-server/pkg/store"
)

func main() {
	// 初始化設定檔案
	config.Init()
	// 初始化資料庫
	store.Init()
	// 初始化伺服器
	server.Init()
}
